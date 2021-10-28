# goat

configurable HTTP Go client that supports:

* tls,
* basic auth
* delayed retries on specific errors
* timeouts
* extra helpers
* etc.

easily configurable http client built on top of a standard http package w/o using any third parties.
the idea is similar to https://github.com/vspaz/pyclient

```go
package main

import (
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
        Host("https://httpbin.org").  // optional 
        Auth("user", "user-password").  // optional 
        Tls("cert.pem", "cert.pem", "ca.crt"). // optional 
        UserAgent("goat"). // optional 
        Retry(3, []int{500, 503}). // optional 
        Delay(0.5). // optional
		ResponseTimeout(10.0).  // optional
		HeadersReadTimeout(2.0).  // optional
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
	
	
	// or just run with default parameters
	client = ghttp.NewClientBuilder().Build()
	resp, err = client.DoGet("https://httpbin.org", nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp.IsOk())
	log.Println(resp.ToString())
}
```
