package apiclient_test

import (
	"testing"

	"github.com/junioryono/Riot-API-Golang/src/apiclient"
	"github.com/junioryono/Riot-API-Golang/src/constants/region"
)

func TestGetMatchlist(t *testing.T) {
	client := newTestClient(t, nil)

	summonerRegion := region.NA1
	summonerContinent := summonerRegion.Continent()
	summonerName := "Mighty Junior"

	summoner, err := client.GetSummonerBySummonerName(summonerRegion, summonerName)
	if err != nil {
		t.Fatalf("Failed to get summoner: %v", err)
	}

	matchlist, err := client.GetMatchlist(summonerContinent, summoner.Puuid, nil)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if matchlist == nil || len(*matchlist) == 0 {
		t.Fatalf("Expected non-empty matchlist but got %v", matchlist)
	}

	count := 1
	opts := &apiclient.GetMatchlistOptions{
		Count: &count,
	}

	matchlist, err = client.GetMatchlist(summonerContinent, summoner.Puuid, opts)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if matchlist == nil || len(*matchlist) != 1 {
		t.Fatalf("Expected matchlist with 1 match but got %v", matchlist)
	}
}

func TestGetMatch(t *testing.T) {
	client := newTestClient(t, nil)

	summonerRegion := region.NA1
	summonerContinent := summonerRegion.Continent()
	summonerName := "Mighty Junior"

	summoner, err := client.GetSummonerBySummonerName(summonerRegion, summonerName)
	if err != nil {
		t.Fatalf("Failed to get summoner: %v", err)
	}

	count := 1
	opts := &apiclient.GetMatchlistOptions{
		Count: &count,
	}

	matchlist, err := client.GetMatchlist(summonerContinent, summoner.Puuid, opts)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if matchlist == nil || len(*matchlist) == 0 {
		t.Fatalf("Expected non-empty matchlist but got %v", matchlist)
	}

	matchID := (*matchlist)[0]

	match, err := client.GetMatch(summonerContinent, matchID)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if match == nil || match.Metadata.MatchID != matchID {
		t.Fatalf("Expected valid match but got %v", match)
	}
}

func TestGetMatchTimeline(t *testing.T) {
	client := newTestClient(t, nil)

	summonerRegion := region.NA1
	summonerContinent := summonerRegion.Continent()
	summonerName := "Mighty Junior"

	summoner, err := client.GetSummonerBySummonerName(summonerRegion, summonerName)
	if err != nil {
		t.Fatalf("Failed to get summoner: %v", err)
	}

	count := 1
	opts := &apiclient.GetMatchlistOptions{
		Count: &count,
	}

	matchlist, err := client.GetMatchlist(summonerContinent, summoner.Puuid, opts)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if matchlist == nil || len(*matchlist) == 0 {
		t.Fatalf("Expected non-empty matchlist but got %v", matchlist)
	}

	matchID := (*matchlist)[0]

	match, err := client.GetMatchTimeline(summonerContinent, matchID)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if match == nil || match.MatchTimelineMetadata.MatchID != matchID {
		t.Fatalf("Expected valid match but got %v", match)
	}
}
