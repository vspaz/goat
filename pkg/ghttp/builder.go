package ghttp

import (
	"github.com/sirupsen/logrus"
	"github.com/vspaz/simplelogger/pkg/logging"
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
	retryCount            uint8
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

func (b *clientBuilder) WithHost(host string) *clientBuilder {
	b.host = host
	return b
}

func (b *clientBuilder) WithAuth(user string, password string) *clientBuilder {
	b.basicAuthUser = user
	b.basicAuthPassword = password
	return b
}

func (b *clientBuilder) WithTls(certFilePath, keyFilePath, CaFilePath string) *clientBuilder {
	b.tlsCertFilePath = certFilePath
	b.tlsKeyFilePath = keyFilePath
	b.tlsCaFilePath = CaFilePath
	return b
}

func (b *clientBuilder) WithUserAgent(userAgent string) *clientBuilder {
	b.userAgent = userAgent
	return b
}

func (b *clientBuilder) WithRetry(count uint8, onErrors []int) *clientBuilder {
	b.retryCount = count
	b.retryOnErrors = onErrors
	return b
}

func (b *clientBuilder) WithDelay(delay float64) *clientBuilder {
	b.delay = time.Duration(delay)
	return b
}

func (b *clientBuilder) WithResponseTimeout(timeout float64) *clientBuilder {
	b.responseTimeout = time.Duration(timeout) * time.Second
	return b
}

func (b *clientBuilder) WithConnectionTimeout(timeout float64) *clientBuilder {
	b.connectionTimeout = time.Duration(timeout) * time.Second
	return b
}

func (b *clientBuilder) WithKeepAlive(timeout float64) *clientBuilder {
	b.keepAlive = time.Duration(timeout) * time.Second
	return b
}

func (b *clientBuilder) WithIdleConnectionTimeout(timeout float64) *clientBuilder {
	b.idleConnectionTimeout = time.Duration(timeout) * time.Second
	return b
}

func (b *clientBuilder) WithHeadersReadTimeout(timeout float64) *clientBuilder {
	b.headersReadTimeout = time.Duration(timeout) * time.Second
	return b
}

func (b *clientBuilder) WithTlsHandshakeTimeout(timeout float64) *clientBuilder {
	b.tlsHandshakeTimeout = time.Duration(timeout) * time.Second
	return b
}

func (b *clientBuilder) WithLogLevel(logLevel string) *clientBuilder {
	b.logger = logging.GetTextLogger(logLevel).Logger
	return b
}

func setDefaults(b *clientBuilder) *clientBuilder {
	if b.logger == nil {
		b.logger = logging.GetTextLogger("info").Logger
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

	b.retryCount++

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
