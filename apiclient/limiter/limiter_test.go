package limiter

import (
	"context"
	"sync"
	"testing"
	"time"
)

func TestLimiter_Obtain_Release(t *testing.T) {
	limiter := NewLimiter(10)

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := limiter.Obtain(context.Background()); err != nil {
				t.Errorf("Failed to obtain token: %v", err)
			}
			time.Sleep(time.Millisecond)
			limiter.Release()
		}()
	}

	wg.Wait()
}

func TestLimiter_Obtain_Release_Concurrent(t *testing.T) {
	limiter := NewLimiter(10)

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := limiter.Obtain(context.Background()); err != nil {
				t.Errorf("Failed to obtain token: %v", err)
			}
			time.Sleep(time.Millisecond)
			limiter.Release()
		}()
	}

	wg.Wait()
}

func TestLimiter_Obtain_Release_MultipleGoroutines(t *testing.T) {
	limiter := NewLimiter(10)

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				if err := limiter.Obtain(context.Background()); err != nil {
					t.Errorf("Failed to obtain token: %v", err)
				}
				time.Sleep(time.Millisecond)
				limiter.Release()
			}
		}()
	}

	wg.Wait()
}

func TestLimiter_Obtain_Release_Timeout(t *testing.T) {
	limiter := NewLimiter(1)

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	if err := limiter.Obtain(ctx); err != nil {
		t.Fatalf("Failed to obtain token: %v", err)
	}

	if err := limiter.Obtain(ctx); err != context.DeadlineExceeded {
		t.Fatalf("Expected context deadline exceeded error, got: %v", err)
	}
}

func TestLimiter_SetCapacity(t *testing.T) {
	limiter := NewLimiter(1)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := limiter.Obtain(context.Background()); err != nil {
				t.Errorf("Failed to obtain token: %v", err)
			}
			time.Sleep(time.Millisecond)
			limiter.Release()
		}()
	}

	time.Sleep(5 * time.Millisecond)
	limiter.SetCapacity(10)
	wg.Wait()
}

func TestLimiter_ReleaseAfterDelay(t *testing.T) {
	limiter := NewLimiter(1)

	if err := limiter.Obtain(context.Background()); err != nil {
		t.Fatalf("Failed to obtain token: %v", err)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := limiter.Obtain(context.Background()); err != nil {
			t.Errorf("Failed to obtain token: %v", err)
		}
		limiter.Release()
	}()

	limiter.ReleaseAfterDelay(10 * time.Millisecond)
	wg.Wait()
}

func TestLimiter_MultipleLimiters(t *testing.T) {
	limiter1 := NewLimiter(5)
	limiter2 := NewLimiter(10)

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := limiter1.Obtain(context.Background()); err != nil {
				t.Errorf("Failed to obtain token from limiter1: %v", err)
			}
			if err := limiter2.Obtain(context.Background()); err != nil {
				t.Errorf("Failed to obtain token from limiter2: %v", err)
			}
			time.Sleep(time.Millisecond)
			limiter1.Release()
			limiter2.Release()
		}()
	}

	wg.Wait()
}
