package ghttp

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type Client interface {
	Get(url string, headers map[string]string) (*Response, error)
}

type HttpClient struct {
	builder *clientBuilder
	client  *http.Client
}

func (c *HttpClient) doRequest(method string, path string, headers map[string]string, body *bytes.Buffer) (*Response, error) {
	req, err := http.NewRequest(method, path, body)
	if err != nil {
		c.builder.logger.Fatal(err)
	}
	headers["User-Agent"] = c.builder.userAgent
	req = setHeaders(req, headers)
	resp, err := c.client.Do(req)
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	c.builder.logger.Printf("status code: '%d'", resp.StatusCode)
	return &Response{
		Status:     resp.Status,
		StatusCode: resp.StatusCode,
		Headers:    resp.Header,
		Body:       respBody,
	}, nil
}

func (c *HttpClient) DoRequest(method string, path string, headers map[string]string, body interface{}) (*Response, error) {
	delay := c.builder.delay
	if headers == nil {

	}
}

func (c *HttpClient) Get(url string, headers map[string]string) (*Response, error) {
	panic("implement me")
}
