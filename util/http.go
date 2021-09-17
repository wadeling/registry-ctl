package util

import (
	"crypto/tls"
	"os"
	"strings"
)

const (
	// Internal TLS ENV
	internalVerifyClientCert = "INTERNAL_VERIFY_CLIENT_CERT"
)

// NewServerTLSConfig returns a modern tls config,
// refer to https://blog.cloudflare.com/exposing-go-on-the-internet/
func NewServerTLSConfig() *tls.Config {
	return &tls.Config{
		PreferServerCipherSuites: true,
		CurvePreferences: []tls.CurveID{
			tls.CurveP256,
			tls.X25519,
		},
		MinVersion: tls.VersionTLS12,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		},
	}
}

func InternalEnableVerifyClientCert() bool {
	return strings.ToLower(os.Getenv(internalVerifyClientCert)) == "true"
}
