package client

import (
	"encoding/json"
	"net/http"
)

type Client struct {
	http *http.Client
}

func NewClient(h *http.Client) *Client {
	return &Client{
		http: h,
	}
}

func (c *Client) GetReq(url string, target interface{}) error {
	resp, err := c.http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}
