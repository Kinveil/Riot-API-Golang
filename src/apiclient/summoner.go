package apiclient

import (
	"encoding/json"
	"fmt"

	"github.com/junioryono/Riot-API-Golang/src/apiclient/ratelimiter"
	"github.com/junioryono/Riot-API-Golang/src/constants/region"
)

type Summoner struct {
	AccountId     string `json:"accountId"`
	Id            string `json:"id"`
	Name          string `json:"name"`
	ProfileIconId int    `json:"profileIconId"`
	Puuid         string `json:"puuid"`
	RevisionDate  int    `json:"revisionDate"`
	SummonerLevel int    `json:"summonerLevel"`
}

func (s Summoner) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

func (s *Summoner) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, s)
}

func (c *client) GetSummonerByRsoPUUID(r region.Region, rsoPuuid string) (*Summoner, error) {
	var res Summoner
	_, err := c.dispatchAndUnmarshal(r, "/lol/summoner/v4/summoners/by-account", fmt.Sprintf("/%s", rsoPuuid), nil, ratelimiter.GetSummonerByRsoPUUID, &res)
	return &res, err
}

func (c *client) GetSummonerByAccountId(r region.Region, accountId string) (*Summoner, error) {
	var res Summoner
	_, err := c.dispatchAndUnmarshal(r, "/lol/summoner/v4/summoners/by-account", fmt.Sprintf("/%s", accountId), nil, ratelimiter.GetSummonerByAccountId, &res)
	return &res, err
}

func (c *client) GetSummonerBySummonerName(r region.Region, name string) (*Summoner, error) {
	var res Summoner
	_, err := c.dispatchAndUnmarshal(r, "/lol/summoner/v4/summoners/by-name", fmt.Sprintf("/%s", name), nil, ratelimiter.GetSummonerBySummonerName, &res)
	return &res, err
}

func (c *client) GetSummonerBySummonerPUUID(r region.Region, puuid string) (*Summoner, error) {
	var res Summoner
	_, err := c.dispatchAndUnmarshal(r, "/lol/summoner/v4/summoners/by-puuid", fmt.Sprintf("/%s", puuid), nil, ratelimiter.GetSummonerBySummonerPUUID, &res)
	return &res, err
}

func (c *client) GetSummonerBySummonerId(r region.Region, summonerId string) (*Summoner, error) {
	var res Summoner
	_, err := c.dispatchAndUnmarshal(r, "/lol/summoner/v4/summoners", fmt.Sprintf("/%s", summonerId), nil, ratelimiter.GetSummonerBySummonerId, &res)
	return &res, err
}
