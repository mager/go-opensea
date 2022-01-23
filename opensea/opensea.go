package opensea

import (
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	apiKey      string
	client      *http.Client
	baseURL     *url.URL
	rateLimitMs time.Duration
}

func NewClient(apiKey string) *Client {
	return &Client{
		apiKey:  apiKey,
		baseURL: &url.URL{Scheme: "https", Host: "api.opensea.io", Path: "/"},
		client: &http.Client{
			Timeout: time.Second * 10,
		},
		rateLimitMs: time.Millisecond * 250,
	}
}
