package ghttp

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
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

func startServer(t *testing.T) *httptest.Server {
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

func TestHttpClientDoGetHeaders(t *testing.T) {
	server := startServer(t)
	defer server.Close()
	client := NewClientBuilder().
		Host(server.URL).
		UserAgent(userAgent).
		Build()
	resp, err := client.DoGet(testEndpoint, contentTypeJson)
	log.Print(err)
	assert.True(t, resp.IsOk())
	assert.Equal(t, "{\"foo\":\"bar\"}", resp.ToString())

	deserializedBody := make(map[string]string)
	assert.Nil(t, resp.FromJson(&deserializedBody))
	assert.Equal(t, map[string]string{"foo": "bar"}, deserializedBody)
}

var retryCount int

func startServerWithRetries(t *testing.T) *httptest.Server {
	return httptest.NewServer(
		http.HandlerFunc(
			func(writer http.ResponseWriter, request *http.Request) {
				if retryCount < 2 {
					writer.WriteHeader(http.StatusInternalServerError)
					retryCount++
				} else {
					writer.WriteHeader(http.StatusOK)
				}
			},
		),
	)
}

func TestGoatClientRetries(t *testing.T) {
	server := startServerWithRetries(t)
	defer server.Close()
	client := NewClientBuilder().
		Host(server.URL).
		Retry(2, []int{500}).
		Build()
	resp, _ := client.DoGet(testEndpoint, nil)
	assert.True(t, resp.IsOk())
}

func startBasicAuthServer(t *testing.T) *httptest.Server {
	return httptest.NewServer(
		http.HandlerFunc(
			func(writer http.ResponseWriter, request *http.Request) {
				user, password, ok := request.BasicAuth()
				assert.True(t, ok)
				assert.Equal(t, user, "user")
				assert.Equal(t, password, "pass")
				writer.WriteHeader(http.StatusOK)
			},
		),
	)
}

func TestGoatClientBasicAuth(t *testing.T) {
	server := startBasicAuthServer(t)
	defer server.Close()
	client := NewClientBuilder().
		Host(server.URL).
		Auth("user", "pass").
		Build()
	resp, _ := client.DoGet(testEndpoint, contentTypeJson)
	assert.True(t, resp.IsOk())
}

func startServerWithTimeouts() *httptest.Server {
	return httptest.NewServer(
		http.HandlerFunc(
			func(writer http.ResponseWriter, request *http.Request) {
				time.Sleep(2 * time.Second)
				writer.WriteHeader(http.StatusOK)
			},
		),
	)
}

func TestGoatClientResponseTimeout(t *testing.T) {
	server := startServerWithTimeouts()
	defer server.Close()
	client := NewClientBuilder().
		Host(server.URL).
		ResponseTimeout(1).
		Build()
	resp, err := client.DoGet(testEndpoint, contentTypeJson)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "Client.Timeout exceeded")
}
