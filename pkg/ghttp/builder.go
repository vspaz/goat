package ghttp

import (
	"log"
	"time"
)

type ClientBuilder interface {
	Host(host string) ClientBuilder
	Auth(user string, password string) ClientBuilder
	Tls(certFilePath, keyFilePath, CaFilePath string) ClientBuilder
	UserAgent(ua string) ClientBuilder
	RetryCount(num int) ClientBuilder
	Delay(num float64) ClientBuilder
	ConnTimeout(timeout time.Duration) ClientBuilder
	ReadTimeout(timeout time.Duration) ClientBuilder
	Logger(logger log.Logger) ClientBuilder
	Build() Client
}

type clientBuilder struct {
	host              string
	basicAuthUser     string
	basicAuthPassword string
	tlsCertFilePath   string
	tlsKeyFilePath    string
	tlsCaFilePath     string
	userAgent         string
	retryCount        int
	delay             float64
	connTimeout       time.Duration
	readTimeout       time.Duration
	logger            log.Logger
}

func (cb *clientBuilder) Host(host string) ClientBuilder {
	panic("implement me")
}

func (cb *clientBuilder) Auth(user string, password string) ClientBuilder {
	panic("implement me")
}

func (cb *clientBuilder) Tls(certFilePath, keyFilePath, CaFilePath string) ClientBuilder {
	panic("implement me")
}

func (cb *clientBuilder) UserAgent(ua string) ClientBuilder {
	panic("implement me")
}

func (cb *clientBuilder) RetryCount(num int) ClientBuilder {
	panic("implement me")
}

func (cb *clientBuilder) Delay(num float64) ClientBuilder {
	panic("implement me")
}

func (cb *clientBuilder) ConnTimeout(timeout time.Duration) ClientBuilder {
	panic("implement me")
}

func (cb *clientBuilder) ReadTimeout(timeout time.Duration) ClientBuilder {
	panic("implement me")
}

func (cb *clientBuilder) Logger(logger log.Logger) ClientBuilder {
	panic("implement me")
}

func (cb *clientBuilder) Build() Client {
	panic("implement me")
}


func NewClientBuilder() ClientBuilder {
	return &clientBuilder{}
}
