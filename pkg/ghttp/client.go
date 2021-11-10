package ghttp

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type GoatClient struct {
	builder *clientBuilder
	client  *http.Client
	mutex   *sync.Mutex
}

func setRequestQueryParams(req *http.Request, params map[string]string) {
	if params == nil {
		return
	}
	q := req.URL.Query()
	for key, value := range params {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()
}

func (g *GoatClient) doRequest(method string, path string, headers map[string]string, body *bytes.Buffer, params map[string]string) (*Response, error) {
	g.builder.logger.Infof("making request to: '%s'", g.builder.host+path)
	req, err := http.NewRequest(method, g.builder.host+path, body)
	if err != nil {
		g.builder.logger.Fatal(err)
	}
	setRequestQueryParams(req, params)
	req = setHeaders(req, headers)
	resp, err := g.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			g.builder.logger.Fatalf("failed to close body %s", err)
		}
	}(resp.Body)
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		g.builder.logger.Error(err)
		return nil, err
	}
	g.builder.logger.Debugf("status code: '%d'", resp.StatusCode)
	return &Response{
		Status:     resp.Status,
		StatusCode: resp.StatusCode,
		Headers:    resp.Header,
		Body:       respBody,
	}, nil
}

func isRetryOnError(statusCode int, errorStatusCodes []int) bool {
	if errorStatusCodes == nil || len(errorStatusCodes) == 0 {
		return false
	}
	for _, errorCode := range errorStatusCodes {
		if statusCode == errorCode {
			return true
		}
	}
	return false
}

func (g *GoatClient) DoRequest(method string, path string, headers, params map[string]string, body interface{}) (*Response, error) {
	delay := g.builder.delay
	if headers == nil {
		headers = map[string]string{}
	}
	headers["User-Agent"] = g.builder.userAgent
	var err error
	var resp *Response
	for attempt := 0; attempt <= g.builder.retryCount; attempt++ {
		resp, err = g.doRequest(method, path, headers, toByteBuffer(headers, body), params)
		if err == nil && !isRetryOnError(resp.StatusCode, g.builder.retryOnErrors) {
			return resp, nil
		}
		delay *= 2
		time.Sleep(time.Second * delay)
		g.builder.logger.Warnf("attempt: '%d'", attempt)
	}
	return nil, err
}

func (g *GoatClient) DoGet(path string, headers, params map[string]string) (*Response, error) {
	return g.DoRequest(http.MethodGet, path, headers, params, nil)
}

func (g *GoatClient) DoDelete(path string, headers, params map[string]string, body interface{}) (*Response, error) {
	return g.DoRequest(http.MethodDelete, path, headers, params, body)
}

func (g *GoatClient) DoPatch(path string, headers, params map[string]string, body interface{}) (*Response, error) {
	return g.DoRequest(http.MethodPatch, path, headers, params, body)
}

func (g *GoatClient) DoPost(path string, headers, params map[string]string, body interface{}) (*Response, error) {
	return g.DoRequest(http.MethodPost, path, headers, params, body)
}

func (g *GoatClient) DoPut(path string, headers, params map[string]string, body interface{}) (*Response, error) {
	return g.DoRequest(http.MethodPut, path, headers, params, body)
}

func (g *GoatClient) DoHead(path string, headers, params map[string]string) (*Response, error) {
	return g.DoRequest(http.MethodHead, path, headers, params, nil)
}
