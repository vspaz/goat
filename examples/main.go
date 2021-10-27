package main

import (
	"fmt"
	"github.com/vspaz/goat/pkg/ghttp"
	"log"
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
		Host("https://httpbin.org").
		Auth("user", "user-password").
		Tls("", "", ""). // e.g. cert.pem, key.pem, ca.crt
		UserAgent("goat").
		Retry(3, nil).
		Delay(0.5).
		ConnTimeout(5).
		ReadTimeout(10).
		Logger(log.Default()).
		Build()
	resp, err := client.DoGet("/get", nil)  // queries https://httpbin.org/get"
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp.StatusCode)
	log.Println(resp.ToString())

	decodedBody := HttpBinGetResponse{}
	resp.FromJson(&decodedBody)
	fmt.Println(decodedBody.Headers.UserAgent)
}
