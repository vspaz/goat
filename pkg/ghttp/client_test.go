package ghttp

import (
	"encoding/json"
	"log"

	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

const (
	userAgent    = "goat"
	contentType  = "application/json"
	testEndpoint = "/endpoint"
)

func TestHttpClient_DoGet(t *testing.T) {
	t.Parallel()
	encodedBody, _ := json.Marshal(map[string]string{"foo": "bar"})
	server := httptest.NewServer(
		http.HandlerFunc(
			func(writer http.ResponseWriter, request *http.Request) {
				assert.Equal(t, request.URL.String(), testEndpoint)
				assert.Equal(t, request.Header["User-Agent"][0], userAgent)
				assert.Equal(t, request.Header["Content-Type"][0], contentType)
				writer.WriteHeader(http.StatusOK)
				if _, err := writer.Write(encodedBody); err != nil {
					t.Fatal(err)
				}
			},
		),
	)
	defer server.Close()

	client := NewClientBuilder().
		Host(server.URL).
		UserAgent(userAgent).
		Auth("user", "pass").
		RetryCount(3).
		ConnTimeout(5 * time.Second).
		Delay(0.5).
		ReadTimeout(10 * time.Second).
		Logger(log.Default()).
		Build()
	resp, _ := client.DoGet(testEndpoint, map[string]string{"Content-Type": contentType})
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "{\"foo\":\"bar\"}", string(resp.Body))
}
