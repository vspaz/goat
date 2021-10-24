package ghttp

import "net/http"

type Response struct {
	Status     string
	StatusCode int
	Headers    http.Header
	Body       []byte
}
