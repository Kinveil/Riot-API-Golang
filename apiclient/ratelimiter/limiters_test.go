package ratelimiter

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestConcurrentRegionLimiters(t *testing.T) {
	rl := &RateLimiter{
		regionLimiters: make(map[string]*RateLimit),
	}
	var wg sync.WaitGroup

	regions := []string{"NA", "EUW", "EUNE", "KR", "JP"}

	for _, region := range regions {
		wg.Add(1)
		go func(region string) {
			defer wg.Done()
			for i := 0; i < 100; i++ {
				limiter := rl.getRegionLimiter(region)
				assert.NotNil(t, limiter)
			}
		}(region)
	}

	wg.Wait()
}

func TestConcurrentMethodLimiters(t *testing.T) {
	rl := &RateLimiter{
		methodLimiters: make(map[string]*RateLimit),
	}
	var wg sync.WaitGroup

	methods := []string{"GetSummoner", "GetMatch", "GetChampion", "GetLeague"}

	for _, method := range methods {
		wg.Add(1)
		go func(method string) {
			defer wg.Done()
			for i := 0; i < 100; i++ {
				limiter := rl.getMethodLimiter(method)
				assert.NotNil(t, limiter)
			}
		}(method)
	}

	wg.Wait()
}

func TestSetCapacity(t *testing.T) {
	limiter := newRateLimit()
	limiter.shortLimiter.SetCapacity(500000)
	limiter.longLimiter.SetCapacity(500000)

	assert.Equal(t, 500000, limiter.shortLimiter.Capacity())
	assert.Equal(t, 500000, limiter.longLimiter.Capacity())
}

func TestConcurrentRegionAndMethodLimiters(t *testing.T) {
	rl := &RateLimiter{
		regionLimiters: make(map[string]*RateLimit),
		methodLimiters: make(map[string]*RateLimit),
	}
	var wg sync.WaitGroup

	regions := []string{"NA", "EUW", "EUNE", "KR", "JP"}
	methods := []string{"GetSummoner", "GetMatch", "GetChampion", "GetLeague"}

	for _, region := range regions {
		regionLimiter := rl.getRegionLimiter(region)
		regionLimiter.shortLimiter.SetCapacity(5)
		regionLimiter.longLimiter.SetCapacity(300)

		for _, method := range methods {
			methodLimiter := rl.getMethodLimiter(region + method)
			methodLimiter.shortLimiter.SetCapacity(20)
		}
	}

	for _, region := range regions {
		for _, method := range methods {
			wg.Add(1)
			go func(region, method string) {
				defer wg.Done()
				for i := 0; i < 1000; i++ {
					regionLimiter := rl.getRegionLimiter(region)
					methodLimiter := rl.getMethodLimiter(region + method)

					rl.waitForLimiters(regionLimiter, methodLimiter, context.Background())
					rl.releaseLimitersAfterDelay(regionLimiter, methodLimiter, time.Millisecond*1)
				}
			}(region, method)
		}
	}

	wg.Wait()
}

func TestConcurrentOneRegionAndTwoMethodsLimiters(t *testing.T) {
	rl := &RateLimiter{
		regionLimiters: make(map[string]*RateLimit),
		methodLimiters: make(map[string]*RateLimit),
	}
	var wg sync.WaitGroup

	regions := []string{"NA"}
	methods := []string{"GetSummoner", "GetMatch"}

	startTime := time.Now()

	for _, region := range regions {
		regionLimiter := rl.getRegionLimiter(region)
		regionLimiter.shortLimiter.SetCapacity(1)
		regionLimiter.longLimiter.SetCapacity(1)

		for _, method := range methods {
			methodLimiter := rl.getMethodLimiter(region + method)
			methodLimiter.shortLimiter.SetCapacity(1)
		}
	}

	for _, region := range regions {
		for _, method := range methods {
			wg.Add(1)
			go func(region, method string) {
				defer wg.Done()

				regionLimiter := rl.getRegionLimiter(region)
				methodLimiter := rl.getMethodLimiter(region + method)

				rl.waitForLimiters(regionLimiter, methodLimiter, context.Background())
				rl.releaseLimitersAfterDelay(regionLimiter, methodLimiter, time.Second*2)
			}(region, method)
		}
	}

	wg.Wait()

	elapsedTime := time.Since(startTime)
	assert.GreaterOrEqual(t, elapsedTime.Milliseconds(), int64(4000))
	assert.LessOrEqual(t, elapsedTime.Milliseconds(), int64(5000))
}

func TestConcurrentTwoRegionsAndOneMethodLimiters(t *testing.T) {
	rl := &RateLimiter{
		regionLimiters: make(map[string]*RateLimit),
		methodLimiters: make(map[string]*RateLimit),
	}
	var wg sync.WaitGroup

	regions := []string{"NA", "EUW"}
	methods := []string{"GetSummoner"}

	startTime := time.Now()

	for _, region := range regions {
		regionLimiter := rl.getRegionLimiter(region)
		regionLimiter.shortLimiter.SetCapacity(1)
		regionLimiter.longLimiter.SetCapacity(1)

		for _, method := range methods {
			methodLimiter := rl.getMethodLimiter(region + method)
			methodLimiter.shortLimiter.SetCapacity(1)
		}
	}

	for _, region := range regions {
		for _, method := range methods {
			wg.Add(1)
			go func(region, method string) {
				defer wg.Done()

				regionLimiter := rl.getRegionLimiter(region)
				methodLimiter := rl.getMethodLimiter(region + method)

				rl.waitForLimiters(regionLimiter, methodLimiter, context.Background())
				rl.releaseLimitersAfterDelay(regionLimiter, methodLimiter, time.Second*2)
			}(region, method)
		}
	}

	wg.Wait()

	elapsedTime := time.Since(startTime)
	assert.GreaterOrEqual(t, elapsedTime.Milliseconds(), int64(2000))
	assert.LessOrEqual(t, elapsedTime.Milliseconds(), int64(3000))
}

func TestConcurrentWaitForLimiters(t *testing.T) {
	rl := &RateLimiter{}
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			regionLimiter := newRateLimit()
			methodLimiter := newRateLimit()

			ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
			defer cancel()

			err := rl.waitForLimiters(regionLimiter, methodLimiter, ctx)
			assert.NoError(t, err)
		}()
	}

	wg.Wait()
}

func TestConcurrentReleaseLimiters(t *testing.T) {
	rl := &RateLimiter{}
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			regionLimiter := newRateLimit()
			methodLimiter := newRateLimit()

			rl.releaseLimiters(regionLimiter, methodLimiter)
		}()
	}

	wg.Wait()
}

func TestConcurrentReleaseLimitersAfterDelay(t *testing.T) {
	rl := &RateLimiter{}
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			regionLimiter := newRateLimit()
			methodLimiter := newRateLimit()

			go rl.releaseLimitersAfterDelay(regionLimiter, methodLimiter, time.Millisecond*50)
			time.Sleep(time.Millisecond * 100)
		}()
	}

	wg.Wait()
}
