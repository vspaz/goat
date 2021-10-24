package main

import (
	"github.com/vspaz/goat/pkg/ghttp"
	"log"
)

func main() {
	client := ghttp.NewClientBuilder().
		Host("http://example.com").
		Auth("user", "user-password").
		Tls("", "", "").
		UserAgent("goat").
		RetryCount(3).
		Delay(0.5).
		ConnTimeout(5).
		ReadTimeout(10).
		Logger(log.Default()).
		Build()
	client.Get("fobar", nil)
}
