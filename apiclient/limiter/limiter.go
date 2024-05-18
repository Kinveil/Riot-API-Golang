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
	waiters  []chan struct{}
}

// NewLimiter initializes a new limiter with the specified capacity
func NewLimiter(capacity int) *Limiter {
	return &Limiter{
		capacity: int32(capacity),
		current:  0,
		waiters:  make([]chan struct{}, 0),
	}
}

// Obtain blocks until a token is available or the context is cancelled
func (l *Limiter) Obtain(ctx context.Context) error {
	waiter := make(chan struct{})
	defer close(waiter)

	l.mu.Lock()
	if atomic.LoadInt32(&l.current) < l.capacity {
		atomic.AddInt32(&l.current, 1)
		l.mu.Unlock()
		return nil
	}
	l.waiters = append(l.waiters, waiter)
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
	if len(l.waiters) > 0 {
		waiter := l.waiters[0]
		l.waiters = l.waiters[1:]
		close(waiter)
		atomic.AddInt32(&l.current, 1)
	}
}

// ReleaseAfterDelay releases the limiter after the specified delay, respecting the provided context.
func (l *Limiter) ReleaseAfterDelay(delay time.Duration) {
	time.AfterFunc(delay, l.Release)
}

// SetCapacity updates the rate limiter's capacity
func (l *Limiter) SetCapacity(newCapacity int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.capacity = int32(newCapacity)
	for atomic.LoadInt32(&l.current) < l.capacity && len(l.waiters) > 0 {
		waiter := l.waiters[0]
		l.waiters = l.waiters[1:]
		close(waiter)
		atomic.AddInt32(&l.current, 1)
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
	for i, w := range l.waiters {
		if w == waiter {
			l.waiters = append(l.waiters[:i], l.waiters[i+1:]...)
			break
		}
	}
}
