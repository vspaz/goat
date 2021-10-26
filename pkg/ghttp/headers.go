package ghttp

import "net/http"

var (
	contentTypeJson = map[string]string{"Content-Type": "application/json"}
)

func setHeaders(r *http.Request, headers map[string]string) *http.Request {
	for headerName, headerValue := range headers {
		r.Header.Set(headerName, headerValue)
	}
	return r
}

func isJson(headers map[string]string) bool {
	contentType, hasKey := headers["Content-Type"]
	return hasKey && contentType == "application/json"
}
