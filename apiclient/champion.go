package apiclient

import (
	"github.com/junioryono/Riot-API-Golang/apiclient/ratelimiter"
	"github.com/junioryono/Riot-API-Golang/constants/region"
)

type ChampionRotations struct {
	FreeChampionIDs              []int `json:"freeChampionIds"`
	FreeChampionIDsForNewPlayers []int `json:"freeChampionIdsForNewPlayers"`
	MaxNewPlayerLevel            int   `json:"maxNewPlayerLevel"`
}

func (c *client) GetChampionRotations(r region.Region) (*ChampionRotations, error) {
	var championRotation ChampionRotations
	_, err := c.dispatchAndUnmarshal(c.ctx, r, "/lol/platform/v3/champion-rotations", "", nil, ratelimiter.GetChampionRotations, &championRotation)
	return &championRotation, err
}
