# goat

Goat, is an HTTP client built on top of a standard Go http package, that is extremely easy to configure; no googling
required. The idea is similar to https://github.com/vspaz/pyclient
it supports out of the box:

* TLS,
* basic auth
* delayed retries on specific errors
* timeouts (connection, read, idle etc.)
* keep-alive connections
* extra helpers
* logging
* etc.

still, if you don't wish any configuration, you can you use it as is with default parameters. Please, see the examples
below:

### Building a client

```go
package main

import (
	"github.com/vspaz/goat/pkg/ghttp"
	"log"
)

func main() {
	client := ghttp.NewClientBuilder().
		WithHost("https://httpbin.org"). // optional 
		WithAuth("user", "user-password"). // optional 
		WithTls("cert.pem", "cert.pem", "ca.crt"). // optional 
		WithTlsHandshakeTimeout(10.0). // optional
		WithUserAgent("goat"). // optional
		WithRetry(3, []int{500, 503}). // optional 
		WithDelay(0.5). // optional
		WithResponseTimeout(10.0). // optional
		WithConnectionTimeout(2.0). // optional 
		WithHeadersReadTimeout(2.0). // optional
		WithLogLevel("info"). // optional 
		Build()

	resp, err := client.DoGet("/get", nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp.IsOk())

	// or just
	client = ghttp.NewClientBuilder().Build()
	resp, err = client.DoGet("https://httpbin.org/get", nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp.IsOk())
}
```

### /GET

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
	client := ghttp.NewClientBuilder().Build()
	resp, err := client.DoGet("https://httpbin.org/get", nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp.IsOk())
	log.Println(resp.ToString())
}
```

### /POST

```go

package main

import (
	"github.com/vspaz/goat/pkg/ghttp"
	"log"
)

func main() {
	client := ghttp.NewClientBuilder().Build()
	resp, err := client.DoPost(
		"https://httpbin.org/post",
		map[string]string{"Content-Type": "application/json"},
		map[string]string{"param": "value"},
		map[string]string{"foo": "bar"})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp.IsOk())
	log.Println(resp.ToString())
}
```