package staticdata_test

import (
	"testing"

	"github.com/junioryono/Riot-API-Golang/src/constants/language"
	"github.com/junioryono/Riot-API-Golang/src/staticdata"
)

func TestGetChampions(t *testing.T) {
	currentPatch := getCurrentPatch(t)

	champions, err := staticdata.GetChampions(currentPatch, language.EnglishUnitedStates)
	if err != nil {
		t.Fatalf("Failed to get champions: %v", err)
	}

	if len(champions) == 0 {
		t.Fatalf("Expected champions to be greater than 0, got %v", len(champions))
	}
}

func TestGetChampion(t *testing.T) {
	currentPatch := getCurrentPatch(t)

	champion, err := staticdata.GetChampion(currentPatch, language.EnglishUnitedStates, "MonkeyKing")
	if err != nil {
		t.Fatalf("Failed to get champion: %v", err)
	}

	if champion.Name != "Wukong" {
		t.Fatalf("Expected champion name to be Wukong, got %v", champion.Name)
	}
}
