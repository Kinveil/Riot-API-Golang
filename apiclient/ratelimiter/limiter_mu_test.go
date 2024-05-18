package ratelimiter

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestLimiterMu_HighConcurrency(t *testing.T) {
	const numGoroutines = 1000
	const numTokens = 100
	limiter := newLimiterMu(numTokens)

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	successObtains := int32(0)
	successReleases := int32(0)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			limiter.Obtain()
			atomic.AddInt32(&successObtains, 1)
			time.Sleep(10 * time.Millisecond) // Simulate some work
			limiter.Release()
			atomic.AddInt32(&successReleases, 1)
		}()
	}
	wg.Wait()

	if successObtains != int32(numGoroutines) || successReleases != int32(numGoroutines) {
		t.Fatalf("Expected %d successful obtains and releases, got %d obtains and %d releases", numGoroutines, successObtains, successReleases)
	}
}

func TestLimiterMu_RateLimiting(t *testing.T) {
	const numGoroutines = 100
	const numTokens = 10
	limiter := newLimiterMu(numTokens)

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	successObtains := int32(0)
	failedObtains := int32(0)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			if atomic.LoadInt32(&limiter.current) < limiter.capacity {
				limiter.Obtain()
				atomic.AddInt32(&successObtains, 1)
				time.Sleep(10 * time.Millisecond) // Simulate some work
				limiter.Release()
			} else {
				atomic.AddInt32(&failedObtains, 1)
			}
		}()
	}
	wg.Wait()

	if successObtains+failedObtains != int32(numGoroutines) {
		t.Fatalf("Expected %d attempts to obtain tokens, got %d", numGoroutines, successObtains+failedObtains)
	}
}

func TestLimiterMu_SetCapacity(t *testing.T) {
	const initialCapacity = 10
	const newCapacity = 5
	limiter := newLimiterMu(initialCapacity)

	for i := 0; i < initialCapacity; i++ {
		limiter.Obtain()
	}
	if atomic.LoadInt32(&limiter.current) != initialCapacity {
		t.Fatalf("Expected %d tokens to be obtained, got %d", initialCapacity, atomic.LoadInt32(&limiter.current))
	}

	limiter.SetCapacity(newCapacity)
	if limiter.Capacity() != newCapacity {
		t.Fatalf("Expected capacity to be %d, got %d", newCapacity, limiter.Capacity())
	}

	// Release some tokens
	for i := 0; i < initialCapacity-newCapacity; i++ {
		limiter.Release()
	}

	if atomic.LoadInt32(&limiter.current) != newCapacity {
		t.Fatalf("Expected %d tokens to remain after releasing, got %d", newCapacity, atomic.LoadInt32(&limiter.current))
	}
}

func TestLimiterMu_HighVolume(t *testing.T) {
	const numGoroutines = 1000
	const numTokens = 100
	limiter := newLimiterMu(numTokens)

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	successCount := int32(0)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for {
				if atomic.LoadInt32(&limiter.current) < limiter.capacity {
					limiter.Obtain()
					atomic.AddInt32(&successCount, 1)
					time.Sleep(10 * time.Millisecond) // Simulate some work
					limiter.Release()
					break
				}
				time.Sleep(10 * time.Millisecond) // Wait before retrying to obtain a token
			}
		}()
	}
	wg.Wait()

	if successCount != int32(numGoroutines) {
		t.Fatalf("Expected %d successful obtains, got %d", numGoroutines, successCount)
	}
}
