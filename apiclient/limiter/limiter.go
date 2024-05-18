package limiter

import (
	"context"
	"sync"
	"sync/atomic"
	"time"
)

type Limiter struct {
	capacity int32
	current  int32
	mu       sync.Mutex
	waiters  map[chan struct{}]struct{}
}

// NewLimiter initializes a new limiter with the specified capacity
func NewLimiter(capacity int) *Limiter {
	return &Limiter{
		capacity: int32(capacity),
		current:  0,
		waiters:  make(map[chan struct{}]struct{}),
	}
}

// Obtain blocks until a token is available or the context is cancelled
func (l *Limiter) Obtain(ctx context.Context) error {
	waiter := make(chan struct{})

	l.mu.Lock()
	if atomic.LoadInt32(&l.current) < l.capacity {
		atomic.AddInt32(&l.current, 1)
		l.mu.Unlock()
		return nil
	}
	l.waiters[waiter] = struct{}{}
	l.mu.Unlock()

	select {
	case <-ctx.Done():
		l.removeWaiter(waiter)
		return ctx.Err()
	case <-waiter:
		return nil
	}
}

// Release explicitly releases a token
func (l *Limiter) Release() {
	l.mu.Lock()
	defer l.mu.Unlock()
	if atomic.AddInt32(&l.current, -1) < 0 {
		atomic.StoreInt32(&l.current, 0)
	}
	for waiter := range l.waiters {
		delete(l.waiters, waiter)
		close(waiter)
		atomic.AddInt32(&l.current, 1)
		break
	}
}

// ReleaseAfterDelay releases the limiter after the specified delay
func (l *Limiter) ReleaseAfterDelay(delay time.Duration) {
	time.AfterFunc(delay, l.Release)
}

// SetCapacity updates the rate limiter's capacity
func (l *Limiter) SetCapacity(newCapacity int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.capacity = int32(newCapacity)
	for atomic.LoadInt32(&l.current) < l.capacity && len(l.waiters) > 0 {
		for waiter := range l.waiters {
			delete(l.waiters, waiter)
			close(waiter)
			atomic.AddInt32(&l.current, 1)
			break
		}
	}
}

// Capacity returns the current capacity of the rate limiter
func (l *Limiter) Capacity() int {
	return int(atomic.LoadInt32(&l.capacity))
}

// removeWaiter removes a specific waiter from the waiters list
func (l *Limiter) removeWaiter(waiter chan struct{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if _, ok := l.waiters[waiter]; ok {
		delete(l.waiters, waiter)
		close(waiter)
	}
}
