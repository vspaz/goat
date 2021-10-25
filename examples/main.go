package main

import (
	"github.com/vspaz/goat/pkg/ghttp"
	"log"
	"time"
)

func main() {
	client := ghttp.NewClientBuilder().
		Host("http://example.com").
		Auth("user", "user-password").
		Tls("", "", "").
		UserAgent("goat").
		RetryCount(3).
		Delay(0.5).
		ConnTimeout(5 * time.Second).
		ReadTimeout(10 * time.Second).
		Logger(log.Default()).
		Build()
	client.DoGet("foobar", nil)
}
