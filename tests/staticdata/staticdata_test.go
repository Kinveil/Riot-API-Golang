package staticdata_test

import (
	"testing"

	"github.com/junioryono/Riot-API-Golang/src/constants/patch"
	"github.com/junioryono/Riot-API-Golang/src/staticdata"
)

func getCurrentPatch(t *testing.T) patch.Patch {
	patches, err := staticdata.GetPatches()
	if err != nil {
		t.Fatalf("Failed to get patches: %v", err)
	}

	return patches.CurrentPatch()
}
