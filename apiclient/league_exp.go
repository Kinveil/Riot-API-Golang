package apiclient

import (
	"fmt"
	"strings"

	"github.com/junioryono/Riot-API-Golang/apiclient/ratelimiter"
	"github.com/junioryono/Riot-API-Golang/constants/queue"
	"github.com/junioryono/Riot-API-Golang/constants/region"
)

func (c *client) GetLeagueExpEntries(r region.Region, q queue.Queue, tier, division string, page int) ([]LeaguePosition, error) {
	var res []LeaguePosition
	_, err := c.dispatchAndUnmarshal(c.ctx, r, "/lol/league-exp/v4/entries", fmt.Sprintf("/%s/%s/%s?page=%d", q.String(), strings.ToUpper(tier), strings.ToUpper(division), page), nil, ratelimiter.GetLeagueExpEntries, &res)
	return res, err
}
