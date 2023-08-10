package apiclient_test

import (
	"testing"

	"github.com/junioryono/Riot-API-Golang/src/constants/queue"
	"github.com/junioryono/Riot-API-Golang/src/constants/region"
)

func TestGetLeagueExpEntries(t *testing.T) {
	client := newTestClient(t, nil)

	q := queue.RankedSolo5x5
	tier := "GOLD"
	division := "II"
	page := 1

	leaguePositions, err := client.GetLeagueExpEntries(region.NA1, q, tier, division, page)
	if err != nil {
		t.Fatalf("Failed to get league exp entries: %v", err)
	}

	if len(leaguePositions) == 0 {
		t.Fatalf("Expected to receive league positions but got an empty slice")
	}
}
