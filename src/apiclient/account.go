package apiclient

import (
	"fmt"

	"github.com/junioryono/Riot-API-Golang/src/apiclient/ratelimiter"
	"github.com/junioryono/Riot-API-Golang/src/constants/continent"
)

type Account struct {
	PUUID    string `json:"puuid"`
	GameName string `json:"gameName"`
	TagLine  string `json:"tagLine"`
}

func (c *client) GetAccountByPUUID(continent continent.Continent, puuid string) (*Account, error) {
	var account Account
	_, err := c.dispatchAndUnmarshal(continent, "/riot/account/v1/accounts/by-puuid", fmt.Sprintf("/%s", puuid), nil, ratelimiter.GetAccountByPUUID, &account)
	return &account, err
}

func (c *client) GetAccountByRiotId(continent continent.Continent, gameName, tagLine string) (*Account, error) {
	var account Account
	_, err := c.dispatchAndUnmarshal(continent, "/riot/account/v1/accounts/by-riot-id", fmt.Sprintf("/%s/%s", gameName, tagLine), nil, ratelimiter.GetAccountByRiotId, &account)
	return &account, err
}
