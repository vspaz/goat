package ghttp

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Client interface {
	DoGet(path string, headers map[string]string) (*Response, error)
}

type HttpClient struct {
	builder *clientBuilder
	client  *http.Client
}

func (c *HttpClient) doRequest(method string, path string, headers map[string]string, body *bytes.Buffer) (*Response, error) {
	c.builder.logger.Printf("making request to: '%s'", c.builder.host+ "/" + path)
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
	resp, err := c.DoRequest(http.MethodGet, path, headers, nil)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return resp, nil
}
