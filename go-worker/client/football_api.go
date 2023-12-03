package client

import (
	"fmt"
	"os"

	"github.com/cblokkeel/footies/types"
)

type FootballAPIClient struct {
	url string

	*Client
}

type MatchResponse struct {
	Response []*types.Match
}

func NewFootballAPIClient(baseClient *Client, url string) *FootballAPIClient {
	return &FootballAPIClient{
		url:    url,
		Client: baseClient,
	}
}

func (c *FootballAPIClient) getHeaders() map[string]string {
	return map[string]string{
		"X-RapidAPI-Key": os.Getenv("FOOTBALL_API_KEY"), // ENV
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
		// TODO better handling error
		return nil, fmt.Errorf("Error while fetching matches")
	}
	return resp.Response, nil
}

func (c *FootballAPIClient) GetMatchByID(ID string) (*types.Match, error) {
	var resp MatchResponse
	url := c.AddQueryParams(c.url, map[string]string{
		"id": ID,
	})
	if err := c.GetReq(url, &resp, c.getHeaders()); err != nil {
		// TODO better handling error
		return nil, fmt.Errorf("Error while fetching match %s", ID)
	}
	if len(resp.Response) == 0 {
		return nil, fmt.Errorf("No match found")
	}
	return resp.Response[0], nil
}
