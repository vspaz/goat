package ghttp

import (
	"encoding/json"
	"log"

	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	userAgent    = "goat"
	contentType  = "application/json"
	testEndpoint = "/test-endpoint"
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
		RetryCount(1).
		ConnTimeout(5).
		Delay(0.5).
		ReadTimeout(10).
		Logger(log.Default()).
		Build()
	resp, _ := client.DoGet(testEndpoint, contentTypeJson)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "{\"foo\":\"bar\"}", resp.ToString())

	deserializedBody := make(map[string]string)
	assert.Nil(t, resp.FromJson(&deserializedBody))
	assert.Equal(t, map[string]string{"foo": "bar"}, deserializedBody)
}
