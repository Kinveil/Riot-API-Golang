package ratelimiter

type MethodId string

func (m MethodId) String() string {
	return string(m)
}

const (
	// ----- Champion Mastery API -----
	GetChampionMasteriesBySummonerId            MethodId = "GetChampionMasteriesBySummonerId"
	GetChampionMasteryBySummonerIdAndChampionId MethodId = "GetChampionMasteryBySummonerIdAndChampionId"
	GetChampionMasteriesTopBySummonerId         MethodId = "GetChampionMasteriesTopBySummonerId"
	GetChampionMasteryScoreTotalBySummonerId    MethodId = "GetChampionMasteryScoreTotalBySummonerId"

	// ----- Champion API -----
	GetChampionRotations MethodId = "GetChampionRotations"

	// ----- Clash API -----
	GetClashPlayersByPUUID      MethodId = "GetClashPlayersByPUUID"
	GetClashPlayersBySummonerId MethodId = "GetClashPlayersBySummonerId"
	GetClashTeamById            MethodId = "GetClashTeamById"
	GetClashTournaments         MethodId = "GetClashTournaments"
	GetClashTournamentByTeamId  MethodId = "GetClashTournamentByTeamId"
	GetClashTournamentById      MethodId = "GetClashTournamentById"

	// ----- League Exp API -----
	GetLeagueExpEntries MethodId = "GetLeagueExpEntries"

	// ----- League API -----
	GetLeagueEntriesChallenger   MethodId = "GetLeagueEntriesChallenger"
	GetLeagueEntriesGrandmaster  MethodId = "GetLeagueEntriesGrandmaster"
	GetLeagueEntriesMaster       MethodId = "GetLeagueEntriesMaster"
	GetLeagueEntries             MethodId = "GetLeagueEntries"
	GetLeagueEntriesById         MethodId = "GetLeagueEntriesById"
	GetLeagueEntriesBySummonerId MethodId = "GetLeagueEntriesBySummonerId"

	// ----- LOL Challenges API -----
	GetChallengesConfig              MethodId = "GetChallengesConfig"
	GetChallengesPercentiles         MethodId = "GetChallengesPercentiles"
	GetChallengesConfigById          MethodId = "GetChallengesConfigById"
	GetChallengesLeaderboardsByLevel MethodId = "GetChallengesLeaderboardsByLevel"
	GetChallengesPercentilesById     MethodId = "GetChallengesPercentilesById"
	GetChallengesPlayerDataByPUUID   MethodId = "GetChallengesPlayerDataByPUUID"

	// ----- LOL Status API -----
	GetStatusPlatformData MethodId = "GetStatusPlatformData"

	// ----- Match API -----
	GetMatchlist     MethodId = "GetMatchlist"
	GetMatch         MethodId = "GetMatch"
	GetMatchTimeline MethodId = "GetMatchTimeline"

	// ----- Spectator API -----
	GetSpectatorActiveGameBySummonerId MethodId = "GetSpectatorActiveGameBySummonerId"
	GetSpectatorFeaturedGames          MethodId = "GetSpectatorFeaturedGames"

	// ----- Summoner API -----
	GetSummonerByRsoPUUID      MethodId = "GetSummonerByRsoPUUID"
	GetSummonerByAccountId     MethodId = "GetSummonerByAccountId"
	GetSummonerBySummonerName  MethodId = "GetSummonerBySummonerName"
	GetSummonerBySummonerPUUID MethodId = "GetSummonerBySummonerPUUID"
	GetSummonerBySummonerId    MethodId = "GetSummonerBySummonerId"
)
