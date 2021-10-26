package ghttp

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
)


type GoatClient struct {
	builder *clientBuilder
	client  *http.Client
}

func (g *GoatClient) doRequest(method string, path string, headers map[string]string, body *bytes.Buffer) (*Response, error) {
	g.builder.logger.Printf("making request to: '%s'", g.builder.host+path)
	req, err := http.NewRequest(method, g.builder.host+path, body)
	if err != nil {
		g.builder.logger.Fatal(err)
	}
	headers["User-Agent"] = g.builder.userAgent
	req = setHeaders(req, headers)
	resp, err := g.client.Do(req)
	if err != nil {
		g.builder.logger.Fatalf("error: %s", err)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	g.builder.logger.Printf("status code: '%d'", resp.StatusCode)
	return &Response{
		Status:     resp.Status,
		StatusCode: resp.StatusCode,
		Headers:    resp.Header,
		Body:       respBody,
	}, nil
}

func (g *GoatClient) DoRequest(method string, path string, headers map[string]string, body interface{}) (*Response, error) {
	delay := g.builder.delay
	if headers == nil {
		headers = map[string]string{}
	}
	var err error
	for attempt := 0; attempt <= g.builder.retryCount; attempt++ {
		resp, err := g.doRequest(method, path, headers, toByteBuffer(headers, body))
		g.builder.logger.Printf(resp.Status)
		g.builder.logger.Printf("%d", resp.StatusCode)
		if err == nil {
			return resp, nil
		}
		delay *= 2
		time.Sleep(time.Second * time.Duration(delay))
		g.builder.logger.Printf("attempt: '%d'", attempt)
	}
	return nil, err
}

func (g *GoatClient) DoGet(path string, headers map[string]string) (*Response, error) {
	return g.DoRequest(http.MethodGet, path, headers, nil)
}

func (g *GoatClient) DoDelete(path string, headers map[string]string, body interface{}) (*Response, error) {
	return g.DoRequest(http.MethodDelete, path, headers, body)
}

func (g *GoatClient) DoPatch(path string, headers map[string]string, body interface{}) (*Response, error) {
	return g.DoRequest(http.MethodPatch, path, headers, body)
}

func (g *GoatClient) DoPost(path string, headers map[string]string, body interface{}) (*Response, error) {
	return g.DoRequest(http.MethodPost, path, headers, body)
}

func (g *GoatClient) DoPut(path string, headers map[string]string, body interface{}) (*Response, error) {
	return g.DoRequest(http.MethodPut, path, headers, body)
}
