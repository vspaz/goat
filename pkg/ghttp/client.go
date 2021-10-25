package ghttp

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
)

type Client interface {
	DoRequest(method string, path string, headers map[string]string, body interface{}) (*Response, error)
	DoGet(path string, headers map[string]string) (*Response, error)
	DoDelete(path string, headers map[string]string, body interface{}) (*Response, error)
	DoPatch(path string, headers map[string]string, body interface{}) (*Response, error)
	DoPost(path string, headers map[string]string, body interface{}) (*Response, error)
	DoPut(path string, headers map[string]string, body interface{}) (*Response, error)
}

type HttpClient struct {
	builder *clientBuilder
	client  *http.Client
}

func (c *HttpClient) doRequest(method string, path string, headers map[string]string, body *bytes.Buffer) (*Response, error) {
	c.builder.logger.Printf("making request to: '%s'", c.builder.host+"/"+path)
	req, err := http.NewRequest(method, c.builder.host+"/"+path, body)
	if err != nil {
		c.builder.logger.Fatal(err)
	}
	headers["User-Agent"] = c.builder.userAgent
	req = setHeaders(req, headers)
	resp, err := c.client.Do(req)
	if err != nil {
		c.builder.logger.Printf("error: %s", err)
	}
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
		headers = map[string]string{}
	}
	var err error
	for attempt := 0; attempt <= c.builder.retryCount; attempt++ {
		resp, err := c.doRequest(method, path, headers, toByteBuffer(headers, body))
		c.builder.logger.Printf(resp.Status)
		c.builder.logger.Printf("%d", resp.StatusCode)
		if err == nil {
			return resp, nil
		}
		delay *= 2
		time.Sleep(time.Second * time.Duration(delay))
		c.builder.logger.Printf("attempt: '%d'", attempt)
	}
	return nil, err
}

func (c *HttpClient) DoGet(path string, headers map[string]string) (*Response, error) {
	return c.DoRequest(http.MethodGet, path, headers, nil)
}

func (c *HttpClient) DoDelete(path string, headers map[string]string, body interface{}) (*Response, error) {
	return c.DoRequest(http.MethodDelete, path, headers, body)
}

func (c *HttpClient) DoPatch(path string, headers map[string]string, body interface{}) (*Response, error) {
	return c.DoRequest(http.MethodPatch, path, headers, body)
}
