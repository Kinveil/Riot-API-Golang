package staticdata_test

import (
	"testing"

	"github.com/junioryono/Riot-API-Golang/src/constants/language"
	"github.com/junioryono/Riot-API-Golang/src/staticdata"
)

func TestGetQueues(t *testing.T) {
	currentPatch := getCurrentPatch(t)

	queues, err := staticdata.GetQueues(currentPatch, language.EnglishUnitedStates)
	if err != nil {
		t.Fatalf("Failed to get queues: %v", err)
	}

	if len(queues) == 0 {
		t.Fatalf("Expected champions to be greater than 0, got %v", len(queues))
	}
}

func TestGetQueue(t *testing.T) {
	currentPatch := getCurrentPatch(t)

	queues, err := staticdata.GetQueues(currentPatch, language.EnglishUnitedStates)
	if err != nil {
		t.Fatalf("Failed to get queues: %v", err)
	}

	if len(queues) == 0 {
		t.Fatalf("Expected queues to be greater than 0, got %v", len(queues))
	}

	queue, err := queues.Queue(420)
	if err != nil {
		t.Fatalf("Failed to get queue: %v", err)
	}

	if queue.Map != "Summoner's Rift" {
		t.Fatalf("Expected queue map to be Summoner's Rift, got %v", queue.Map)
	}
}
