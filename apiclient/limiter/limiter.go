package limiter

import (
	"container/heap"
	"context"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

type Limiter struct {
	capacity int32
	current  int32
	mu       sync.Mutex
	waiters  priorityQueue
}

type waiter struct {
	ch       chan struct{}
	priority int
	index    int
}

type priorityQueue []*waiter

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*waiter)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

// NewLimiter initializes a new limiter with the specified capacity
func NewLimiter(capacity int) *Limiter {
	return &Limiter{
		capacity: int32(capacity),
		current:  0,
		waiters:  make(priorityQueue, 0),
	}
}

// Obtain blocks until a token is available or the context is cancelled
func (l *Limiter) Obtain(ctx context.Context, priority int) error {
	for {
		if atomic.LoadInt32(&l.current) < atomic.LoadInt32(&l.capacity) {
			if atomic.CompareAndSwapInt32(&l.current, atomic.LoadInt32(&l.current), atomic.LoadInt32(&l.current)+1) {
				return nil
			}
			continue
		}

		w := &waiter{
			ch:       make(chan struct{}),
			priority: priority,
		}

		l.mu.Lock()
		heap.Push(&l.waiters, w)
		l.mu.Unlock()

		select {
		case <-ctx.Done():
			l.removeWaiter(w)
			return ctx.Err()
		case <-w.ch:
			return nil
		}
	}
}

// Release explicitly releases a token
func (l *Limiter) Release() {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.waiters.Len() > 0 {
		w := heap.Pop(&l.waiters).(*waiter)
		close(w.ch)
	} else {
		atomic.AddInt32(&l.current, -1)
	}
}

// ReleaseAfterDelay releases the limiter after the specified delay
func (l *Limiter) ReleaseAfterDelay(delay time.Duration) {
	time.AfterFunc(delay, l.Release)
}

// SetCapacity updates the rate limiter's capacity
func (l *Limiter) SetCapacity(newCapacity int) {
	if newCapacity < 0 {
		return
	}

	if newCapacity > math.MaxInt32 {
		return
	}

	newCapacityInt32 := int32(newCapacity)
	oldCapacity := atomic.LoadInt32(&l.capacity)

	if newCapacityInt32 == oldCapacity {
		return
	}

	if newCapacityInt32 < oldCapacity {
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

	l.mu.Lock()
	defer l.mu.Unlock()

	oldCapacity = atomic.LoadInt32(&l.capacity)
	if newCapacityInt32 <= oldCapacity {
		return
	}

	atomic.StoreInt32(&l.capacity, newCapacityInt32)

	additionalCapacity := int(newCapacityInt32 - oldCapacity)
	current := int(atomic.LoadInt32(&l.current))

	toRelease := min(additionalCapacity, l.waiters.Len())
	toRelease = min(toRelease, int(newCapacityInt32)-current)

	for i := 0; i < toRelease; i++ {
		w := heap.Pop(&l.waiters).(*waiter)
		close(w.ch)
		atomic.AddInt32(&l.current, 1)
	}
}

// Capacity returns the current capacity of the rate limiter
func (l *Limiter) Capacity() int {
	return int(atomic.LoadInt32(&l.capacity))
}

// removeWaiter removes a specific waiter from the waiters list
func (l *Limiter) removeWaiter(w *waiter) {
	l.mu.Lock()
	defer l.mu.Unlock()

	for i := 0; i < l.waiters.Len(); i++ {
		if l.waiters[i] == w {
			heap.Remove(&l.waiters, i)
			close(w.ch)
			break
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
