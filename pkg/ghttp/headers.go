package ghttp

import "net/http"

func setHeaders(r *http.Request, headers map[string]string) *http.Request {
	for headerName, headerValue := range headers {
		r.Header.Set(headerName, headerValue)
	}
	return r
}
