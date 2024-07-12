package apiclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Kinveil/Riot-API-Golang/apiclient/ratelimiter"
	"github.com/Kinveil/Riot-API-Golang/constants/continent"
	"github.com/Kinveil/Riot-API-Golang/constants/league/rank"
	"github.com/Kinveil/Riot-API-Golang/constants/league/tier"
	"github.com/Kinveil/Riot-API-Golang/constants/queue_ranked"
	"github.com/Kinveil/Riot-API-Golang/constants/region"
)

type Client interface {
	// Helper methods to set the API key, usage conservation, and max retries.

	SetUsageConservation(conserveUsage ratelimiter.ConserveUsage)
	SetAPIKey(apiKey string)
	SetMaxRetries(maxRetries int)
	WithContext(ctx context.Context) Client

	// Account API

	GetAccountByPuuid(continent continent.Continent, puuid string) (*Account, error)
	GetAccountByRiotID(continent continent.Continent, gameName, tagLine string) (*Account, error)

	// Champion Mastery API

	GetChampionMasteriesBySummonerID(region region.Region, summonerID string) ([]ChampionMastery, error)
	GetChampionMasteryBySummonerIDAndChampionID(region region.Region, summonerID string, championID int) (*ChampionMastery, error)
	GetChampionMasteriesTopBySummonerID(region region.Region, summonerID string) ([]ChampionMastery, error)
	GetChampionMasteryScoreTotalBySummonerID(region region.Region, summonerID string) (int, error)

	// Champion API

	GetChampionRotations(region region.Region) (*ChampionRotations, error)

	// Clash API

	GetClashPlayersByPuuid(region region.Region, puuid string) (*ClashPlayers, error)
	GetClashPlayersBySummonerID(region region.Region, summonerID string) (*ClashPlayers, error)
	GetClashTeamByID(region region.Region, teamID string) (*ClashTeam, error)
	GetClashTournaments(region region.Region) (*ClashTournaments, error)
	GetClashTournamentByTeamID(region region.Region, teamID string) (*ClashTournament, error)
	GetClashTournamentByID(region region.Region, tournamentID string) (*ClashTournament, error)

	// League Exp API

	GetLeagueExpEntries(region region.Region, q queue_ranked.String, tier tier.String, rank rank.String, page int) ([]LeagueEntry, error)

	// League API

	GetLeagueEntriesChallenger(region region.Region, q queue_ranked.String) (*LeagueList, error)
	GetLeagueEntriesGrandmaster(region region.Region, q queue_ranked.String) (*LeagueList, error)
	GetLeagueEntriesMaster(region region.Region, q queue_ranked.String) (*LeagueList, error)
	GetLeagueEntries(region region.Region, q queue_ranked.String, tier tier.String, rank rank.String, page int) ([]LeagueEntry, error)
	GetLeagueEntriesByID(region region.Region, leagueID string) (*LeagueList, error)
	GetLeagueEntriesBySummonerID(region region.Region, summonerID string) ([]LeagueEntry, error)

	// LOL Challenges API

	GetChallengesConfig(region region.Region) (*ChallengesConfig, error)
	GetChallengesPercentiles(region region.Region) (*ChallengesPercentiles, error)
	GetChallengesConfigByID(region region.Region, challengeID string) (*ChallengesConfig, error)
	GetChallengesLeaderboardsByLevel(region region.Region, challengeID, level string) (*ChallengesLeaderboards, error)
	GetChallengesPercentilesByID(region region.Region, challengeID string) (*ChallengesPercentiles, error)
	GetChallengesPlayerDataByPuuid(region region.Region, puuid string) (*ChallengesPlayerData, error)

	// LOL Status API

	GetStatusPlatformData(region region.Region) (*StatusPlatformData, error)

	// Match API

	GetMatchlist(continent continent.Continent, puuid string, opts *GetMatchlistOptions) (*Matchlist, error)
	GetMatch(continent continent.Continent, matchID string) (*Match, error)
	GetMatchTimeline(continent continent.Continent, matchID string) (*MatchTimeline, error)

	// Spectator API

	GetSpectatorActiveGameByPuuid(region region.Region, summonerID string) (*ActiveGame, error)
	GetSpectatorFeaturedGames(region region.Region) (*FeaturedGames, error)

	// Summoner API

	GetSummonerByRsoPuuid(region region.Region, rsoPuuid string) (*Summoner, error)
	GetSummonerByAccountID(region region.Region, accountID string) (*Summoner, error)
	GetSummonerByPuuid(region region.Region, puuid string) (*Summoner, error)
	GetSummonerBySummonerID(region region.Region, summonerID string) (*Summoner, error)
}

// client is the internal implementation of Client.
type client struct {
	ratelimiter *ratelimiter.RateLimiter
	ctx         context.Context
}

// New returns a Client configured for the given API client and underlying HTTP
// client. The returned Client is threadsafe.
func New(apiKey string) Client {
	requests := make(chan *ratelimiter.APIRequest)

	ratelimiter := ratelimiter.NewRateLimiter(requests, apiKey)
	go ratelimiter.Start()

	return &client{
		ratelimiter: ratelimiter,
		ctx:         context.Background(),
	}
}

func (c *client) WithContext(ctx context.Context) Client {
	return &client{
		ratelimiter: c.ratelimiter,
		ctx:         ctx,
	}
}

func (c *client) SetUsageConservation(conserveUsage ratelimiter.ConserveUsage) {
	c.ratelimiter.SetUsageConservation(conserveUsage)
}

func (c *client) SetAPIKey(apiKey string) {
	c.ratelimiter.SetAPIKey(apiKey)
}

func (c *client) SetMaxRetries(maxRetries int) {
	c.ratelimiter.SetMaxRetries(maxRetries)
}

type HostProvider interface {
	Host() string
	String() string
}

func (c *client) dispatchAndUnmarshal(regionOrContinent HostProvider, method string, relativePath string, parameters url.Values, methodID ratelimiter.MethodID, dest interface{}) (*http.Response, error) {
	var suffix, separator string

	if len(parameters) > 0 {
		suffix = fmt.Sprintf("?%s", parameters.Encode())
	}

	if !strings.HasPrefix(relativePath, "/") {
		separator = "/"
	}

	URL := regionOrContinent.Host() + method + separator + relativePath + suffix

	responseChan := make(chan *http.Response, 1)
	errorChan := make(chan error, 1)
	newRequest := ratelimiter.APIRequest{
		Context:  c.ctx,
		Region:   strings.ToUpper(regionOrContinent.String()),
		MethodID: methodID,
		URL:      URL,
		Response: responseChan,
		Error:    errorChan,
	}

	// Insert the request into the rate limiter
	select {
	case <-c.ctx.Done():
		return nil, c.ctx.Err()
	case c.ratelimiter.Requests <- &newRequest:
	}

	// Wait for the response
	select {
	case <-c.ctx.Done():
		return nil, c.ctx.Err()
	case err := <-errorChan:
		return nil, err
	case response := <-responseChan:
		if response == nil {
			return nil, fmt.Errorf("received nil response (%s)", URL)
		}

		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			if err, ok := StatusToError[response.StatusCode]; ok {
				return nil, fmt.Errorf("status code %d: %s (%s)", response.StatusCode, err, URL)
			}

			return nil, fmt.Errorf("status code %d: unknown error (%s)", response.StatusCode, URL)
		}

		decoder := json.NewDecoder(response.Body)
		if err := decoder.Decode(dest); err != nil {
			return nil, fmt.Errorf("failed to decode response: %w (%s)", err, URL)
		}

		return response, nil
	}
}
