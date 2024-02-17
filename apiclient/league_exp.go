package apiclient

import (
	"fmt"

	"github.com/Kinveil-Engineering-Analysis/Riot-API-Golang/apiclient/ratelimiter"
	"github.com/Kinveil-Engineering-Analysis/Riot-API-Golang/constants/league/rank"
	"github.com/Kinveil-Engineering-Analysis/Riot-API-Golang/constants/league/tier"
	"github.com/Kinveil-Engineering-Analysis/Riot-API-Golang/constants/queue_ranked"
	"github.com/Kinveil-Engineering-Analysis/Riot-API-Golang/constants/region"
)

func (c *client) GetLeagueExpEntries(r region.Region, q queue_ranked.String, tier tier.String, rank rank.String, page int) ([]LeagueEntry, error) {
	var res []LeagueEntry
	_, err := c.dispatchAndUnmarshal(r, "/lol/league-exp/v4/entries", fmt.Sprintf("/%s/%s/%s?page=%d", q, tier, rank, page), nil, ratelimiter.GetLeagueExpEntries, &res)
	return res, err
}
