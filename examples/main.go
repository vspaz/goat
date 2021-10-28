package main

import (
	"github.com/vspaz/goat/pkg/ghttp"
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
	logger := ghttp.ConfigureLogger()
	client := ghttp.NewClientBuilder().
		Host("https://httpbin.org").
		Auth("user", "user-password").
		Tls("", "", ""). // e.g. cert.pem, key.pem, ca.crt
		UserAgent("goat").
		Retry(3, nil).
		Delay(0.5).
		ResponseTimeout(10.0).
		HeadersReadTimeout(2.0).
		Logger(logger).
		Build()
	resp, err := client.DoGet("/get", nil) // queries https://httpbin.org/get"
	if err != nil {
		logger.Fatal(err)
	}
	logger.Println(resp.StatusCode)
	logger.Println(resp.ToString())

	decodedBody := HttpBinGetResponse{}
	resp.FromJson(&decodedBody)
	logger.Println(decodedBody.Headers.UserAgent)

	// or just run as is with default parameters
	client = ghttp.NewClientBuilder().Build()
	resp, err = client.DoGet("https://httpbin.org/get", nil)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Println(resp.StatusCode)
	logger.Println(resp.ToString())
}
