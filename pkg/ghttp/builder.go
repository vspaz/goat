package ghttp

import "time"

type ClientBuilder interface {
	Host(host string) ClientBuilder
	Auth(user string, password string) ClientBuilder
	Tls(certFilePath, keyFilePath, CaFilePath string) ClientBuilder
	UserAgent(ua string) ClientBuilder
	RetryCount(num int) ClientBuilder
	Delay(num float64) ClientBuilder
	ConnTimeout(timeout time.Duration) ClientBuilder
	ReadTimeout(timeout time.Duration) ClientBuilder
}

type clientBuilder struct {
	host string

	basicAuthUser     string
	basicAuthPassword string

	tlsCertFilePath string
	tlsKeyFilePath  string
	tlsCaFilePath   string

	userAgent string

	retryCount int
	delay      float64

	connTimeout time.Duration
	readTimeout time.Duration
}
