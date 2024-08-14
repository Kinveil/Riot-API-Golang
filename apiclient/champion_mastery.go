package apiclient

import (
	"fmt"

	"github.com/Kinveil/Riot-API-Golang/apiclient/ratelimiter"
	"github.com/Kinveil/Riot-API-Golang/constants/region"
)

type ChampionMastery struct {
	Puuid                        string `json:"puuid"`
	ChampionID                   int    `json:"championId"`
	ChampionLevel                int    `json:"championLevel"`
	ChampionPoints               int    `json:"championPoints"`
	LastPlayTime                 int    `json:"lastPlayTime"`
	ChampionPointsSinceLastLevel int    `json:"championPointsSinceLastLevel"`
	ChampionPointsUntilNextLevel int    `json:"championPointsUntilNextLevel"`
	ChestGranted                 bool   `json:"chestGranted"`
	TokensEarned                 int    `json:"tokensEarned"`
	SummonerID                   string `json:"summonerId"`
}

func (c *uniqueClient) GetChampionMasteriesBySummonerID(r region.Region, summonerID string) ([]ChampionMastery, error) {
	var res []ChampionMastery
	_, err := c.dispatchAndUnmarshal(r, "/lol/champion-mastery/v4/champion-masteries/by-summoner", fmt.Sprintf("/%s", summonerID), nil, ratelimiter.GetChampionMasteriesBySummonerID, &res)
	return res, err
}

func (c *uniqueClient) GetChampionMasteryBySummonerIDAndChampionID(r region.Region, summonerID string, championID int) (*ChampionMastery, error) {
	var res ChampionMastery
	_, err := c.dispatchAndUnmarshal(r, "/lol/champion-mastery/v4/champion-masteries/by-summoner", fmt.Sprintf("/%s/by-champion/%d", summonerID, championID), nil, ratelimiter.GetChampionMasteryBySummonerIDAndChampionID, &res)
	return &res, err
}

func (c *uniqueClient) GetChampionMasteriesTopBySummonerID(r region.Region, summonerID string) ([]ChampionMastery, error) {
	var res []ChampionMastery
	_, err := c.dispatchAndUnmarshal(r, "/lol/champion-mastery/v4/champion-masteries/by-summoner", fmt.Sprintf("/%s/top", summonerID), nil, ratelimiter.GetChampionMasteriesTopBySummonerID, &res)
	return res, err
}

func (c *uniqueClient) GetChampionMasteryScoreTotalBySummonerID(r region.Region, summonerID string) (int, error) {
	var res int
	_, err := c.dispatchAndUnmarshal(r, "/lol/champion-mastery/v4/scores/by-summoner", fmt.Sprintf("/%s", summonerID), nil, ratelimiter.GetChampionMasteryScoreTotalBySummonerID, &res)
	return res, err
}
