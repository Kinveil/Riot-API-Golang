package apiclient_test

import (
	"testing"

	"github.com/junioryono/Riot-API-Golang/src/constants/region"
)

func TestGetStatusPlatformData(t *testing.T) {
	client := newTestClient(t, nil)

	shardStatus, err := client.GetStatusPlatformData(region.NA1)
	if err != nil {
		t.Fatalf("Failed to get shard data: %v", err)
	}

	if shardStatus == nil {
		t.Fatalf("Expected to receive shard status but got nil")
	}
}
