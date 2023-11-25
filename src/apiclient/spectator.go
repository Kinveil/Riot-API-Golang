package apiclient

import (
	"fmt"

	"github.com/junioryono/Riot-API-Golang/src/apiclient/ratelimiter"
	"github.com/junioryono/Riot-API-Golang/src/constants/region"
)

type ActiveGame struct {
	GameId            int64                   `json:"gameId"`            // The Id of the game
	MapId             int64                   `json:"mapId"`             // The Id of the map
	GameMode          string                  `json:"gameMode"`          // The game mode
	GameType          string                  `json:"gameType"`          // The game type
	GameQueueConfigId int64                   `json:"gameQueueConfigId"` // The queue type (queue types are documented on the Game Constants page)
	Participants      []ActiveGameParticipant `json:"participants"`      // The participant information
	Observers         Observer                `json:"observers"`         // The observer information
	PlatformId        string                  `json:"platformId"`        // The Id of the platform on which the game is being played
	BannedChampions   []BannedChampion        `json:"bannedChampions"`   // Banned champion information
	GameStartTime     int64                   `json:"gameStartTime"`     // The game start time represented in epoch milliseconds
	GameLength        int64                   `json:"gameLength"`        // The amount of time in seconds that has passed since the game started
}

type ActiveGameParticipant struct {
	PUUID         string                             `json:"puuid"`         // The PUUID of the player
	ProfileIconId int64                              `json:"profileIconId"` // The Id of the profile icon used by this participant
	ChampionId    int64                              `json:"championId"`    // The Id of the champion played by this participant
	SummonerName  string                             `json:"summonerName"`  // The summoner name of this participant
	Runes         []CurrentGameParticipantRuneDTO    `json:"runes"`         // The runes used by this participant
	Bot           bool                               `json:"bot"`           // Flag indicating whether or not this participant is a bot
	TeamId        int64                              `json:"teamId"`        // The team Id of this participant, indicating the participant's team
	Spell2Id      int64                              `json:"spell2Id"`      // The Id of the second summoner spell used by this participant
	Masteries     []CurrentGameParticipantMasteryDTO `json:"masteries"`     // The masteries used by this participant
	Spell1Id      int64                              `json:"spell1Id"`      // The Id of the first summoner spell used by this participant
	SummonerId    string                             `json:"summonerId"`    // The encrypted summoner Id of this participant
	Perks         Perks                              `json:"perks"`
}

type Perks struct {
	PerkIds      []int64 `json:"perkIds"`
	PerkStyle    int64   `json:"perkStyle"`
	PerkSubStyle int64   `json:"perkSubStyle"`
}

type CurrentGameParticipantRuneDTO struct {
	Count  int   `json:"count"`  // The count of this rune used by the participant
	RuneId int64 `json:"runeId"` // The Id of the rune
}

type CurrentGameParticipantMasteryDTO struct {
	MasteryId int64 `json:"masteryId"` // The Id of the mastery
	Rank      int   `json:"rank"`      // The number of points put into this mastery by the user
}

type Observer struct {
	EncryptionKey string `json:"encryptionKey"` // Key used to decrypt the spectator grid game data for playback
}

type BannedChampion struct {
	PickTurn   int   `json:"pickTurn"`   // The turn during which the champion was banned
	ChampionId int64 `json:"championId"` // The Id of the banned champion
	TeamId     int64 `json:"teamId"`     // The Id of the team that banned the champion
}

func (c *client) GetSpectatorActiveGameBySummonerId(r region.Region, summonerId string) (*ActiveGame, error) {
	var res ActiveGame
	_, err := c.dispatchAndUnmarshal(r, "/lol/spectator/v4/active-games/by-summoner", fmt.Sprintf("/%s", summonerId), nil, ratelimiter.GetSpectatorActiveGameBySummonerId, &res)
	return &res, err
}

type FeaturedGames struct {
	ClientRefreshInterval int64                 `json:"clientRefreshInterval"` // The suggested interval to wait before requesting FeaturedGames again
	GameList              []FeaturedGameInfoDTO `json:"gameList"`              // 	The list of featured games
}

type FeaturedGameInfoDTO struct {
	GameId            int64                        `json:"gameId"`            // The Id of the game
	GameStartTime     int64                        `json:"gameStartTime"`     // The game start time represented in epoch milliseconds
	PlatformId        string                       `json:"platformId"`        // The Id of the platform on which the game is being played
	GameMode          string                       `json:"gameMode"`          // The game mode
	MapId             int64                        `json:"mapId"`             // The Id of the map
	GameType          string                       `json:"gameType"`          // The game type
	BannedChampions   []BannedChampion             `json:"bannedChampions"`   // Banned champion information
	Observers         Observer                     `json:"observers"`         // The observer information
	Participants      []FeaturedGameParticipantDTO `json:"participants"`      // The participant information
	GameLength        int64                        `json:"gameLength"`        // The amount of time in seconds that has passed since the game started
	GameQueueConfigId int64                        `json:"gameQueueConfigId"` // The queue type (queue types are documented on the Game Constants page)
}

type FeaturedGameParticipantDTO struct {
	ProfileIconId int64  `json:"profileIconId"` // The Id of the profile icon used by this participant
	ChampionId    int64  `json:"championId"`    // The Id of the champion played by this participant
	SummonerName  string `json:"summonerName"`  // The summoner name of this participant
	Bot           bool   `json:"bot"`           // Flag indicating whether or not this participant is a bot
	Spell2Id      int64  `json:"spell2Id"`      // The Id of the second summoner spell used by this participant
	TeamId        int64  `json:"teamId"`        // The team Id of this participant, indicating the participant's team
	Spell1Id      int64  `json:"spell1Id"`      // The Id of the first summoner spell used by this participant
	Perks         Perks  `json:"perks"`
}

func (c *client) GetSpectatorFeaturedGames(r region.Region) (*FeaturedGames, error) {
	var res FeaturedGames
	_, err := c.dispatchAndUnmarshal(r, "/lol/spectator/v4/featured-games", "", nil, ratelimiter.GetSpectatorFeaturedGames, &res)
	return &res, err
}
