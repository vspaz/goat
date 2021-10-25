package main

import (
	"github.com/vspaz/goat/pkg/ghttp"
	"log"
	"time"
)

func main() {
	client := ghttp.NewClientBuilder().
		Host("https://httpbin.org").
		Auth("user", "user-password").
		Tls("", "", "").
		UserAgent("goat").
		RetryCount(3).
		Delay(0.5).
		ConnTimeout(5 * time.Second).
		ReadTimeout(10 * time.Second).
		Logger(log.Default()).
		Build()
	resp, err := client.DoGet("/get", nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp.StatusCode)
	log.Println(string(resp.Body))
}
