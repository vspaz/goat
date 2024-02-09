package ghttp

import (
	"crypto/tls"
	"crypto/x509"
	"os"
)

func loadKeyPair(certFilePath string, keyFilePath string) tls.Certificate {
	if certFilePath == "" && keyFilePath == "" {
		return tls.Certificate{}
	}
	certs, err := tls.LoadX509KeyPair(certFilePath, keyFilePath)
	if err != nil {
		panic(err)
	}
	return certs
}

func loadCa(caFilePath string) *x509.CertPool {
	if caFilePath == "" {
		return nil
	}
	ca, err := os.ReadFile(caFilePath)
	if err != nil {
		panic(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(ca)
	return caCertPool
}

func createTlsConfig(certFilePath, keyFilePath, caFilePath string) *tls.Config {
	return &tls.Config{
		Certificates: []tls.Certificate{loadKeyPair(certFilePath, keyFilePath)},
		RootCAs:      loadCa(caFilePath),
	}
}
