package testutils

import "net/http"

type roundTripFunc func(req *http.Request) *http.Response

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func newHTTPClient(fn roundTripFunc) *http.Client { // nolint: interfacer
	return &http.Client{
		Transport: fn,
	}
}

func CreateDummyHTTPClient(handler func(r *http.Request) *http.Response) *http.Client {
	return newHTTPClient(func(req *http.Request) *http.Response {
		return handler(req)
	})
}
