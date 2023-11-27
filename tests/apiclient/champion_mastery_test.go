package apiclient_test

import (
	"testing"

	"github.com/junioryono/Riot-API-Golang/src/constants/region"
)

func TestGetChampionMasteriesBySummonerID(t *testing.T) {
	client := newTestClient(t, nil)

	summoner, err := client.GetSummonerByName(region.NA1, "Mighty Junior")
	if err != nil {
		t.Fatalf("Failed to get summoner: %v", err)
	}

	championMasteries, err := client.GetChampionMasteriesBySummonerID(region.NA1, summoner.ID)
	if err != nil {
		t.Fatalf("Failed to get champion masteries: %v", err)
	}

	if len(championMasteries) == 0 {
		t.Fatalf("Expected to receive champion masteries but got an empty slice")
	}
}

func TestGetChampionMasteryBySummonerIDAndChampionID(t *testing.T) {
	client := newTestClient(t, nil)

	summoner, err := client.GetSummonerByName(region.NA1, "Mighty Junior")
	if err != nil {
		t.Fatalf("Failed to get summoner: %v", err)
	}

	championMastery, err := client.GetChampionMasteryBySummonerIDAndChampionID(region.NA1, summoner.ID, 1)
	if err != nil {
		t.Fatalf("Failed to get champion mastery: %v", err)
	}

	if championMastery.ChampionID != 1 {
		t.Fatalf("Expected champion ID to be 1 but got %d", championMastery.ChampionID)
	}
}

func TestGetChampionMasteriesTopBySummonerID(t *testing.T) {
	client := newTestClient(t, nil)

	summoner, err := client.GetSummonerByName(region.NA1, "Mighty Junior")
	if err != nil {
		t.Fatalf("Failed to get summoner: %v", err)
	}

	championMasteries, err := client.GetChampionMasteriesTopBySummonerID(region.NA1, summoner.ID)
	if err != nil {
		t.Fatalf("Failed to get champion masteries: %v", err)
	}

	if len(championMasteries) == 0 {
		t.Fatalf("Expected to receive champion masteries but got an empty slice")
	}
}

func TestGetChampionMasteryScoreTotalBySummonerID(t *testing.T) {
	client := newTestClient(t, nil)

	summoner, err := client.GetSummonerByName(region.NA1, "Mighty Junior")
	if err != nil {
		t.Fatalf("Failed to get summoner: %v", err)
	}

	score, err := client.GetChampionMasteryScoreTotalBySummonerID(region.NA1, summoner.ID)
	if err != nil {
		t.Fatalf("Failed to get champion mastery score: %v", err)
	}

	if score == 0 {
		t.Fatalf("Expected to receive a champion mastery score but got 0")
	}
}
