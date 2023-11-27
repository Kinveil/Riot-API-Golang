package apiclient

import (
	"fmt"

	"github.com/junioryono/Riot-API-Golang/src/apiclient/ratelimiter"
	"github.com/junioryono/Riot-API-Golang/src/constants/region"
)

type ActiveGame struct {
	GameID            int64                   `json:"gameId"`            // The ID of the game
	MapID             int64                   `json:"mapId"`             // The ID of the map
	GameMode          string                  `json:"gameMode"`          // The game mode
	GameType          string                  `json:"gameType"`          // The game type
	GameQueueConfigID int64                   `json:"gameQueueConfigId"` // The queue type (queue types are documented on the Game Constants page)
	Participants      []ActiveGameParticipant `json:"participants"`      // The participant information
	Observers         Observer                `json:"observers"`         // The observer information
	PlatformID        string                  `json:"platformId"`        // The ID of the platform on which the game is being played
	BannedChampions   []BannedChampion        `json:"bannedChampions"`   // Banned champion information
	GameStartTime     int64                   `json:"gameStartTime"`     // The game start time represented in epoch milliseconds
	GameLength        int64                   `json:"gameLength"`        // The amount of time in seconds that has passed since the game started
}

type ActiveGameParticipant struct {
	PUUID         string                             `json:"puuid"`         // The PUUID of the player
	ProfileIconID int64                              `json:"profileIconId"` // The ID of the profile icon used by this participant
	ChampionID    int64                              `json:"championId"`    // The ID of the champion played by this participant
	SummonerName  string                             `json:"summonerName"`  // The summoner name of this participant
	Runes         []CurrentGameParticipantRuneDTO    `json:"runes"`         // The runes used by this participant
	Bot           bool                               `json:"bot"`           // Flag indicating whether or not this participant is a bot
	TeamID        int64                              `json:"teamId"`        // The team ID of this participant, indicating the participant's team
	Spell1ID      int64                              `json:"spell1Id"`      // The ID of the first summoner spell used by this participant
	Spell2ID      int64                              `json:"spell2Id"`      // The ID of the second summoner spell used by this participant
	Masteries     []CurrentGameParticipantMasteryDTO `json:"masteries"`     // The masteries used by this participant
	SummonerID    string                             `json:"summonerId"`    // The encrypted summoner ID of this participant
	Perks         Perks                              `json:"perks"`
}

type Perks struct {
	PerkIDs      []int64 `json:"perkIds"`
	PerkStyle    int64   `json:"perkStyle"`
	PerkSubStyle int64   `json:"perkSubStyle"`
}

type CurrentGameParticipantRuneDTO struct {
	Count  int   `json:"count"`  // The count of this rune used by the participant
	RuneID int64 `json:"runeId"` // The ID of the rune
}

type CurrentGameParticipantMasteryDTO struct {
	MasteryID int64 `json:"masteryId"` // The ID of the mastery
	Rank      int   `json:"rank"`      // The number of points put into this mastery by the user
}

type Observer struct {
	EncryptionKey string `json:"encryptionKey"` // Key used to decrypt the spectator grid game data for playback
}

type BannedChampion struct {
	PickTurn   int   `json:"pickTurn"`   // The turn during which the champion was banned
	ChampionID int64 `json:"championId"` // The ID of the banned champion
	TeamID     int64 `json:"teamId"`     // The ID of the team that banned the champion
}

func (c *client) GetSpectatorActiveGameBySummonerID(r region.Region, summonerID string) (*ActiveGame, error) {
	var res ActiveGame
	_, err := c.dispatchAndUnmarshal(r, "/lol/spectator/v4/active-games/by-summoner", fmt.Sprintf("/%s", summonerID), nil, ratelimiter.GetSpectatorActiveGameBySummonerID, &res)
	return &res, err
}

type FeaturedGames struct {
	ClientRefreshInterval int64                 `json:"clientRefreshInterval"` // The suggested interval to wait before requesting FeaturedGames again
	GameList              []FeaturedGameInfoDTO `json:"gameList"`              // 	The list of featured games
}

type FeaturedGameInfoDTO struct {
	GameID            int64                        `json:"gameId"`            // The ID of the game
	GameStartTime     int64                        `json:"gameStartTime"`     // The game start time represented in epoch milliseconds
	PlatformID        string                       `json:"platformId"`        // The ID of the platform on which the game is being played
	GameMode          string                       `json:"gameMode"`          // The game mode
	MapID             int64                        `json:"mapId"`             // The ID of the map
	GameType          string                       `json:"gameType"`          // The game type
	BannedChampions   []BannedChampion             `json:"bannedChampions"`   // Banned champion information
	Observers         Observer                     `json:"observers"`         // The observer information
	Participants      []FeaturedGameParticipantDTO `json:"participants"`      // The participant information
	GameLength        int64                        `json:"gameLength"`        // The amount of time in seconds that has passed since the game started
	GameQueueConfigID int64                        `json:"gameQueueConfigId"` // The queue type (queue types are documented on the Game Constants page)
}

type FeaturedGameParticipantDTO struct {
	ProfileIconID int64  `json:"profileIconId"` // The ID of the profile icon used by this participant
	ChampionID    int64  `json:"championId"`    // The ID of the champion played by this participant
	SummonerName  string `json:"summonerName"`  // The summoner name of this participant
	Bot           bool   `json:"bot"`           // Flag indicating whether or not this participant is a bot
	Spell1ID      int64  `json:"spell1Id"`      // The ID of the first summoner spell used by this participant
	Spell2ID      int64  `json:"spell2Id"`      // The ID of the second summoner spell used by this participant
	TeamID        int64  `json:"teamId"`        // The team ID of this participant, indicating the participant's team
	Perks         Perks  `json:"perks"`
}

func (c *client) GetSpectatorFeaturedGames(r region.Region) (*FeaturedGames, error) {
	var res FeaturedGames
	_, err := c.dispatchAndUnmarshal(r, "/lol/spectator/v4/featured-games", "", nil, ratelimiter.GetSpectatorFeaturedGames, &res)
	return &res, err
}
