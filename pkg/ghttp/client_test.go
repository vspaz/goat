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

func assertRequest(t *testing.T, request *http.Request) {
	assert.Equal(t, request.URL.String(), testEndpoint)
	assert.Equal(t, request.Header["User-Agent"][0], userAgent)
	assert.Equal(t, request.Header["Content-Type"][0], contentType)
}

func createDefaultClient(url string) *GoatClient {
	return NewClientBuilder().
		Host(url).
		UserAgent(userAgent).
		Auth("user", "pass").
		Retry(1, nil).
		ConnTimeout(5).
		Delay(0.5).
		ReadTimeout(10).
		Logger(log.Default()).
		Build()
}

func getTestServer(t *testing.T) *httptest.Server {
	return httptest.NewServer(
		http.HandlerFunc(
			func(writer http.ResponseWriter, request *http.Request) {
				assertRequest(t, request)
				writer.WriteHeader(http.StatusOK)
				encodedBody, _ := json.Marshal(map[string]string{"foo": "bar"})
				if _, err := writer.Write(encodedBody); err != nil {
					t.Fatal(err)
				}
			},
		),
	)
}

func TestHttpClient_DoGet(t *testing.T) {
	t.Parallel()
	server := getTestServer(t)
	defer server.Close()

	client := createDefaultClient(server.URL)
	resp, _ := client.DoGet(testEndpoint, contentTypeJson)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "{\"foo\":\"bar\"}", resp.ToString())

	deserializedBody := make(map[string]string)
	assert.Nil(t, resp.FromJson(&deserializedBody))
	assert.Equal(t, map[string]string{"foo": "bar"}, deserializedBody)
}
