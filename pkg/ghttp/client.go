package ghttp

type Client interface {
	Get(url string, headers map[string]string) (*Response, error)
}