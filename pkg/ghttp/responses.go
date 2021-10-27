package ghttp

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status     string
	StatusCode int
	Headers    http.Header
	Body       []byte
}

func (r *Response) ToBytes() []byte {
	return r.Body
}

func (r *Response) ToString() string {
	return string(r.Body)
}

func (r *Response) FromJson(deserialized interface{}) error {
	return json.Unmarshal(r.ToBytes(), deserialized)
}

func (r *Response) IsOk() bool {
	return r.StatusCode == http.StatusOK
}

func (r *Response) IsAccepted() bool {
	return r.StatusCode == http.StatusAccepted
}
