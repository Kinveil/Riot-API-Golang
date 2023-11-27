package apiclient

import (
	"encoding/json"
	"fmt"

	"github.com/junioryono/Riot-API-Golang/src/apiclient/ratelimiter"
	"github.com/junioryono/Riot-API-Golang/src/constants/region"
)

type Summoner struct {
	AccountID     string `json:"accountId"`
	ID            string `json:"id"`
	Name          string `json:"name"`
	ProfileIconID int    `json:"profileIconId"`
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

func (c *client) GetSummonerByRsoPuuid(r region.Region, rsoPuuid string) (*Summoner, error) {
	var res Summoner
	_, err := c.dispatchAndUnmarshal(r, "/lol/summoner/v4/summoners/by-account", fmt.Sprintf("/%s", rsoPuuid), nil, ratelimiter.GetSummonerByRsoPuuid, &res)
	return &res, err
}

func (c *client) GetSummonerByAccountID(r region.Region, accountID string) (*Summoner, error) {
	var res Summoner
	_, err := c.dispatchAndUnmarshal(r, "/lol/summoner/v4/summoners/by-account", fmt.Sprintf("/%s", accountID), nil, ratelimiter.GetSummonerByAccountID, &res)
	return &res, err
}

func (c *client) GetSummonerBySummonerName(r region.Region, name string) (*Summoner, error) {
	var res Summoner
	_, err := c.dispatchAndUnmarshal(r, "/lol/summoner/v4/summoners/by-name", fmt.Sprintf("/%s", name), nil, ratelimiter.GetSummonerBySummonerName, &res)
	return &res, err
}

func (c *client) GetSummonerBySummonerPuuid(r region.Region, puuid string) (*Summoner, error) {
	var res Summoner
	_, err := c.dispatchAndUnmarshal(r, "/lol/summoner/v4/summoners/by-puuid", fmt.Sprintf("/%s", puuid), nil, ratelimiter.GetSummonerBySummonerPuuid, &res)
	return &res, err
}

func (c *client) GetSummonerBySummonerID(r region.Region, summonerID string) (*Summoner, error) {
	var res Summoner
	_, err := c.dispatchAndUnmarshal(r, "/lol/summoner/v4/summoners", fmt.Sprintf("/%s", summonerID), nil, ratelimiter.GetSummonerBySummonerID, &res)
	return &res, err
}
