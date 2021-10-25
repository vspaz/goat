package ghttp

import "net/http"

func setHeaders(r *http.Request, headers map[string]string) *http.Request {
	for headerName, headerValue := range headers {
		r.Header.Set(headerName, headerValue)
	}
	return r
}

func isJson(headers map[string]string) bool {
	contentType, _ := headers["Content-Type"]
	return contentType == "application/json"
}

func isPlainTest(headers map[string]string) bool {
	contentType, _ := headers["Content-Type"]
	return contentType == "text/plain" || contentType == "text/html"
}
