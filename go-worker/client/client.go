package client

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type Client struct {
	http *http.Client
}

func NewClient(h *http.Client) *Client {
	return &Client{
		http: h,
	}
}

func (c *Client) GetReq(url string, target interface{}, headers map[string]string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

func (c *Client) AddQueryParams(baseUrl string, queryParams map[string]string) string {
	u, err := url.Parse(baseUrl)
	if err != nil {
		return ""
	}
	q := u.Query()
	for key, value := range queryParams {
		q.Add(key, value)
	}
	u.RawQuery = q.Encode()
	return u.String()
}
