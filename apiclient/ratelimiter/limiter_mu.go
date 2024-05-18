package ratelimiter

import (
	"sync"
	"sync/atomic"
)

type limiterMu struct {
	capacity int32
	current  int32
	cond     *sync.Cond
}

// newLimiterMu initializes a new limiter with the specified capacity
func newLimiterMu(capacity int) *limiterMu {
	mu := &sync.Mutex{}
	return &limiterMu{
		capacity: int32(capacity),
		current:  0,
		cond:     sync.NewCond(mu),
	}
}

// Obtain blocks until a token is available
func (l *limiterMu) Obtain() {
	l.cond.L.Lock()
	defer l.cond.L.Unlock()

	for atomic.LoadInt32(&l.current) >= l.capacity {
		l.cond.Wait()
	}

	atomic.AddInt32(&l.current, 1)
}

// Release explicitly releases a token
func (l *limiterMu) Release() {
	l.cond.L.Lock()
	defer l.cond.L.Unlock()

	if atomic.AddInt32(&l.current, -1) < l.capacity {
		l.cond.Signal()
	}
}

// SetCapacity updates the rate limiter's capacity
func (l *limiterMu) SetCapacity(newCapacity int) {
	l.cond.L.Lock()
	defer l.cond.L.Unlock()

	l.capacity = int32(newCapacity)
	l.cond.Broadcast()
}

// Capacity returns the current capacity of the rate limiter
func (l *limiterMu) Capacity() int {
	return int(atomic.LoadInt32(&l.capacity))
}
