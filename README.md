# goat

configurable HTTP Go client that supports:

* tls,
* basic auth
* retries
* timeouts
* extra helpers
* etc.

easily configurable http client build on top of standard http package w/o using any third parties.
the idea is similar to https://github.com/vspaz/pyclient

```go
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
		Host("https://httpbin.org").  // required
		Auth("user", "user-password").  // optional
		Tls("cert.pem", "cert.pem", "ca.crt"). // optional
		UserAgent("goat"). // optional
		RetryCount(3). // optional
		Delay(0.5). // optional
		ConnTimeout(5). // optional
		ReadTimeout(10). // optional
		Logger(log.Default()). // optional
		Build()
	resp, err := client.DoGet("/get", nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp.StatusCode)
	log.Println(resp.ToString())

	deserializedBody := HttpBinGetResponse{}
	resp.FromJson(&deserializedBody)
}
```