package utils

import "net/http"

type MockRoundTripper struct {
	RtFunc func(r *http.Request) (*http.Response, error)
}

func (rt *MockRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	return rt.RtFunc(r)
}
