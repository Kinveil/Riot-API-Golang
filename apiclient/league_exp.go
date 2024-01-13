package apiclient

import (
	"fmt"
	"strings"

	"github.com/Kinveil-Engineering-Analysis/Riot-API-Golang/apiclient/ratelimiter"
	"github.com/Kinveil-Engineering-Analysis/Riot-API-Golang/constants/queue_ranked"
	"github.com/Kinveil-Engineering-Analysis/Riot-API-Golang/constants/region"
)

func (c *client) GetLeagueExpEntries(r region.Region, q queue_ranked.String, tier, division string, page int) ([]LeagueEntry, error) {
	var res []LeagueEntry
	_, err := c.dispatchAndUnmarshal(c.ctx, r, "/lol/league-exp/v4/entries", fmt.Sprintf("/%s/%s/%s?page=%d", string(q), strings.ToUpper(tier), strings.ToUpper(division), page), nil, ratelimiter.GetLeagueExpEntries, &res)
	return res, err
}
