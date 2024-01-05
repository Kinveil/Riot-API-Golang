package apiclient

import (
	"fmt"
	"strings"

	"github.com/junioryono/Riot-API-Golang/apiclient/ratelimiter"
	"github.com/junioryono/Riot-API-Golang/constants/queue_ranked"
	"github.com/junioryono/Riot-API-Golang/constants/region"
)

type LeagueList struct {
	LeagueID string              `json:"leagueId"`
	Tier     string              `json:"tier"`
	Entries  []LeaguePosition    `json:"entries"`
	Queue    queue_ranked.String `json:"queue"`
	Name     string              `json:"name"`
}

type LeaguePosition struct {
	FreshBlood   bool        `json:"freshBlood"`
	HotStreak    bool        `json:"hotStreak"`
	Inactive     bool        `json:"inactive"`
	LeagueID     string      `json:"leagueId"`
	LeaguePoints int         `json:"leaguePoints"`
	Losses       int         `json:"losses"`
	QueueType    string      `json:"queueType"`
	Rank         string      `json:"rank"`
	SummonerID   string      `json:"summonerId"`
	SummonerName string      `json:"summonerName"`
	Tier         string      `json:"tier"`
	Veteran      bool        `json:"veteran"`
	Wins         int         `json:"wins"`
	MiniSeries   *MiniSeries `json:"miniSeries"`
}

type MiniSeries struct {
	Wins     int    `json:"wins"`
	Losses   int    `json:"losses"`
	Target   int    `json:"target"`
	Progress string `json:"progress"`
}

func (c *client) GetLeagueEntriesChallenger(r region.Region, q queue_ranked.String) (*LeagueList, error) {
	var res LeagueList
	_, err := c.dispatchAndUnmarshal(c.ctx, r, "/lol/league/v4/challengerleagues/by-queue", fmt.Sprintf("/%s", string(q)), nil, ratelimiter.GetLeagueEntriesChallenger, &res)
	return &res, err
}

func (c *client) GetLeagueEntriesGrandmaster(r region.Region, q queue_ranked.String) (*LeagueList, error) {
	var res LeagueList
	_, err := c.dispatchAndUnmarshal(c.ctx, r, "/lol/league/v4/grandmasterleagues/by-queue", fmt.Sprintf("/%s", string(q)), nil, ratelimiter.GetLeagueEntriesGrandmaster, &res)
	return &res, err
}

func (c *client) GetLeagueEntriesMaster(r region.Region, q queue_ranked.String) (*LeagueList, error) {
	var res LeagueList
	_, err := c.dispatchAndUnmarshal(c.ctx, r, "/lol/league/v4/masterleagues/by-queue", fmt.Sprintf("/%s", string(q)), nil, ratelimiter.GetLeagueEntriesMaster, &res)
	return &res, err
}

func (c *client) GetLeagueEntries(r region.Region, q queue_ranked.String, tier, division string, page int) ([]LeaguePosition, error) {
	var res []LeaguePosition
	_, err := c.dispatchAndUnmarshal(c.ctx, r, "/lol/league/v4/entries", fmt.Sprintf("/%s/%s/%s?page=%d", string(q), strings.ToUpper(tier), strings.ToUpper(division), page), nil, ratelimiter.GetLeagueEntries, &res)
	return res, err
}

func (c *client) GetLeagueEntriesByID(r region.Region, leagueID string) (*LeagueList, error) {
	var res LeagueList
	_, err := c.dispatchAndUnmarshal(c.ctx, r, "/lol/league/v4/leagues", fmt.Sprintf("/%s", leagueID), nil, ratelimiter.GetLeagueEntriesByID, &res)
	return &res, err
}

func (c *client) GetLeagueEntriesBySummonerID(r region.Region, summonerID string) ([]LeaguePosition, error) {
	var res []LeaguePosition
	_, err := c.dispatchAndUnmarshal(c.ctx, r, "/lol/league/v4/entries/by-summoner", fmt.Sprintf("/%s", summonerID), nil, ratelimiter.GetLeagueEntriesBySummonerID, &res)
	return res, err
}
