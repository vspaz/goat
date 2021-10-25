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

func TestHttpClient_DoGet(t *testing.T) {
	encodedBody, _ := json.Marshal(map[string]string{"foo": "bar"})
	server := httptest.NewServer(
		http.HandlerFunc(
			func(writer http.ResponseWriter, request *http.Request) {
				assert.Equal(t, request.URL.String(), "/foo")
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
		UserAgent("goat").
		Auth("user", "pass").
		RetryCount(3).
		ConnTimeout(5 * time.Second).
		Delay(0.5).
		ReadTimeout(10 * time.Second).
		Logger(log.Default()).
		Build()
	resp, _ := client.DoGet("/foo", nil)
	assert.Equal(t, resp.StatusCode, http.StatusOK)
}
