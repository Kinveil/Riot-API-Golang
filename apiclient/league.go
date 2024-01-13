package apiclient

import (
	"fmt"
	"strings"

	"github.com/Kinveil-Engineering-Analysis/Riot-API-Golang/apiclient/ratelimiter"
	"github.com/Kinveil-Engineering-Analysis/Riot-API-Golang/constants/league/rank"
	"github.com/Kinveil-Engineering-Analysis/Riot-API-Golang/constants/league/tier"
	"github.com/Kinveil-Engineering-Analysis/Riot-API-Golang/constants/queue_ranked"
	"github.com/Kinveil-Engineering-Analysis/Riot-API-Golang/constants/region"
)

type LeagueList struct {
	LeagueID string              `json:"leagueId"`
	Tier     tier.String         `json:"tier"`
	Entries  []LeagueItem        `json:"entries"`
	Queue    queue_ranked.String `json:"queue"`
	Name     string              `json:"name"`
}

type LeagueItem struct {
	SummonerID   string      `json:"summonerId"`
	SummonerName string      `json:"summonerName"`
	LeaguePoints int         `json:"leaguePoints"`
	Rank         rank.String `json:"rank"`
	Wins         int         `json:"wins"`
	Losses       int         `json:"losses"`
	Veteran      bool        `json:"veteran"`
	Inactive     bool        `json:"inactive"`
	FreshBlood   bool        `json:"freshBlood"`
	HotStreak    bool        `json:"hotStreak"`
}

type LeagueEntry struct {
	FreshBlood   bool                `json:"freshBlood"`
	HotStreak    bool                `json:"hotStreak"`
	Inactive     bool                `json:"inactive"`
	LeagueID     string              `json:"leagueId"`
	LeaguePoints int                 `json:"leaguePoints"`
	Losses       int                 `json:"losses"`
	QueueType    queue_ranked.String `json:"queueType"`
	Rank         rank.String         `json:"rank"`
	SummonerID   string              `json:"summonerId"`
	SummonerName string              `json:"summonerName"`
	Tier         tier.String         `json:"tier"`
	Veteran      bool                `json:"veteran"`
	Wins         int                 `json:"wins"`
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

func (c *client) GetLeagueEntries(r region.Region, q queue_ranked.String, tier, division string, page int) ([]LeagueEntry, error) {
	var res []LeagueEntry
	_, err := c.dispatchAndUnmarshal(c.ctx, r, "/lol/league/v4/entries", fmt.Sprintf("/%s/%s/%s?page=%d", string(q), strings.ToUpper(tier), strings.ToUpper(division), page), nil, ratelimiter.GetLeagueEntries, &res)
	return res, err
}

func (c *client) GetLeagueEntriesByID(r region.Region, leagueID string) (*LeagueList, error) {
	var res LeagueList
	_, err := c.dispatchAndUnmarshal(c.ctx, r, "/lol/league/v4/leagues", fmt.Sprintf("/%s", leagueID), nil, ratelimiter.GetLeagueEntriesByID, &res)
	return &res, err
}

func (c *client) GetLeagueEntriesBySummonerID(r region.Region, summonerID string) ([]LeagueEntry, error) {
	var res []LeagueEntry
	_, err := c.dispatchAndUnmarshal(c.ctx, r, "/lol/league/v4/entries/by-summoner", fmt.Sprintf("/%s", summonerID), nil, ratelimiter.GetLeagueEntriesBySummonerID, &res)
	return res, err
}
