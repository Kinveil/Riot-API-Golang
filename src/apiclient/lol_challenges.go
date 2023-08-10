package apiclient

import (
	"fmt"

	"github.com/junioryono/Riot-API-Golang/src/apiclient/ratelimiter"
	"github.com/junioryono/Riot-API-Golang/src/constants/region"
)

type ChallengesConfig struct {
	ChallengeId string `json:"challengeId"`
	Levels      []struct {
		LevelId string `json:"levelId"`
		Order   int    `json:"order"`
	} `json:"levels"`
	QueueId string `json:"queueId"`
}

type ChallengesLeaderboards struct {
	Entries []struct {
		Entries []struct {
			Percentile float64 `json:"percentile"`
			Rank       int     `json:"rank"`
			Total      int     `json:"total"`
		} `json:"entries"`
		LevelId string `json:"levelId"`
	} `json:"entries"`
	QueueId string `json:"queueId"`
}

type ChallengesPercentiles struct {
	Percentiles []struct {
		Percentile float64 `json:"percentile"`
		Total      int     `json:"total"`
	} `json:"percentiles"`
	QueueId string `json:"queueId"`
}

type ChallengesPlayerData struct {
	Entries []struct {
		Entries []struct {
			Percentile float64 `json:"percentile"`
			Rank       int     `json:"rank"`
			Total      int     `json:"total"`
		} `json:"entries"`
		LevelId string `json:"levelId"`
	} `json:"entries"`
	QueueId string `json:"queueId"`
}

func (c *client) GetChallengesConfig(r region.Region) (*ChallengesConfig, error) {
	var res ChallengesConfig
	_, err := c.dispatchAndUnmarshal(r, "/lol/challenges/v1/config", "", nil, ratelimiter.GetChallengesConfig, &res)
	return &res, err
}

func (c *client) GetChallengesPercentiles(r region.Region) (*ChallengesPercentiles, error) {
	var res ChallengesPercentiles
	_, err := c.dispatchAndUnmarshal(r, "/lol/challenges/v1/percentiles", "", nil, ratelimiter.GetChallengesPercentiles, &res)
	return &res, err
}

func (c *client) GetChallengesConfigById(r region.Region, challengeId string) (*ChallengesConfig, error) {
	var res ChallengesConfig
	_, err := c.dispatchAndUnmarshal(r, "/lol/challenges/v1/config", fmt.Sprintf("/%s", challengeId), nil, ratelimiter.GetChallengesConfigById, &res)
	return &res, err
}

func (c *client) GetChallengesLeaderboardsByLevel(r region.Region, challengeId, level string) (*ChallengesLeaderboards, error) {
	var res ChallengesLeaderboards
	_, err := c.dispatchAndUnmarshal(r, "/lol/challenges/v1/leaderboards", fmt.Sprintf("/%s/%s", challengeId, level), nil, ratelimiter.GetChallengesLeaderboardsByLevel, &res)
	return &res, err
}

func (c *client) GetChallengesPercentilesById(r region.Region, challengeId string) (*ChallengesPercentiles, error) {
	var res ChallengesPercentiles
	_, err := c.dispatchAndUnmarshal(r, "/lol/challenges/v1/percentiles", fmt.Sprintf("/%s", challengeId), nil, ratelimiter.GetChallengesPercentilesById, &res)
	return &res, err
}

func (c *client) GetChallengesPlayerDataByPUUID(r region.Region, puuid string) (*ChallengesPlayerData, error) {
	var res ChallengesPlayerData
	_, err := c.dispatchAndUnmarshal(r, "/lol/challenges/v1/player-data", fmt.Sprintf("/%s", puuid), nil, ratelimiter.GetChallengesPlayerDataByPUUID, &res)
	return &res, err
}
