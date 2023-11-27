package staticdata

import (
	"fmt"

	"github.com/junioryono/Riot-API-Golang/src/constants/language"
	"github.com/junioryono/Riot-API-Golang/src/constants/patch"
)

type SummonerSpells []SummonerSpell

type SummonerSpell struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Tooltip       string    `json:"tooltip"`
	MaxRank       int       `json:"maxrank"`
	Cooldown      []float64 `json:"cooldown"`
	CooldownBurn  string    `json:"cooldownBurn"`
	Cost          []int     `json:"cost"`
	CostBurn      string    `json:"costBurn"`
	Key           string    `json:"key"`
	SummonerLevel int       `json:"summonerLevel"`
	Modes         []string  `json:"modes"`
	CostType      string    `json:"costType"`
	MaxAmmo       string    `json:"maxammo"`
	Range         []int     `json:"range"`
	RangeBurn     string    `json:"rangeBurn"`
	Image         Image     `json:"image"`
	Resource      string    `json:"resource"`
}

func GetSummonerSpells(v patch.Patch, lang language.Language) (SummonerSpells, error) {
	type Response struct {
		Data map[string]SummonerSpell `json:"data"`
	}

	var res Response
	err := getJSON(fmt.Sprintf("http://ddragon.leagueoflegends.com/cdn/%s/data/%s/summoner.json", v, lang), &res)

	var summonerSpells SummonerSpells
	for _, summonerSpell := range res.Data {
		summonerSpells = append(summonerSpells, summonerSpell)
	}

	return summonerSpells, err
}

func (s SummonerSpells) SummonerSpell(summonerSpellId string) (SummonerSpell, error) {
	for _, summonerSpell := range s {
		if summonerSpell.ID == summonerSpellId {
			return summonerSpell, nil
		}
	}

	return SummonerSpell{}, fmt.Errorf("summoner spell %s not found", summonerSpellId)
}
