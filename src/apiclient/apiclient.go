package apiclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/junioryono/Riot-API-Golang/src/apiclient/ratelimiter"
	"github.com/junioryono/Riot-API-Golang/src/constants/continent"
	"github.com/junioryono/Riot-API-Golang/src/constants/queue"
	"github.com/junioryono/Riot-API-Golang/src/constants/region"
)

type Client interface {
	// Helper methods to set the API key, usage conservation, and max retries.

	SetUsageConservation(conserveUsage ratelimiter.ConserveUsage)
	SetAPIKey(apiKey string)
	SetMaxRetries(maxRetries int)

	// Account API

	GetAccountByPUUID(continent continent.Continent, puuid string) (*Account, error)
	GetAccountByRiotId(continent continent.Continent, gameName, tagLine string) (*Account, error)

	// Champion Mastery API

	GetChampionMasteriesBySummonerId(region region.Region, summonerId string) ([]ChampionMastery, error)
	GetChampionMasteryBySummonerIdAndChampionId(region region.Region, summonerId string, championId int) (*ChampionMastery, error)
	GetChampionMasteriesTopBySummonerId(region region.Region, summonerId string) ([]ChampionMastery, error)
	GetChampionMasteryScoreTotalBySummonerId(region region.Region, summonerId string) (int, error)

	// Champion API

	GetChampionRotations(region region.Region) (*ChampionRotations, error)

	// Clash API

	GetClashPlayersByPUUID(region region.Region, puuid string) (*ClashPlayers, error)
	GetClashPlayersBySummonerId(region region.Region, summonerId string) (*ClashPlayers, error)
	GetClashTeamById(region region.Region, teamId string) (*ClashTeam, error)
	GetClashTournaments(region region.Region) (*ClashTournaments, error)
	GetClashTournamentByTeamId(region region.Region, teamId string) (*ClashTournament, error)
	GetClashTournamentById(region region.Region, tournamentId string) (*ClashTournament, error)

	// League Exp API

	GetLeagueExpEntries(region region.Region, q queue.Queue, tier, division string, page int) ([]LeaguePosition, error)

	// League API

	GetLeagueEntriesChallenger(region region.Region, q queue.Queue) (*LeagueList, error)
	GetLeagueEntriesGrandmaster(region region.Region, q queue.Queue) (*LeagueList, error)
	GetLeagueEntriesMaster(region region.Region, q queue.Queue) (*LeagueList, error)
	GetLeagueEntries(region region.Region, q queue.Queue, tier, division string, page int) ([]LeaguePosition, error)
	GetLeagueEntriesById(region region.Region, leagueId string) (*LeagueList, error)
	GetLeagueEntriesBySummonerId(region region.Region, summonerId string) ([]LeaguePosition, error)

	// LOL Challenges API

	GetChallengesConfig(region region.Region) (*ChallengesConfig, error)
	GetChallengesPercentiles(region region.Region) (*ChallengesPercentiles, error)
	GetChallengesConfigById(region region.Region, challengeId string) (*ChallengesConfig, error)
	GetChallengesLeaderboardsByLevel(region region.Region, challengeId, level string) (*ChallengesLeaderboards, error)
	GetChallengesPercentilesById(region region.Region, challengeId string) (*ChallengesPercentiles, error)
	GetChallengesPlayerDataByPUUID(region region.Region, puuid string) (*ChallengesPlayerData, error)

	// LOL Status API

	GetStatusPlatformData(region region.Region) (*StatusPlatformData, error)

	// Match API

	GetMatchlist(continent continent.Continent, puuid string, opts *GetMatchlistOptions) (*Matchlist, error)
	GetMatch(continent continent.Continent, matchId string) (*Match, error)
	GetMatchTimeline(continent continent.Continent, matchId string) (*MatchTimeline, error)

	// Spectator API

	GetSpectatorActiveGameBySummonerId(region region.Region, summonerId string) (*ActiveGame, error)
	GetSpectatorFeaturedGames(region region.Region) (*FeaturedGames, error)

	// Summoner API

	GetSummonerByRsoPUUID(region region.Region, rsoPuuid string) (*Summoner, error)
	GetSummonerByAccountId(region region.Region, accountId string) (*Summoner, error)
	GetSummonerBySummonerName(region region.Region, name string) (*Summoner, error)
	GetSummonerBySummonerPUUID(region region.Region, puuid string) (*Summoner, error)
	GetSummonerBySummonerId(region region.Region, summonerId string) (*Summoner, error)
}

// client is the internal implementation of Client.
type client struct {
	ratelimiter *ratelimiter.RateLimiter
}

// New returns a Client configured for the given API client and underlying HTTP
// client. The returned Client is threadsafe.
func New(apiKey string) Client {
	requests := make(chan *ratelimiter.APIRequest)

	ratelimiter := ratelimiter.NewRateLimiter(requests, apiKey)
	go ratelimiter.Start()

	return &client{
		ratelimiter: ratelimiter,
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

func (c *client) dispatchAndUnmarshal(regionOrContinent HostProvider, method string, relativePath string, parameters url.Values, methodId ratelimiter.MethodId, dest interface{}) (*http.Response, error) {
	var suffix, separator string

	if len(parameters) > 0 {
		suffix = fmt.Sprintf("?%s", parameters.Encode())
	}

	if !strings.HasPrefix(relativePath, "/") {
		separator = "/"
	}

	URL := regionOrContinent.Host() + method + separator + relativePath + suffix

	responseChan := make(chan *http.Response)
	newRequest := ratelimiter.APIRequest{
		Region:   strings.ToUpper(regionOrContinent.String()),
		MethodId: methodId,
		URL:      URL,
		Response: responseChan,
	}

	c.ratelimiter.Requests <- &newRequest
	response := <-responseChan

	if response == nil {
		return nil, fmt.Errorf("response is nil")
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return response, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, dest)
	return response, err
}
