package ratelimiter

import "sync"

type limiterMu struct {
	mutex    sync.Mutex
	cond     *sync.Cond
	capacity int
	current  int
}

func newLimiterMu(capacity int) *limiterMu {
	l := &limiterMu{
		capacity: capacity,
		current:  0,
	}

	l.cond = sync.NewCond(&l.mutex)

	return l
}

func (l *limiterMu) Obtain() {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	for l.current >= l.capacity {
		l.cond.Wait() // Wait until Release() is called
	}

	l.current++
}

func (l *limiterMu) Release() {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	l.current--
	l.cond.Signal() // Wake up one goroutine waiting in Obtain(), if any
}

func (l *limiterMu) SetCapacity(newCapacity int) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	l.capacity = newCapacity

	// Wake up any goroutines that may be waiting in Obtain()
	l.cond.Broadcast()
}

func (l *limiterMu) Capacity() int {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	return l.capacity
}
