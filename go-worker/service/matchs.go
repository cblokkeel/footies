package service

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/cblokkeel/footies/cache"
	"github.com/cblokkeel/footies/client"
	"github.com/cblokkeel/footies/constants"
	"github.com/cblokkeel/footies/dto"
	"github.com/cblokkeel/footies/pubsub"
	"github.com/cblokkeel/footies/types"
	"github.com/redis/go-redis/v9"
)

type MatchService struct {
	cache  cache.Cacher
	pubsub pubsub.PubSub[*redis.Message] // not ideal, maybe change this is kind of crappy
	client *client.FootballAPIClient
}

const (
	activeMatchesKey string = "active_matches"
)

func NewMatchService(cache cache.Cacher, pubsub pubsub.PubSub[*redis.Message], client *client.FootballAPIClient) *MatchService {
	return &MatchService{
		cache,
		pubsub,
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
		return nil, err
	}
	match := rawMatch.ToDTO()
	s.cache.Set(ctx, cacheKey, match, time.Minute)
	return match, nil
}

func (s *MatchService) MonitorMatch(ctx context.Context, matchID string) error {
	isMonitored := s.cache.HasSet(ctx, activeMatchesKey, matchID)
	if isMonitored {
		log.Printf("Match %s is already being monitored\n", matchID)
		return nil
	}

	if err := s.cache.AddSet(ctx, activeMatchesKey, matchID); err != nil {
		return fmt.Errorf("error adding match %s to active matches", matchID)
	}

	var lastChrono int
	for {
		matchInfo, err := s.GetMatchByID(ctx, matchID)
		if err != nil {
			return err
		}
		if matchInfo.Status == dto.Finished {
			log.Printf("Match %s has ended", matchID)
			break
		}
		if matchInfo.Elapsed != lastChrono { // Later check other informations
			if err := s.pubsub.Publish(ctx, "match_%s_update", matchInfo.Elapsed); err != nil {
				return fmt.Errorf("error publishing update for match %s", matchID)
			}
			lastChrono = matchInfo.Elapsed
		}
		sleepTime, err := strconv.Atoi(os.Getenv(constants.InBetweenRefreshInformation))
		if err != nil {
			return fmt.Errorf("invalid configuration, %s is not a number", os.Getenv(constants.InBetweenRefreshInformation))
		}
		time.Sleep(time.Duration(sleepTime) * time.Second)
	}

	if err := s.cache.RemoveSet(ctx, activeMatchesKey, matchID); err != nil {
		return fmt.Errorf("could not remove %s to monitored match keys %v", matchID, activeMatchesKey)
	}
	return nil
}
