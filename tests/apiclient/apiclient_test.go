package apiclient_test

import (
	"fmt"
	"sync"
	"testing"

	"github.com/junioryono/Riot-API-Golang/src/apiclient"
	"github.com/junioryono/Riot-API-Golang/src/apiclient/ratelimiter"
	"github.com/junioryono/Riot-API-Golang/src/constants/region"
)

func newTestClient(t *testing.T, conserveUsage *ratelimiter.ConserveUsage) apiclient.Client {
	key := "RGAPI-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
	client := apiclient.New(key)

	if conserveUsage != nil {
		client.SetUsageConservation(*conserveUsage)
	}

	if client == nil {
		t.Fatalf("Expected client to be not nil")
	}

	return client
}

func TestRateLimiter(t *testing.T) {
	client := newTestClient(t, nil)

	var wg sync.WaitGroup
	errCh := make(chan error, 1)

	for i := 0; i < 230; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			_, err := client.GetSummonerBySummonerName(region.NA1, "Mighty Junior")
			if err != nil {
				errCh <- fmt.Errorf("Failed to get summoner on iteration %d: %v", i, err)
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(errCh) // Close channel after all goroutines are done
	}()

	if err := <-errCh; err != nil {
		t.Fatalf("%v", err) // This will fail the test if there is any error
	}
}
