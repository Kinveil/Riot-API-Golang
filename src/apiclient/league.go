package apiclient

import (
	"fmt"
	"strings"

	"github.com/junioryono/Riot-API-Golang/src/apiclient/ratelimiter"
	"github.com/junioryono/Riot-API-Golang/src/constants/queue"
	"github.com/junioryono/Riot-API-Golang/src/constants/region"
)

type LeagueList struct {
	LeagueId string           `json:"leagueId"`
	Tier     string           `json:"tier"`
	Entries  []LeaguePosition `json:"entries"`
	Queue    queue.Queue      `json:"queue"`
	Name     string           `json:"name"`
}

type LeaguePosition struct {
	FreshBlood   bool        `json:"freshBlood"`
	HotStreak    bool        `json:"hotStreak"`
	Inactive     bool        `json:"inactive"`
	LeagueId     string      `json:"leagueId"`
	LeaguePoints int         `json:"leaguePoints"`
	Losses       int         `json:"losses"`
	QueueType    string      `json:"queueType"`
	Rank         string      `json:"rank"`
	SummonerId   string      `json:"summonerId"`
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

func (c *client) GetLeagueEntriesChallenger(r region.Region, q queue.Queue) (*LeagueList, error) {
	var res LeagueList
	_, err := c.dispatchAndUnmarshal(r, "/lol/league/v4/challengerleagues/by-queue", fmt.Sprintf("/%s", q.String()), nil, ratelimiter.GetLeagueEntriesChallenger, &res)
	return &res, err
}

func (c *client) GetLeagueEntriesGrandmaster(r region.Region, q queue.Queue) (*LeagueList, error) {
	var res LeagueList
	_, err := c.dispatchAndUnmarshal(r, "/lol/league/v4/grandmasterleagues/by-queue", fmt.Sprintf("/%s", q.String()), nil, ratelimiter.GetLeagueEntriesGrandmaster, &res)
	return &res, err
}

func (c *client) GetLeagueEntriesMaster(r region.Region, q queue.Queue) (*LeagueList, error) {
	var res LeagueList
	_, err := c.dispatchAndUnmarshal(r, "/lol/league/v4/masterleagues/by-queue", fmt.Sprintf("/%s", q.String()), nil, ratelimiter.GetLeagueEntriesMaster, &res)
	return &res, err
}

func (c *client) GetLeagueEntries(r region.Region, q queue.Queue, tier, division string, page int) ([]LeaguePosition, error) {
	var res []LeaguePosition
	_, err := c.dispatchAndUnmarshal(r, "/lol/league/v4/entries", fmt.Sprintf("/%s/%s/%s?page=%d", q.String(), strings.ToUpper(tier), strings.ToUpper(division), page), nil, ratelimiter.GetLeagueEntries, &res)
	return res, err
}

func (c *client) GetLeagueEntriesById(r region.Region, leagueId string) (*LeagueList, error) {
	var res LeagueList
	_, err := c.dispatchAndUnmarshal(r, "/lol/league/v4/leagues", fmt.Sprintf("/%s", leagueId), nil, ratelimiter.GetLeagueEntriesById, &res)
	return &res, err
}

func (c *client) GetLeagueEntriesBySummonerId(r region.Region, summonerId string) ([]LeaguePosition, error) {
	var res []LeaguePosition
	_, err := c.dispatchAndUnmarshal(r, "/lol/league/v4/entries/by-summoner", fmt.Sprintf("/%s", summonerId), nil, ratelimiter.GetLeagueEntriesBySummonerId, &res)
	return res, err
}
