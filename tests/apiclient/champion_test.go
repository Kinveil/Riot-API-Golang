package apiclient_test

import (
	"testing"

	"github.com/junioryono/Riot-API-Golang/src/constants/region"
)

func TestGetChampionRotations(t *testing.T) {
	client := newTestClient(t, nil)

	championRotations, err := client.GetChampionRotations(region.NA1)
	if err != nil {
		t.Fatalf("Failed to get champion rotations: %v", err)
	}

	if len(championRotations.FreeChampionIds) == 0 {
		t.Fatalf("Expected to receive free champion Ids but got an empty slice")
	}

	if len(championRotations.FreeChampionIdsForNewPlayers) == 0 {
		t.Fatalf("Expected to receive free champion Ids for new players but got an empty slice")
	}
}
