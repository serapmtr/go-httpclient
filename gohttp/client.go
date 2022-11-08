package gohttp

import (
	"net/http"
	"sync"

	"github.com/serapmtr/go-httpclient/core"
)

type httpClient struct {
	builder *clientBuilder

	client     *http.Client // * default value is nil
	clientOnce sync.Once
}

type Client interface {
	// with '...' we can call the function without headers
	// or with multiple headers
	// httpClient.Get(url) OR httpClient.Get(url, nil, nil ,nil)
	Get(url string, headers ...http.Header) (*core.Response, error)
	Put(url string, body interface{}, headers ...http.Header) (*core.Response, error)
	Post(url string, body interface{}, headers ...http.Header) (*core.Response, error)
	Patch(url string, body interface{}, headers ...http.Header) (*core.Response, error)
	Delete(url string, headers ...http.Header) (*core.Response, error)
}

// func (c *HttpClient)headers Get() {} // method

// func Get() {} // package functio

func (c *httpClient) Get(url string, headers ...http.Header) (*core.Response, error) {
	return c.do(http.MethodGet, url, getHeaders(headers...), nil)
}

func (c *httpClient) Post(url string, body interface{}, headers ...http.Header) (*core.Response, error) {

	return c.do(http.MethodPost, url, getHeaders(headers...), body)
}

func (c *httpClient) Put(url string, body interface{}, headers ...http.Header) (*core.Response, error) {
	return c.do(http.MethodPut, url, getHeaders(headers...), nil)
}

func (c *httpClient) Patch(url string, body interface{}, headers ...http.Header) (*core.Response, error) {
	return c.do(http.MethodPatch, url, getHeaders(headers...), nil)
}

func (c *httpClient) Delete(url string, headers ...http.Header) (*core.Response, error) {
	return c.do(http.MethodDelete, url, getHeaders(headers...), nil)
}
