package ghttp

import (
	"log"
	"net/http"
	"time"
)

type clientBuilder struct {
	host              string
	basicAuthUser     string
	basicAuthPassword string
	tlsCertFilePath   string
	tlsKeyFilePath    string
	tlsCaFilePath     string
	userAgent         string
	retryCount        int
	retryOnErrors     []int
	delay             float64
	connTimeout       time.Duration
	readTimeout       time.Duration
	logger            *log.Logger
}

func (b *clientBuilder) Host(host string) *clientBuilder {
	b.host = host
	return b
}

func (b *clientBuilder) Auth(user string, password string) *clientBuilder {
	b.basicAuthUser = user
	b.basicAuthPassword = password
	return b
}

func (b *clientBuilder) Tls(certFilePath, keyFilePath, CaFilePath string) *clientBuilder {
	b.tlsCertFilePath = certFilePath
	b.tlsKeyFilePath = keyFilePath
	b.tlsCaFilePath = CaFilePath
	return b
}

func (b *clientBuilder) UserAgent(userAgent string) *clientBuilder {
	b.userAgent = userAgent
	return b
}

func (b *clientBuilder) Retry(count int, onErrors []int) *clientBuilder {
	b.retryCount = count
	b.retryOnErrors = onErrors
	return b
}

func (b *clientBuilder) Delay(delay float64) *clientBuilder {
	b.delay = delay
	return b
}

func (b *clientBuilder) ConnTimeout(timeout float64) *clientBuilder {
	b.connTimeout = time.Duration(timeout) * time.Second
	return b
}

func (b *clientBuilder) ReadTimeout(timeout float64) *clientBuilder {
	b.readTimeout = time.Duration(timeout) * time.Second
	return b
}

func (b *clientBuilder) Logger(logger *log.Logger) *clientBuilder {
	b.logger = logger
	return b
}

func (b *clientBuilder) Build() *GoatClient {
	return &GoatClient{
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
}

func NewClientBuilder() *clientBuilder {
	return &clientBuilder{}
}
