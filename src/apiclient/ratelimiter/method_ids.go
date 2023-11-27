package ratelimiter

type MethodID string

func (m MethodID) String() string {
	return string(m)
}

const (
	// ----- Account API -----
	GetAccountByPuuid  MethodID = "GetAccountByPuuid"
	GetAccountByRiotID MethodID = "GetAccountByRiotID"

	// ----- Champion Mastery API -----
	GetChampionMasteriesBySummonerID            MethodID = "GetChampionMasteriesBySummonerID"
	GetChampionMasteryBySummonerIDAndChampionID MethodID = "GetChampionMasteryBySummonerIDAndChampionID"
	GetChampionMasteriesTopBySummonerID         MethodID = "GetChampionMasteriesTopBySummonerID"
	GetChampionMasteryScoreTotalBySummonerID    MethodID = "GetChampionMasteryScoreTotalBySummonerID"

	// ----- Champion API -----
	GetChampionRotations MethodID = "GetChampionRotations"

	// ----- Clash API -----
	GetClashPlayersByPuuid      MethodID = "GetClashPlayersByPuuid"
	GetClashPlayersBySummonerID MethodID = "GetClashPlayersBySummonerID"
	GetClashTeamByID            MethodID = "GetClashTeamByID"
	GetClashTournaments         MethodID = "GetClashTournaments"
	GetClashTournamentByTeamID  MethodID = "GetClashTournamentByTeamID"
	GetClashTournamentByID      MethodID = "GetClashTournamentByID"

	// ----- League Exp API -----
	GetLeagueExpEntries MethodID = "GetLeagueExpEntries"

	// ----- League API -----
	GetLeagueEntriesChallenger   MethodID = "GetLeagueEntriesChallenger"
	GetLeagueEntriesGrandmaster  MethodID = "GetLeagueEntriesGrandmaster"
	GetLeagueEntriesMaster       MethodID = "GetLeagueEntriesMaster"
	GetLeagueEntries             MethodID = "GetLeagueEntries"
	GetLeagueEntriesByID         MethodID = "GetLeagueEntriesByID"
	GetLeagueEntriesBySummonerID MethodID = "GetLeagueEntriesBySummonerID"

	// ----- LOL Challenges API -----
	GetChallengesConfig              MethodID = "GetChallengesConfig"
	GetChallengesPercentiles         MethodID = "GetChallengesPercentiles"
	GetChallengesConfigByID          MethodID = "GetChallengesConfigByID"
	GetChallengesLeaderboardsByLevel MethodID = "GetChallengesLeaderboardsByLevel"
	GetChallengesPercentilesByID     MethodID = "GetChallengesPercentilesByID"
	GetChallengesPlayerDataByPuuid   MethodID = "GetChallengesPlayerDataByPuuid"

	// ----- LOL Status API -----
	GetStatusPlatformData MethodID = "GetStatusPlatformData"

	// ----- Match API -----
	GetMatchlist     MethodID = "GetMatchlist"
	GetMatch         MethodID = "GetMatch"
	GetMatchTimeline MethodID = "GetMatchTimeline"

	// ----- Spectator API -----
	GetSpectatorActiveGameBySummonerID MethodID = "GetSpectatorActiveGameBySummonerID"
	GetSpectatorFeaturedGames          MethodID = "GetSpectatorFeaturedGames"

	// ----- Summoner API -----
	GetSummonerByRsoPuuid   MethodID = "GetSummonerByRsoPuuid"
	GetSummonerByAccountID  MethodID = "GetSummonerByAccountID"
	GetSummonerByName       MethodID = "GetSummonerByName"
	GetSummonerByPuuid      MethodID = "GetSummonerByPuuid"
	GetSummonerBySummonerID MethodID = "GetSummonerBySummonerID"
)
