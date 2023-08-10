package apiclient_test

import (
	"testing"

	"github.com/junioryono/Riot-API-Golang/src/constants/queue"
	"github.com/junioryono/Riot-API-Golang/src/constants/region"
)

func TestGetLeagueEntriesChallenger(t *testing.T) {
	client := newTestClient(t, nil)
	q := queue.RankedSolo5x5

	leagueList, err := client.GetLeagueEntriesChallenger(region.NA1, q)
	if err != nil {
		t.Fatalf("Failed to get league entries for Challenger: %v", err)
	}

	if leagueList == nil {
		t.Fatalf("Expected to receive league entries but got nil")
	}
}

func TestGetLeagueEntriesGrandmaster(t *testing.T) {
	client := newTestClient(t, nil)
	q := queue.RankedSolo5x5

	leagueList, err := client.GetLeagueEntriesGrandmaster(region.NA1, q)
	if err != nil {
		t.Fatalf("Failed to get league entries for Grandmaster: %v", err)
	}

	if leagueList == nil {
		t.Fatalf("Expected to receive league entries but got nil")
	}
}

func TestGetLeagueEntriesMaster(t *testing.T) {
	client := newTestClient(t, nil)
	q := queue.RankedSolo5x5

	leagueList, err := client.GetLeagueEntriesMaster(region.NA1, q)
	if err != nil {
		t.Fatalf("Failed to get league entries for Master: %v", err)
	}

	if leagueList == nil {
		t.Fatalf("Expected to receive league entries but got nil")
	}
}

func TestGetLeagueEntries(t *testing.T) {
	client := newTestClient(t, nil)
	q := queue.RankedSolo5x5
	tier := "Gold"
	division := "II"
	page := 1

	leaguePositions, err := client.GetLeagueEntries(region.NA1, q, tier, division, page)
	if err != nil {
		t.Fatalf("Failed to get league entries: %v", err)
	}

	if len(leaguePositions) == 0 {
		t.Fatalf("Expected to receive league positions but got an empty slice")
	}
}

func TestGetLeagueEntriesBySummonerId(t *testing.T) {
	client := newTestClient(t, nil)

	summoner, err := client.GetSummonerBySummonerName(region.NA1, "Mighty Junior")
	if err != nil {
		t.Fatalf("Failed to get summoner: %v", err)
	}

	leaguePositions, err := client.GetLeagueEntriesBySummonerId(region.NA1, summoner.Id)
	if err != nil {
		t.Fatalf("Failed to get league entries by summoner Id: %v", err)
	}

	if len(leaguePositions) == 0 {
		t.Fatalf("Expected to receive league positions but got an empty slice")
	}
}
