package ghttp

import (
	"log"
	"net/http"
	"time"
)

type ClientBuilder interface {
	Host(host string) ClientBuilder
	Auth(user string, password string) ClientBuilder
	Tls(certFilePath, keyFilePath, CaFilePath string) ClientBuilder
	UserAgent(userAgent string) ClientBuilder
	RetryCount(retryCount int) ClientBuilder
	Delay(delay float64) ClientBuilder
	ConnTimeout(timeout time.Duration) ClientBuilder
	ReadTimeout(timeout time.Duration) ClientBuilder
	Logger(logger *log.Logger) ClientBuilder
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
	logger            *log.Logger
}

func (b *clientBuilder) Host(host string) ClientBuilder {
	b.host = host
	return b
}

func (b *clientBuilder) Auth(user string, password string) ClientBuilder {
	b.basicAuthUser = user
	b.basicAuthPassword = password
	return b
}

func (b *clientBuilder) Tls(certFilePath, keyFilePath, CaFilePath string) ClientBuilder {
	b.tlsCertFilePath = certFilePath
	b.tlsKeyFilePath = keyFilePath
	b.tlsCaFilePath = CaFilePath
	return b
}

func (b *clientBuilder) UserAgent(userAgent string) ClientBuilder {
	b.userAgent = userAgent
	return b
}

func (b *clientBuilder) RetryCount(retryCount int) ClientBuilder {
	b.retryCount = retryCount
	return b
}

func (b *clientBuilder) Delay(delay float64) ClientBuilder {
	b.delay = delay
	return b
}

func (b *clientBuilder) ConnTimeout(timeout time.Duration) ClientBuilder {
	b.connTimeout = timeout
	return b
}

func (b *clientBuilder) ReadTimeout(timeout time.Duration) ClientBuilder {
	b.readTimeout = timeout
	return b
}

func (b *clientBuilder) Logger(logger *log.Logger) ClientBuilder {
	b.logger = logger
	return b
}

func (b *clientBuilder) Build() Client {
	client := &HttpClient{
		builder: b,
		client: &http.Client{
			Timeout: b.connTimeout,
			Transport: &Auth{
				&http.Transport{
					ResponseHeaderTimeout: b.readTimeout,
					TLSClientConfig: createTlsConfig(
						b.tlsCertFilePath,
						b.tlsKeyFilePath,
						b.tlsCaFilePath,
					),
				},
				b.basicAuthUser,
				b.basicAuthPassword,
			},
		},
	}
	return &client
}

func NewClientBuilder() ClientBuilder {
	return &clientBuilder{}
}
