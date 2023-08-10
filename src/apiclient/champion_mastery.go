package apiclient

import (
	"fmt"

	"github.com/junioryono/Riot-API-Golang/src/apiclient/ratelimiter"
	"github.com/junioryono/Riot-API-Golang/src/constants/region"
)

type ChampionMastery struct {
	Puuid                        string `json:"puuid"`
	ChampionId                   int    `json:"championId"`
	ChampionLevel                int    `json:"championLevel"`
	ChampionPoints               int    `json:"championPoints"`
	LastPlayTime                 int    `json:"lastPlayTime"`
	ChampionPointsSinceLastLevel int    `json:"championPointsSinceLastLevel"`
	ChampionPointsUntilNextLevel int    `json:"championPointsUntilNextLevel"`
	ChestGranted                 bool   `json:"chestGranted"`
	TokensEarned                 int    `json:"tokensEarned"`
	SummonerId                   string `json:"summonerId"`
}

func (c *client) GetChampionMasteriesBySummonerId(r region.Region, summonerId string) ([]ChampionMastery, error) {
	var res []ChampionMastery
	_, err := c.dispatchAndUnmarshal(r, "/lol/champion-mastery/v4/champion-masteries/by-summoner", fmt.Sprintf("/%s", summonerId), nil, ratelimiter.GetChampionMasteriesBySummonerId, &res)
	return res, err
}

func (c *client) GetChampionMasteryBySummonerIdAndChampionId(r region.Region, summonerId string, championId int) (*ChampionMastery, error) {
	var res ChampionMastery
	_, err := c.dispatchAndUnmarshal(r, "/lol/champion-mastery/v4/champion-masteries/by-summoner", fmt.Sprintf("/%s/by-champion/%d", summonerId, championId), nil, ratelimiter.GetChampionMasteryBySummonerIdAndChampionId, &res)
	return &res, err
}

func (c *client) GetChampionMasteriesTopBySummonerId(r region.Region, summonerId string) ([]ChampionMastery, error) {
	var res []ChampionMastery
	_, err := c.dispatchAndUnmarshal(r, "/lol/champion-mastery/v4/champion-masteries/by-summoner", fmt.Sprintf("/%s/top", summonerId), nil, ratelimiter.GetChampionMasteriesTopBySummonerId, &res)
	return res, err
}

func (c *client) GetChampionMasteryScoreTotalBySummonerId(r region.Region, summonerId string) (int, error) {
	var res int
	_, err := c.dispatchAndUnmarshal(r, "/lol/champion-mastery/v4/scores/by-summoner", fmt.Sprintf("/%s", summonerId), nil, ratelimiter.GetChampionMasteryScoreTotalBySummonerId, &res)
	return res, err
}
