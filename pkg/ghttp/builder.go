package ghttp

import "time"

type ClientBuilder interface {
	Host(host string) ClientBuilder
	Auth(user string, password string) ClientBuilder
	Tls(certFilePath, keyFilePath, CaFilePath string) ClientBuilder
	UserAgent(ua string) ClientBuilder
	Retry(num int) ClientBuilder
	Delay(num float64) ClientBuilder
	ConnTimeout(timeout time.Duration) ClientBuilder
	ReadTimeout(timeout time.Duration) ClientBuilder
}
