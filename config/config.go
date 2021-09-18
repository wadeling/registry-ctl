package config

import (
	"fmt"
	"github.com/docker/distribution/configuration"
	storagedriver "github.com/docker/distribution/registry/storage/driver"
	"github.com/docker/distribution/registry/storage/driver/factory"
	_ "github.com/docker/distribution/registry/storage/driver/filesystem"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// DefaultConfig ...
var DefaultConfig = &Configuration{}

// Configuration loads the configuration of registry controller.
type Configuration struct {
	Protocol    string `yaml:"protocol"`
	Port        string `yaml:"port"`
	LogLevel    string `yaml:"log_level"`
	HTTPSConfig struct {
		Cert string `yaml:"cert"`
		Key  string `yaml:"key"`
	} `yaml:"https_config,omitempty"`
	RegistryConfig string                      `yaml:"registry_config"`
	StorageDriver  storagedriver.StorageDriver `yaml:"-"`
}

// Load the configuration options from the specified yaml file.
func (c *Configuration) Load(yamlFilePath string, detectEnv bool) error {
	if len(yamlFilePath) != 0 {
		// Try to load from file first
		data, err := ioutil.ReadFile(yamlFilePath)
		if err != nil {
			return err
		}
		if err = yaml.Unmarshal(data, c); err != nil {
			return err
		}
	}

	if detectEnv {
		c.loadEnvs()
	}

	if err := c.setStorageDriver(); err != nil {
		log.Errorf("failed to load storage driver, err:%v", err)
		return err
	}

	return nil
}

// setStorageDriver set the storage driver according the registry's configuration.
func (c *Configuration) setStorageDriver() error {
	fp, err := os.Open(c.RegistryConfig)
	if err != nil {
		return err
	}
	defer fp.Close()
	rConf, err := configuration.Parse(fp)
	if err != nil {
		return fmt.Errorf("error parsing registry configuration %s: %v", c.RegistryConfig, err)
	}
	storageDriver, err := factory.Create(rConf.Storage.Type(), rConf.Storage.Parameters())
	if err != nil {
		return err
	}
	c.StorageDriver = storageDriver
	return nil
}

// GetLogLevel returns the log level
func GetLogLevel() string {
	return DefaultConfig.LogLevel
}

// loadEnvs Load env variables
func (c *Configuration) loadEnvs() {
	prot := os.Getenv("REGISTRYCTL_PROTOCOL")
	if len(prot) != 0 {
		c.Protocol = prot
	}

	p := os.Getenv("PORT")
	if len(p) != 0 {
		c.Port = p
	}

	// Only when protocol is https
	if c.Protocol == "HTTPS" {
		cert := os.Getenv("REGISTRYCTL_HTTPS_CERT")
		if len(cert) != 0 {
			c.HTTPSConfig.Cert = cert
		}

		certKey := os.Getenv("REGISTRYCTL_HTTPS_KEY")
		if len(certKey) != 0 {
			c.HTTPSConfig.Key = certKey
		}
	}

	loggerLevel := os.Getenv("LOG_LEVEL")
	if len(loggerLevel) != 0 {
		c.LogLevel = loggerLevel
	}

	registryConf := os.Getenv("REGISTRY_CONFIG")
	if len(registryConf) != 0 {
		c.RegistryConfig = registryConf
	}

}
