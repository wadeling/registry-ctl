package main

import (
	"crypto/tls"
	"flag"
	"github.com/wadeling/registry-ctl/config"
	"github.com/wadeling/registry-ctl/handlers"
	"github.com/wadeling/registry-ctl/registry"
	"github.com/wadeling/registry-ctl/util"
	"log"
	"net/http"
)

// RegistryCtl for registry controller
type RegistryCtl struct {
	ServerConf config.Configuration
	Handler    http.Handler
}

// Start the registry controller
func (s *RegistryCtl) Start() {
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
		log.Fatal(err)
	}

	return
}

func main() {
	configPath := flag.String("c", "", "Specify registryCtl configuration file path")
	flag.Parse()

	if configPath == nil || len(*configPath) == 0 {
		flag.Usage()
		log.Fatal("Config file should be specified")
	}
	if err := config.DefaultConfig.Load(*configPath, true); err != nil {
		log.Fatalf("Failed to load configurations with error: %s\n", err)
	}

	// start registry
	err := registry.StartRegistry()
	if err != nil {
		log.Fatalf("failed to start registry:%v\n", err)
		return
	}

	// start http server
	regCtl := &RegistryCtl{
		ServerConf: *config.DefaultConfig,
		Handler:    handlers.NewHandlerChain(*config.DefaultConfig),
	}
	regCtl.Start()
}
