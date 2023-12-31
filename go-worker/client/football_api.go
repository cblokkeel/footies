package client

import (
	"github.com/cblokkeel/footies/types"
)

type FootballAPIClient struct {
	url    string
	apiKey string

	*Client
}

type MatchResponse struct {
	Response []*types.Match `json:"response"`
}

func NewFootballAPIClient(baseClient *Client, url string, apiKey string) *FootballAPIClient {
	return &FootballAPIClient{
		url:    url,
		apiKey: apiKey,
		Client: baseClient,
	}
}

func (c *FootballAPIClient) getHeaders() map[string]string {
	return map[string]string{
		"X-RapidAPI-Key": c.apiKey,
	}
}

func (c *FootballAPIClient) GetMatchsByLeague(date string, season string, leagueID string) ([]*types.Match, error) {
	var resp MatchResponse
	url := c.AddQueryParams(c.url, map[string]string{
		"date":   date,
		"season": season,
		"league": leagueID,
	})
	if err := c.GetReq(url, &resp, c.getHeaders()); err != nil {
		return nil, err
	}
	return resp.Response, nil
}

func (c *FootballAPIClient) GetMatchByID(ID string) (*types.Match, error) {
	var resp MatchResponse
	url := c.AddQueryParams(c.url, map[string]string{
		"id": ID,
	})
	if err := c.GetReq(url, &resp, c.getHeaders()); err != nil {
		return nil, err
	}
	return resp.Response[0], nil
}
