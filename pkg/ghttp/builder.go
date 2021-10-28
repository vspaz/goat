package ghttp

import (
	"github.com/sirupsen/logrus"
	"net"
	"net/http"
	"time"
)

type clientBuilder struct {
	host                  string
	basicAuthUser         string
	basicAuthPassword     string
	tlsCertFilePath       string
	tlsKeyFilePath        string
	tlsCaFilePath         string
	userAgent             string
	retryCount            int
	retryOnErrors         []int
	delay                 time.Duration
	responseTimeout       time.Duration
	connectionTimeout     time.Duration
	keepAlive             time.Duration
	idleConnectionTimeout time.Duration
	tlsHandshakeTimeout   time.Duration
	headersReadTimeout    time.Duration
	logger                *logrus.Logger
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
	b.delay = time.Duration(delay)
	return b
}

func (b *clientBuilder) ResponseTimeout(timeout float64) *clientBuilder {
	b.responseTimeout = time.Duration(timeout) * time.Second
	return b
}

func (b *clientBuilder) ConnectionTimeout(timeout float64) *clientBuilder {
	b.connectionTimeout = time.Duration(timeout) * time.Second
	return b
}

func (b *clientBuilder) KeepAlive(timeout float64) *clientBuilder {
	b.keepAlive = time.Duration(timeout) * time.Second
	return b
}

func (b *clientBuilder) IdleConnectionTimeout(timeout float64) *clientBuilder {
	b.idleConnectionTimeout = time.Duration(timeout) * time.Second
	return b
}

func (b *clientBuilder) HeadersReadTimeout(timeout float64) *clientBuilder {
	b.headersReadTimeout = time.Duration(timeout) * time.Second
	return b
}

func (b *clientBuilder) TlsHandshakeTimeout(timeout float64) *clientBuilder {
	b.tlsHandshakeTimeout = time.Duration(timeout) * time.Second
	return b
}

func (b *clientBuilder) Logger(logger *logrus.Logger) *clientBuilder {
	b.logger = logger
	return b
}

func setDefaults(b *clientBuilder) *clientBuilder {
	if b.logger == nil {
		b.logger = &logrus.Logger{};
	}

	if b.responseTimeout == 0 {
		b.responseTimeout = time.Duration(5) * time.Second
	}

	if b.connectionTimeout == 0 {
		b.connectionTimeout = time.Duration(5) * time.Second
	}

	if b.keepAlive == 0 {
		b.keepAlive = time.Duration(2*60) * time.Second
	}

	if b.headersReadTimeout == 0 {
		b.headersReadTimeout = time.Duration(5) * time.Second
	}

	if b.idleConnectionTimeout == 0 {
		b.idleConnectionTimeout = time.Duration(2*60) * time.Second
	}

	if b.tlsHandshakeTimeout == 0 {
		b.tlsHandshakeTimeout = time.Duration(10) * time.Second
	}

	if b.delay == 0 {
		b.delay = time.Duration(1)
	}
	return b
}

func (b *clientBuilder) Build() *GoatClient {
	b = setDefaults(b)
	return &GoatClient{
		builder: b,
		client: &http.Client{
			Timeout: b.responseTimeout,
			Transport: &Auth{
				&http.Transport{
					DialContext: (&net.Dialer{
						Timeout:   b.connectionTimeout,
						KeepAlive: b.keepAlive,
					}).DialContext,
					IdleConnTimeout:       b.idleConnectionTimeout,
					ResponseHeaderTimeout: b.headersReadTimeout,
					TLSHandshakeTimeout:   b.tlsHandshakeTimeout,
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
