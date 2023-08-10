package staticdata_test

import (
	"testing"

	"github.com/junioryono/Riot-API-Golang/src/staticdata"
)

func TestGetPatches(t *testing.T) {
	patches, err := staticdata.GetPatches()
	if err != nil {
		t.Fatalf("Failed to get patches: %v", err)
	}

	if len(patches) == 0 {
		t.Error("Failed to get patches: no patches returned")
	}
}

func TestGetPatch(t *testing.T) {
	patches, err := staticdata.GetPatches()
	if err != nil {
		t.Fatalf("Failed to get patches: %v", err)
	}

	if len(patches) == 0 {
		t.Error("Failed to get patches: no patches returned")
	}

	currentPatch := patches.CurrentPatch()

	if currentPatch != patches[0] {
		t.Errorf("Failed to get current patch: expected %v, got %v", patches[0], currentPatch)
	}
}

func TestGetPatchesWithStartTime(t *testing.T) {
	patches, err := staticdata.GetPatchesWithStartTime()
	if err != nil {
		t.Fatalf("Failed to get patches: %v", err)
	}

	if len(patches) == 0 {
		t.Error("Failed to get patches: no patches returned")
	}
}

func TestGetPatchWithStartTime(t *testing.T) {
	patches, err := staticdata.GetPatchesWithStartTime()
	if err != nil {
		t.Fatalf("Failed to get patches: %v", err)
	}

	if len(patches) == 0 {
		t.Error("Failed to get patches: no patches returned")
	}

	currentPatch := patches.CurrentPatch()

	if currentPatch.Patch != patches[len(patches)-1].Patch {
		t.Errorf("Failed to get current patch: expected %v, got %v", patches[0], currentPatch)
	}
}
