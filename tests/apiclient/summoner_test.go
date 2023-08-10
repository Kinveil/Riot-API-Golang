package apiclient_test

import (
	"testing"

	"github.com/junioryono/Riot-API-Golang/src/constants/region"
)

func TestGetSummonerBySummonerName(t *testing.T) {
	client := newTestClient(t, nil)

	summoner, err := client.GetSummonerBySummonerName(region.NA1, "Mighty Junior")
	if err != nil {
		t.Fatalf("Failed to get summoner: %v", err)
	}

	if summoner.Name != "Mighty Junior" {
		t.Errorf("Summoner name is not Mighty Junior: %s", summoner.Name)
	}
}
