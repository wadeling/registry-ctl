package main

import (
	"crypto/tls"
	"flag"
	"github.com/sirupsen/logrus"
	"github.com/wadeling/registry-ctl/config"
	"github.com/wadeling/registry-ctl/handlers"
	"github.com/wadeling/registry-ctl/registry"
	"github.com/wadeling/registry-ctl/util"
	"net/http"
)

// RegistryCtl for registry controller
type RegistryCtl struct {
	ServerConf config.Configuration
	Handler    http.Handler
}

// Start the registry controller
func (s *RegistryCtl) Start() {
	logrus.Info("registry ctl start...")
	regCtl := &http.Server{
		Addr:      ":" + s.ServerConf.Port,
		Handler:   s.Handler,
		TLSConfig: util.NewServerTLSConfig(),
	}

	// start http server
	var err error
	if s.ServerConf.Protocol == "https" {
		if util.InternalEnableVerifyClientCert() {
			regCtl.TLSConfig.ClientAuth = tls.RequireAndVerifyClientCert
		}
		err = regCtl.ListenAndServeTLS(s.ServerConf.HTTPSConfig.Cert, s.ServerConf.HTTPSConfig.Key)
	} else {
		err = regCtl.ListenAndServe()
	}

	if err != nil {
		logrus.Fatal(err)
	}

	return
}

func initLog() {
	lvl, err := logrus.ParseLevel(config.GetLogLevel())
	if err != nil {
		lvl = logrus.InfoLevel
	}
	logrus.SetLevel(lvl)
}

func main() {
	configPath := flag.String("c", "", "Specify registryCtl configuration file path")
	flag.Parse()

	if configPath == nil || len(*configPath) == 0 {
		flag.Usage()
		logrus.Fatal("Config file should be specified")
	}
	if err := config.DefaultConfig.Load(*configPath, true); err != nil {
		logrus.Fatalf("Failed to load configurations with error: %s\n", err)
	}

	// init log level
	initLog()

	// start registry
	go func() {
		err := registry.StartRegistry(config.DefaultConfig.RegistryConfig)
		if err != nil {
			logrus.Fatalf("failed to start registry:%v\n", err)
		}
	}()

	// start http server
	regCtl := &RegistryCtl{
		ServerConf: *config.DefaultConfig,
		Handler:    handlers.NewHandlerChain(*config.DefaultConfig),
	}
	regCtl.Start()
}
