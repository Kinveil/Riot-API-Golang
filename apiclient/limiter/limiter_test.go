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
			if err := limiter.Obtain(context.Background(), 1); err != nil {
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
			if err := limiter.Obtain(context.Background(), 1); err != nil {
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
				if err := limiter.Obtain(context.Background(), 1); err != nil {
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

	if err := limiter.Obtain(ctx, 1); err != nil {
		t.Fatalf("Failed to obtain token: %v", err)
	}

	if err := limiter.Obtain(ctx, 1); err != context.DeadlineExceeded {
		t.Fatalf("Expected context deadline exceeded error, got: %v", err)
	}
}

func TestLimiter_SetCapacity(t *testing.T) {
	limiter := NewLimiter(1)

	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := limiter.Obtain(context.Background(), 1); err != nil {
				t.Errorf("Failed to obtain token: %v", err)
			}
			time.Sleep(time.Millisecond)
			limiter.Release()
		}()
	}

	time.Sleep(5 * time.Millisecond)
	limiter.SetCapacity(10)

	time.Sleep(5 * time.Millisecond)
	limiter.SetCapacity(1000)

	wg.Wait()
}

func TestLimiter_ReleaseAfterDelay(t *testing.T) {
	limiter := NewLimiter(1)

	if err := limiter.Obtain(context.Background(), 1); err != nil {
		t.Fatalf("Failed to obtain token: %v", err)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := limiter.Obtain(context.Background(), 1); err != nil {
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
			if err := limiter1.Obtain(context.Background(), 1); err != nil {
				t.Errorf("Failed to obtain token from limiter1: %v", err)
			}
			if err := limiter2.Obtain(context.Background(), 1); err != nil {
				t.Errorf("Failed to obtain token from limiter2: %v", err)
			}
			time.Sleep(time.Millisecond)
			limiter1.Release()
			limiter2.Release()
		}()
	}

	wg.Wait()
}

func TestLimiter_FIFO_Order(t *testing.T) {
	limiter := NewLimiter(200)

	// Obtain the only available token
	if err := limiter.Obtain(context.Background(), 1); err != nil {
		t.Fatalf("Failed to obtain initial token: %v", err)
	}

	const waiters = 500
	orders := make(chan int, waiters)
	starts := make([]chan struct{}, waiters)
	for i := range starts {
		starts[i] = make(chan struct{})
	}

	var wg sync.WaitGroup
	for i := 0; i < waiters; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			<-starts[index] // Wait for signal to start
			if err := limiter.Obtain(context.Background(), 1); err != nil {
				t.Errorf("Waiter %d failed to obtain token: %v", index, err)
				return
			}
			orders <- index
			time.Sleep(time.Millisecond) // Hold the token briefly
			limiter.Release()
		}(i)
	}

	// Start goroutines in a specific order
	for i := 0; i < waiters; i++ {
		close(starts[i])
		time.Sleep(time.Millisecond) // Ensure previous goroutine has time to call Obtain()
	}

	// Release the initial token and let the waiters proceed
	limiter.Release()

	wg.Wait()
	close(orders)

	// Check if the orders are in FIFO sequence
	expected := 0
	for order := range orders {
		if order != expected {
			t.Errorf("Expected order %d, but got %d", expected, order)
		}
		expected++
	}
}

func TestLimiter_FIFO_WithCancellation(t *testing.T) {
	limiter := NewLimiter(200)

	// Obtain the only available token
	if err := limiter.Obtain(context.Background(), 1); err != nil {
		t.Fatalf("Failed to obtain initial token: %v", err)
	}

	const waiters = 500
	orders := make(chan int, waiters)
	starts := make([]chan struct{}, waiters)
	for i := range starts {
		starts[i] = make(chan struct{})
	}
	contexts := make([]context.Context, waiters)
	cancels := make([]context.CancelFunc, waiters)

	var wg sync.WaitGroup
	for i := 0; i < waiters; i++ {
		contexts[i], cancels[i] = context.WithCancel(context.Background())
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			<-starts[index] // Wait for signal to start
			if err := limiter.Obtain(contexts[index], 1); err != nil {
				if err != context.Canceled {
					t.Errorf("Waiter %d failed with unexpected error: %v", index, err)
				}
				return
			}
			orders <- index
			time.Sleep(time.Millisecond) // Hold the token briefly
			limiter.Release()
		}(i)
	}

	// Start goroutines in a specific order
	for i := 0; i < waiters; i++ {
		close(starts[i])
		time.Sleep(time.Millisecond) // Ensure previous goroutine has time to call Obtain()
	}

	// Cancel the middle waiter
	cancels[2]()

	time.Sleep(10 * time.Millisecond) // Give time for cancellation to take effect

	// Release the initial token and let the waiters proceed
	limiter.Release()

	wg.Wait()
	close(orders)

	// Check if the orders are in FIFO sequence, skipping the cancelled waiter
	expected := 0
	for order := range orders {
		if order != expected {
			t.Errorf("Expected order %d, but got %d", expected, order)
		}

		expected++
	}
}

func TestLimiter_Obtain_Priority_Concurrent(t *testing.T) {
	limiter := NewLimiter(1)

	const concurrentRequests = 100
	results := make(chan int, concurrentRequests)

	var wg sync.WaitGroup
	for i := 0; i < concurrentRequests; i++ {
		wg.Add(1)
		go func(priority int) {
			defer wg.Done()
			err := limiter.Obtain(context.Background(), priority)
			if err != nil {
				t.Errorf("Failed to obtain token: %v", err)
				return
			}
			results <- priority
			time.Sleep(time.Millisecond) // Hold the token briefly
			limiter.Release()
		}(i % 5) // Use priorities 0-4
	}

	wg.Wait()
	close(results)

	// Check if higher priorities were generally served first
	var priorities []int
	for p := range results {
		priorities = append(priorities, p)
	}

	if !isDescending(priorities) {
		t.Errorf("Priorities were not generally served in descending order: %v", priorities)
	}
}

func TestLimiter_Obtain_Priority_Fairness(t *testing.T) {
	const (
		highPriority    = 10
		lowPriority     = 1
		requestsPerType = 100
		iterations      = 1
	)

	limiter := NewLimiter(1)

	var (
		highPriorityCount int32
		lowPriorityCount  int32
		mu                sync.Mutex
	)

	var wg sync.WaitGroup
	for i := 0; i < requestsPerType*3; i++ {
		wg.Add(1)
		go func(isHighPriority bool) {
			defer wg.Done()
			priority := lowPriority
			if isHighPriority {
				priority = highPriority
			}

			err := limiter.Obtain(context.Background(), priority)
			if err != nil {
				t.Errorf("Failed to obtain token: %v", err)
				return
			}

			mu.Lock()
			if isHighPriority {
				highPriorityCount++
			} else {
				lowPriorityCount++
			}
			mu.Unlock()

			time.Sleep(time.Millisecond) // Hold the token longer
			limiter.Release()
		}(i%3 != 2) // 2/3 high priority, 1/3 low priority
	}

	wg.Wait()

	highPercentage := float64(highPriorityCount) / float64(highPriorityCount+lowPriorityCount) * 100
	t.Logf("High priority percentage: %.2f%% (High: %d, Low: %d)", highPercentage, highPriorityCount, lowPriorityCount)

	// Check if high priority requests were served significantly more often
	if highPercentage < 60 { // Expecting at least 60% high priority
		t.Errorf("Expected significantly more high priority requests to be served. High: %d, Low: %d", highPriorityCount, lowPriorityCount)
	}
}

func TestLimiter_Obtain_Priority_Ordering(t *testing.T) {
	limiter := NewLimiter(1)

	// Obtain the only available token
	if err := limiter.Obtain(context.Background(), 1); err != nil {
		t.Fatalf("Failed to obtain initial token: %v", err)
	}

	const waiters = 100
	priorities := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3} // Repeat this pattern
	results := make(chan int, waiters)

	var wg sync.WaitGroup
	for i := 0; i < waiters; i++ {
		wg.Add(1)
		go func(priority int) {
			defer wg.Done()
			err := limiter.Obtain(context.Background(), priority)
			if err != nil {
				t.Errorf("Failed to obtain token: %v", err)
				return
			}
			results <- priority
			time.Sleep(time.Millisecond) // Hold the token briefly
			limiter.Release()
		}(priorities[i%len(priorities)])
	}

	// Release the initial token and let the waiters proceed
	limiter.Release()

	wg.Wait()
	close(results)

	// Check if the priorities are in descending order
	var obtainedPriorities []int
	for p := range results {
		obtainedPriorities = append(obtainedPriorities, p)
	}

	if !isDescending(obtainedPriorities) {
		t.Errorf("Priorities were not served in descending order: %v", obtainedPriorities)
	}
}

func isDescending(s []int) bool {
	// Do not test the first 5 elements because they are not guaranteed to be in order
	for i := 5; i < len(s); i++ {
		if s[i-1] < s[i] {
			return false
		}
	}

	return true
}
