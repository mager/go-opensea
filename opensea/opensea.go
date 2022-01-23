package opensea

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"go.uber.org/zap"
)

type Client struct {
	apiKey      string
	client      *http.Client
	baseURL     *url.URL
	limitAssets int
	logger      *zap.SugaredLogger
	rateLimitMs time.Duration
}

func NewClient(apiKey string) *Client {
	return &Client{
		apiKey:  apiKey,
		baseURL: &url.URL{Scheme: "https", Host: "api.opensea.io", Path: "/"},
		client: &http.Client{
			Timeout: time.Second * 10,
		},
		limitAssets: 50,
		logger:      zap.NewExample().Sugar(),
		rateLimitMs: time.Millisecond * 250,
	}
}

func (c *Client) NewRequest(method string, u *url.URL, body interface{}) (*http.Request, error) {
	var (
		err error
		buf []byte
	)

	if body != nil {
		buf, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), bytes.NewBuffer(buf))
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-API-KEY", c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}
