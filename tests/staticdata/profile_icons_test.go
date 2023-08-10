package staticdata_test

import (
	"testing"

	"github.com/junioryono/Riot-API-Golang/src/constants/language"
	"github.com/junioryono/Riot-API-Golang/src/staticdata"
)

func TestGetProfileIcons(t *testing.T) {
	currentPatch := getCurrentPatch(t)

	profileIcons, err := staticdata.GetProfileIcons(currentPatch, language.EnglishUnitedStates)
	if err != nil {
		t.Fatalf("Failed to get items: %v", err)
	}

	if len(profileIcons) == 0 {
		t.Fatalf("Expected profile icons to be greater than 0, got %v", len(profileIcons))
	}
}

func TestGetProfileIcon(t *testing.T) {
	currentPatch := getCurrentPatch(t)

	profileIcons, err := staticdata.GetProfileIcons(currentPatch, language.EnglishUnitedStates)
	if err != nil {
		t.Fatalf("Failed to get items: %v", err)
	}

	if len(profileIcons) == 0 {
		t.Fatalf("Expected profile icons to be greater than 0, got %v", len(profileIcons))
	}

	profileIcon, err := profileIcons.ProfileIcon(1)
	if err != nil {
		t.Fatalf("Failed to get item: %v", err)
	}

	if profileIcon.Id != 1 {
		t.Fatalf("Expected profile icon id to be 1, got %v", profileIcon.Id)
	}
}
