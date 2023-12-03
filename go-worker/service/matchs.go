package service

import (
	"context"
	"fmt"
	"time"

	"github.com/cblokkeel/footies/cache"
	"github.com/cblokkeel/footies/client"
	"github.com/cblokkeel/footies/dto"
	"github.com/cblokkeel/footies/types"
)

type MatchService struct {
	cache  cache.Cacher
	client *client.FootballAPIClient
}

func NewMatchService(cache cache.Cacher, client *client.FootballAPIClient) *MatchService {
	return &MatchService{
		cache,
		client,
	}
}

func (s *MatchService) GetMatchByLeague(ctx context.Context, date string, leagueID string, season string) ([]*dto.MatchDTO, error) {
	cacheKey := fmt.Sprintf("matchs_%s_%s", leagueID, date)
	var cachedMatchs []*dto.MatchDTO
	if exists := s.cache.Get(ctx, cacheKey, &cachedMatchs); exists {
		return cachedMatchs, nil
	}

	rawMatchs, err := s.client.GetMatchsByLeague(date, season, leagueID)
	if err != nil {
		return nil, err
	}
	matchs := types.ToMatchDTOs(rawMatchs)
	s.cache.Set(ctx, cacheKey, matchs, time.Minute)
	return matchs, nil
}

func (s *MatchService) GetMatchByID(ctx context.Context, matchID string) (*dto.MatchDTO, error) {
	cacheKey := fmt.Sprintf("match_%s", matchID)
	var cachedMatch *dto.MatchDTO
	if exists := s.cache.Get(ctx, cacheKey, &cachedMatch); exists {
		return cachedMatch, nil
	}

	rawMatch, err := s.client.GetMatchByID(matchID)
	if err != nil {
		fmt.Println("ICI")
		return nil, err
	}
	match := rawMatch.ToDTO()
	s.cache.Set(ctx, cacheKey, match, time.Minute)
	return match, nil
}
