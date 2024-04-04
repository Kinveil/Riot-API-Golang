package summoner_spell

import (
	"encoding/json"
	"testing"
)

type TestStruct struct {
	ID ID `json:"id"`
}

func TestUnmarshalJSON(t *testing.T) {
	expectedResult := TestStruct{ID: ID(4)}

	validIntId := []byte(`{"id": 4}`)
	validStringId := []byte(`{"id": "SummonerFlash"}`)

	// Test case 1: Should correctly unmarshal number id
	var id TestStruct
	err := json.Unmarshal(validIntId, &id)
	if err != nil {
		t.Errorf("Expected no error. Error: %v", err)
	}

	if expectedResult != id {
		t.Errorf("Expected: %v - Got: %v", expectedResult, id)
	}

	// Test case 2: Should correctly unmarshal string id
	var id2 TestStruct
	err = json.Unmarshal(validStringId, &id2)
	if err != nil {
		t.Errorf("Expected no error. Error: %v", err)
	}

	if expectedResult != id2 {
		t.Errorf("Expected: %v - Got: %v", expectedResult, id)
	}
}
