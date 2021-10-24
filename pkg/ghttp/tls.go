package ghttp

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
)

func createTlsConfig(certFilePath, keyFilePath, caFilePath string) *tls.Config {
	if certFilePath == "" && keyFilePath == "" && caFilePath == "" {
		return &tls.Config{}
	}
	certs, err := tls.LoadX509KeyPair(certFilePath, keyFilePath)
	if err != nil {
		panic(err)
	}

	ca, err := ioutil.ReadFile(caFilePath)
	if err != nil {
		panic(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(ca)

	return &tls.Config{
		Certificates: []tls.Certificate{certs},
		RootCAs:      caCertPool,
	}
}
