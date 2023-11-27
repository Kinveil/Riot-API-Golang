package apiclient_test

import (
	"testing"

	"github.com/junioryono/Riot-API-Golang/src/constants/region"
)

func TestGetChallengesConfig(t *testing.T) {
	client := newTestClient(t, nil)

	config, err := client.GetChallengesConfig(region.NA1)
	if err != nil {
		t.Fatalf("Failed to get challenges config: %v", err)
	}

	if config == nil {
		t.Fatalf("Expected to receive challenges config but got nil")
	}
}

func TestGetChallengesPercentiles(t *testing.T) {
	client := newTestClient(t, nil)

	percentiles, err := client.GetChallengesPercentiles(region.NA1)
	if err != nil {
		t.Fatalf("Failed to get challenges percentiles: %v", err)
	}

	if percentiles == nil {
		t.Fatalf("Expected to receive challenges percentiles but got nil")
	}
}

func TestGetChallengesPlayerDataByPuuid(t *testing.T) {
	client := newTestClient(t, nil)

	summoner, err := client.GetSummonerByName(region.NA1, "Mighty Junior")
	if err != nil {
		t.Fatalf("Failed to get summoner: %v", err)
	}

	playerData, err := client.GetChallengesPlayerDataByPuuid(region.NA1, summoner.Puuid)
	if err != nil {
		t.Fatalf("Failed to get challenges player data by Puuid: %v", err)
	}

	if playerData == nil {
		t.Fatalf("Expected to receive challenges player data but got nil")
	}
}
