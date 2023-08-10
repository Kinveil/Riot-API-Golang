package staticdata_test

import (
	"testing"

	"github.com/junioryono/Riot-API-Golang/src/constants/language"
	"github.com/junioryono/Riot-API-Golang/src/staticdata"
)

func TestGetItems(t *testing.T) {
	currentPatch := getCurrentPatch(t)

	items, err := staticdata.GetItems(currentPatch, language.EnglishUnitedStates)
	if err != nil {
		t.Fatalf("Failed to get items: %v", err)
	}

	if len(items.Data) == 0 {
		t.Fatalf("Expected items to be greater than 0, got %v", len(items.Data))
	}
}

func TestGetItem(t *testing.T) {
	currentPatch := getCurrentPatch(t)

	items, err := staticdata.GetItems(currentPatch, language.EnglishUnitedStates)
	if err != nil {
		t.Fatalf("Failed to get items: %v", err)
	}

	if len(items.Data) == 0 {
		t.Fatalf("Expected items to be greater than 0, got %v", len(items.Data))
	}

	item, err := items.Item("1001")
	if err != nil {
		t.Fatalf("Failed to get item: %v", err)
	}

	if item.Name != "Boots" {
		t.Fatalf("Expected item name to be Boots, got %v", item.Name)
	}
}
