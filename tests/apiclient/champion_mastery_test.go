package apiclient_test

import (
	"testing"

	"github.com/junioryono/Riot-API-Golang/src/constants/region"
)

func TestGetChampionMasteriesBySummonerId(t *testing.T) {
	client := newTestClient(t, nil)

	summoner, err := client.GetSummonerBySummonerName(region.NA1, "Mighty Junior")
	if err != nil {
		t.Fatalf("Failed to get summoner: %v", err)
	}

	championMasteries, err := client.GetChampionMasteriesBySummonerId(region.NA1, summoner.Id)
	if err != nil {
		t.Fatalf("Failed to get champion masteries: %v", err)
	}

	if len(championMasteries) == 0 {
		t.Fatalf("Expected to receive champion masteries but got an empty slice")
	}
}

func TestGetChampionMasteryBySummonerIdAndChampionId(t *testing.T) {
	client := newTestClient(t, nil)

	summoner, err := client.GetSummonerBySummonerName(region.NA1, "Mighty Junior")
	if err != nil {
		t.Fatalf("Failed to get summoner: %v", err)
	}

	championMastery, err := client.GetChampionMasteryBySummonerIdAndChampionId(region.NA1, summoner.Id, 1)
	if err != nil {
		t.Fatalf("Failed to get champion mastery: %v", err)
	}

	if championMastery.ChampionId != 1 {
		t.Fatalf("Expected champion Id to be 1 but got %d", championMastery.ChampionId)
	}
}

func TestGetChampionMasteriesTopBySummonerId(t *testing.T) {
	client := newTestClient(t, nil)

	summoner, err := client.GetSummonerBySummonerName(region.NA1, "Mighty Junior")
	if err != nil {
		t.Fatalf("Failed to get summoner: %v", err)
	}

	championMasteries, err := client.GetChampionMasteriesTopBySummonerId(region.NA1, summoner.Id)
	if err != nil {
		t.Fatalf("Failed to get champion masteries: %v", err)
	}

	if len(championMasteries) == 0 {
		t.Fatalf("Expected to receive champion masteries but got an empty slice")
	}
}

func TestGetChampionMasteryScoreTotalBySummonerId(t *testing.T) {
	client := newTestClient(t, nil)

	summoner, err := client.GetSummonerBySummonerName(region.NA1, "Mighty Junior")
	if err != nil {
		t.Fatalf("Failed to get summoner: %v", err)
	}

	score, err := client.GetChampionMasteryScoreTotalBySummonerId(region.NA1, summoner.Id)
	if err != nil {
		t.Fatalf("Failed to get champion mastery score: %v", err)
	}

	if score == 0 {
		t.Fatalf("Expected to receive a champion mastery score but got 0")
	}
}
