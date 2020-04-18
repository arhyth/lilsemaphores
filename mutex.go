package lilsem

import (
	"context"

	"golang.org/x/sync/semaphore"
)

var (
	mutexWeight int64 = 1
)

type Mutex struct {
	sem *semaphore.Weighted
}

func (m *Mutex) Wait() {
	// sem.Acquire/2 blocks until it successfully acquires semaphore.
	// Since we acquiring or release semaphore total weight
	// in a single call, only 1 goroutine can proceed at a time.
	m.sem.Acquire(context.Background(), mutexWeight)
}

func (m *Mutex) Signal() {
	m.sem.Release(mutexWeight)
}

func NewMutex() *Mutex {
	return &Mutex{
		sem: semaphore.NewWeighted(mutexWeight),
	}
}
