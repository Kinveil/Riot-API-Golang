package apiclient

import (
	"github.com/Kinveil/Riot-API-Golang/apiclient/ratelimiter"
	"github.com/Kinveil/Riot-API-Golang/constants/region"
)

type ChampionRotations struct {
	FreeChampionIDs              []int `json:"freeChampionIds"`
	FreeChampionIDsForNewPlayers []int `json:"freeChampionIdsForNewPlayers"`
	MaxNewPlayerLevel            int   `json:"maxNewPlayerLevel"`
}

func (c *uniqueClient) GetChampionRotations(r region.Region) (*ChampionRotations, error) {
	var championRotation ChampionRotations
	err := c.dispatchAndUnmarshal(r, "/lol/platform/v3/champion-rotations", "", nil, ratelimiter.GetChampionRotations, &championRotation)
	return &championRotation, err
}
