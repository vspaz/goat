package ghttp

import "net/http"

type Auth struct {
	transport *http.Transport
	username  string
	password  string
}

func (auth *Auth) RoundTrip(request *http.Request) (*http.Response, error) {
	request.SetBasicAuth(auth.username, auth.password)
	return auth.transport.RoundTrip(request)
}
