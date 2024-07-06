package ratelimiter

import (
	"context"
	"time"

	"github.com/Kinveil/Riot-API-Golang/apiclient/limiter"
)

func newRateLimit() *RateLimit {
	return &RateLimit{
		shortLimiter: limiter.NewLimiter(initialRegionLimit),
		longLimiter:  limiter.NewLimiter(initialRegionLimit),
		blockedUntil: time.Time{},
	}
}

func (rl *RateLimiter) getRegionLimiter(region string) *RateLimit {
	rl.regionMutex.Lock()
	defer rl.regionMutex.Unlock()

	if limiter, ok := rl.regionLimiters[region]; ok {
		return limiter
	}

	limiter := newRateLimit()
	rl.regionLimiters[region] = limiter
	return limiter
}

func (rl *RateLimiter) getMethodLimiter(method string) *RateLimit {
	rl.methodMutex.Lock()
	defer rl.methodMutex.Unlock()

	if limiter, ok := rl.methodLimiters[method]; ok {
		return limiter
	}

	limiter := newRateLimit()
	rl.methodLimiters[method] = limiter
	return limiter
}

func (rl *RateLimiter) waitForLimiters(ctx context.Context, regionLimiter, methodLimiter *RateLimit, isRetryRequest bool) error {
	var lCtx context.Context
	if ctx == nil {
		lCtx = context.Background()
	} else {
		lCtx = ctx
	}

	if err := rl.waitForBlock(lCtx, regionLimiter.blockedUntil); err != nil {
		return err
	}

	if err := rl.waitForBlock(lCtx, methodLimiter.blockedUntil); err != nil {
		return err
	}

	// If this is a not a retry request, obtain tokens from all limiters
	// We do this because retry requests do not release tokens
	if !isRetryRequest {
		if err := regionLimiter.shortLimiter.Obtain(lCtx); err != nil {
			return err
		}

		if err := regionLimiter.longLimiter.Obtain(lCtx); err != nil {
			return err
		}

		if err := methodLimiter.shortLimiter.Obtain(lCtx); err != nil {
			return err
		}
	}

	return nil
}

func (rl *RateLimiter) waitForBlock(ctx context.Context, blockedUntil time.Time) error {
	if time.Now().Before(blockedUntil) {
		timeout := time.Until(blockedUntil)
		waitCtx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()

		<-waitCtx.Done()
		if waitCtx.Err() != nil && waitCtx.Err() != context.DeadlineExceeded {
			return waitCtx.Err()
		}
	}

	return nil
}

func (rl *RateLimiter) releaseLimiters(regionLimiter, methodLimiter *RateLimit) {
	regionLimiter.shortLimiter.Release()
	regionLimiter.longLimiter.Release()
	methodLimiter.shortLimiter.Release()
}

func (rl *RateLimiter) releaseLimitersAfterDelay(regionLimiter, methodLimiter *RateLimit, delay time.Duration) {
	<-time.After(delay)
	rl.releaseLimiters(regionLimiter, methodLimiter)
}
