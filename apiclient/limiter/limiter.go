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
	for {
		if atomic.LoadInt32(&l.current) < atomic.LoadInt32(&l.capacity) {
			if atomic.CompareAndSwapInt32(&l.current, atomic.LoadInt32(&l.current), atomic.LoadInt32(&l.current)+1) {
				return nil
			}

			continue
		}

		w := make(chan struct{})

		l.mu.Lock()
		l.waiters = append(l.waiters, w)
		l.mu.Unlock()

		select {
		case <-ctx.Done():
			l.removeWaiter(w)
			return ctx.Err()
		case <-w:
			return nil
		}
	}
}

// Release explicitly releases a token
func (l *Limiter) Release() {
	var w chan struct{}

	l.mu.Lock()

	if len(l.waiters) > 0 {
		w = l.waiters[0]
		l.waiters = l.waiters[1:]
	} else {
		atomic.AddInt32(&l.current, -1)
	}

	l.mu.Unlock()

	if w != nil {
		close(w)
	}
}

// ReleaseAfterDelay releases the limiter after the specified delay
func (l *Limiter) ReleaseAfterDelay(delay time.Duration) {
	time.AfterFunc(delay, l.Release)
}

// SetCapacity updates the rate limiter's capacity
func (l *Limiter) SetCapacity(newCapacity int) {
	newCapacityInt32 := int32(newCapacity)
	oldCapacity := atomic.LoadInt32(&l.capacity)

	if newCapacityInt32 == oldCapacity {
		// Capacity unchanged, no action needed
		return
	}

	if newCapacityInt32 < oldCapacity {
		// If capacity decreased, we can handle this without locking
		atomic.StoreInt32(&l.capacity, newCapacityInt32)

		for {
			current := atomic.LoadInt32(&l.current)
			if current <= newCapacityInt32 {
				break
			}

			if atomic.CompareAndSwapInt32(&l.current, current, newCapacityInt32) {
				break
			}
		}

		return
	}

	// If capacity increased, we need to lock to handle waiters
	l.mu.Lock()

	// Check again in case capacity changed while waiting for lock
	oldCapacity = atomic.LoadInt32(&l.capacity)
	if newCapacityInt32 <= oldCapacity {
		l.mu.Unlock()
		return
	}

	atomic.StoreInt32(&l.capacity, newCapacityInt32)

	additionalCapacity := int(newCapacityInt32 - oldCapacity)
	current := int(atomic.LoadInt32(&l.current))

	toRelease := min(additionalCapacity, len(l.waiters))
	toRelease = min(toRelease, int(newCapacityInt32)-current)

	var waitersToRelease []chan struct{}
	for i := 0; i < toRelease; i++ {
		w := l.waiters[0]
		l.waiters = l.waiters[1:]
		waitersToRelease = append(waitersToRelease, w)
		atomic.AddInt32(&l.current, 1)
	}

	l.mu.Unlock()

	for _, w := range waitersToRelease {
		close(w)
	}
}

// Capacity returns the current capacity of the rate limiter
func (l *Limiter) Capacity() int {
	return int(atomic.LoadInt32(&l.capacity))
}

// removeWaiter removes a specific waiter from the waiters list
func (l *Limiter) removeWaiter(ch chan struct{}) {
	var waiterToRelease chan struct{}

	l.mu.Lock()
	for i, w := range l.waiters {
		if w == ch {
			l.waiters = append(l.waiters[:i], l.waiters[i+1:]...)
			waiterToRelease = w
			break
		}
	}
	l.mu.Unlock()

	if waiterToRelease != nil {
		close(waiterToRelease)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
