package staticdata_test

import (
	"testing"

	"github.com/junioryono/Riot-API-Golang/src/constants/language"
	"github.com/junioryono/Riot-API-Golang/src/staticdata"
)

func TestGetSummonerSpells(t *testing.T) {
	currentPatch := getCurrentPatch(t)

	summonerSpells, err := staticdata.GetSummonerSpells(currentPatch, language.EnglishUnitedStates)
	if err != nil {
		t.Fatalf("Failed to get summonerSpells: %v", err)
	}

	if len(summonerSpells) == 0 {
		t.Fatalf("Expected summoner spells to be greater than 0, got %v", len(summonerSpells))
	}
}

func TestGetSummonerSpell(t *testing.T) {
	currentPatch := getCurrentPatch(t)

	summonerSpells, err := staticdata.GetSummonerSpells(currentPatch, language.EnglishUnitedStates)
	if err != nil {
		t.Fatalf("Failed to get summonerSpells: %v", err)
	}

	if len(summonerSpells) == 0 {
		t.Fatalf("Expected summoner spells to be greater than 0, got %v", len(summonerSpells))
	}

	summonerSpell, err := summonerSpells.SummonerSpell("SummonerBarrier")
	if err != nil {
		t.Fatalf("Failed to get summoner spell: %v", err)
	}

	if summonerSpell.ID != "SummonerBarrier" {
		t.Fatalf("Expected SummonerBarrier, got %v", summonerSpell.ID)
	}
}
