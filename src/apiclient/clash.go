package apiclient

import (
	"fmt"

	"github.com/junioryono/Riot-API-Golang/src/apiclient/ratelimiter"
	"github.com/junioryono/Riot-API-Golang/src/constants/region"
)

type ClashTeam struct {
	ID           string        `json:"id"`
	TournamentID int           `json:"tournamentId"`
	Name         string        `json:"name"`
	IconID       int           `json:"iconId"`
	Tier         int           `json:"tier"`
	Captain      string        `json:"captain"`
	Abbreviation string        `json:"abbreviation"`
	Players      []ClashPlayer `json:"players"`
}

type ClashPlayers struct {
	Players []ClashPlayer `json:"players"`
}

type ClashPlayer struct {
	TeamID       string `json:"teamId"`
	Position     string `json:"position"`
	SummonerID   string `json:"summonerId"`
	SummonerName string `json:"summonerName"`
	GameName     string `json:"gameName"`
	TeamRole     string `json:"teamRole"`
	TeamName     string `json:"teamName"`
}

type ClashTournaments struct {
	Tournaments []ClashTournament `json:"tournaments"`
}

type ClashTournament struct {
	ID               int    `json:"id"`
	ThemeId          int    `json:"themeId"`
	NameKey          string `json:"nameKey"`
	NameKeySecondary string `json:"nameKeySecondary"`
	Schedule         []struct {
		ID               int    `json:"id"`
		RegistrationTime int64  `json:"registrationTime"`
		StartTime        int64  `json:"startTime"`
		WaitTime         int    `json:"waitTime"`
		NotifyAt         int    `json:"notifyAt"`
		State            string `json:"state"`
	} `json:"schedule"`
	CompletedPhase int `json:"completedPhase"`
	Prizes         []struct {
		TeamID string `json:"teamId"`
		Prize  int    `json:"prize"`
	} `json:"prizes"`
	Stages []struct {
		ID                  int   `json:"id"`
		Stage               int   `json:"stage"`
		StartTime           int64 `json:"startTime"`
		WaitTime            int   `json:"waitTime"`
		NotifyAt            int   `json:"notifyAt"`
		RegistrationTime    int64 `json:"registrationTime"`
		Length              int   `json:"length"`
		TeamSize            int   `json:"teamSize"`
		MaxTeams            int   `json:"maxTeams"`
		IsMajor             bool  `json:"isMajor"`
		RegistrationEnabled bool  `json:"registrationEnabled"`
	} `json:"stages"`
	MaxNumTeams int   `json:"maxNumTeams"`
	SignupTime  int64 `json:"signupTime"`
	StartTime   int64 `json:"startTime"`
	Winners     []struct {
		TeamID string `json:"teamId"`
		Place  int    `json:"place"`
	} `json:"winners"`
	TeamSize int `json:"teamSize"`
	Entrants []struct {
		ID                    string        `json:"id"`
		TeamID                string        `json:"teamId"`
		TeamName              string        `json:"teamName"`
		TeamIconId            int           `json:"teamIconId"`
		TeamIcon              string        `json:"teamIcon"`
		TeamTier              int           `json:"teamTier"`
		Players               []ClashPlayer `json:"players"`
		ProvisioningFlowId    string        `json:"provisioningFlowId"`
		ProvisioningFlowState string        `json:"provisioningFlowState"`
	} `json:"entrants"`
}

func (c *client) GetClashPlayersByPuuid(r region.Region, puuid string) (*ClashPlayers, error) {
	var res ClashPlayers
	_, err := c.dispatchAndUnmarshal(r, "/lol/clash/v1/players/by-puuid", fmt.Sprintf("/%s", puuid), nil, ratelimiter.GetClashPlayersByPuuid, &res)
	return &res, err
}

func (c *client) GetClashPlayersBySummonerID(r region.Region, summonerID string) (*ClashPlayers, error) {
	var res ClashPlayers
	_, err := c.dispatchAndUnmarshal(r, "/lol/clash/v1/players/by-summoner", fmt.Sprintf("/%s", summonerID), nil, ratelimiter.GetClashPlayersBySummonerID, &res)
	return &res, err
}

func (c *client) GetClashTeamByID(r region.Region, teamID string) (*ClashTeam, error) {
	var res ClashTeam
	_, err := c.dispatchAndUnmarshal(r, "/lol/clash/v1/teams", fmt.Sprintf("/%s", teamID), nil, ratelimiter.GetClashTeamByID, &res)
	return &res, err
}

func (c *client) GetClashTournaments(r region.Region) (*ClashTournaments, error) {
	var res ClashTournaments
	_, err := c.dispatchAndUnmarshal(r, "/lol/clash/v1/tournaments", "", nil, ratelimiter.GetClashTournaments, &res)
	return &res, err
}

func (c *client) GetClashTournamentByTeamID(r region.Region, teamID string) (*ClashTournament, error) {
	var res ClashTournament
	_, err := c.dispatchAndUnmarshal(r, "/lol/clash/v1/tournaments/by-team", fmt.Sprintf("/%s", teamID), nil, ratelimiter.GetClashTournamentByTeamID, &res)
	return &res, err
}

func (c *client) GetClashTournamentByID(r region.Region, tournamentID string) (*ClashTournament, error) {
	var res ClashTournament
	_, err := c.dispatchAndUnmarshal(r, "/lol/clash/v1/tournaments", fmt.Sprintf("/%s", tournamentID), nil, ratelimiter.GetClashTournamentByID, &res)
	return &res, err
}
