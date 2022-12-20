package main

import (
	"github.com/vspaz/goat/pkg/ghttp"
	"github.com/vspaz/simplelogger/pkg/logging"
)

type HttpBinGetResponse struct {
	Args struct {
	} `json:"args"`
	Headers struct {
		AcceptEncoding string `json:"Accept-Encoding"`
		Authorization  string `json:"Authorization"`
		Host           string `json:"Host"`
		UserAgent      string `json:"User-Agent"`
		XAmznTraceId   string `json:"X-Amzn-Trace-Id"`
	} `json:"headers"`
	Origin string `json:"origin"`
	Url    string `json:"url"`
}

func main() {
	client := ghttp.NewClientBuilder().
		WithHost("https://httpbin.org").
		WithAuth("user", "user-password").
		WithTls("", "", ""). // e.g. cert.pem, key.pem, ca.crt
		WithUserAgent("goat").
		WithRetry(3, nil).
		WithDelay(0.5).
		WithResponseTimeout(10.0).
		WithHeadersReadTimeout(2.0).
		WithLogLevel("info").
		Build()
	resp, err := client.DoGet("/get", nil, nil) // queries https://httpbin.org/get"
	logger := logging.GetTextLogger("info").Logger
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info(resp.StatusCode)
	logger.Info(resp.ToString())

	decodedBody := HttpBinGetResponse{}
	resp.FromJson(&decodedBody)
	logger.Info(decodedBody.Headers.UserAgent)

	// or just run as is with default parameters
	client = ghttp.NewClientBuilder().Build()
	resp, err = client.DoGet("https://httpbin.org/get", nil, nil)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info(resp.StatusCode)
	logger.Info(resp.ToString())
}
