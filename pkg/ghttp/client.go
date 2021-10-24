package ghttp

import "net/http"

type Client interface {
	Get(url string, headers map[string]string) (*Response, error)
}

type HttpClient struct {
	builder *clientBuilder
	client  *http.Client
}

func (c *HttpClient) Get(url string, headers map[string]string) (*Response, error) {
	panic("implement me")
}

