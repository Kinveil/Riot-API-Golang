package main

import (
	"fmt"

	"github.com/junioryono/Riot-API-Golang/src/apiclient"
	"github.com/junioryono/Riot-API-Golang/src/apiclient/ratelimiter"
	"github.com/junioryono/Riot-API-Golang/src/constants/region"
)

func main() {
	apiKey := "RGAPI-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
	client := apiclient.New(apiKey)

	conservation := ratelimiter.ConserveUsage{
		RegionPercent: 30,
		MethodPercent: 30,
		IgnoreLimits: []ratelimiter.MethodId{
			ratelimiter.GetLeagueEntriesChallenger,
			ratelimiter.GetLeagueEntriesGrandmaster,
			ratelimiter.GetLeagueEntriesMaster,
			ratelimiter.GetLeagueEntries,
		},
	}

	client.SetUsageConservation(conservation)

	summoner, err := client.GetSummonerBySummonerName(region.NA1, "Mighty Junior")
	if err != nil {
		panic(err)
	}

	matchlist, err := client.GetMatchlist(region.NA1.Continent(), summoner.Puuid, nil)
	if err != nil {
		panic(err)
	}

	matchId := (*matchlist)[0]

	match, err := client.GetMatch(region.NA1.Continent(), matchId)
	if err != nil {
		panic(err)
	}

	fmt.Println(match.Info.GameCreation)
}
