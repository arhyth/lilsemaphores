package lilsem

import (
	"context"

	"golang.org/x/sync/semaphore"
)

type Multiplex struct {
	sem *semaphore.Weighted
}

func (m *Multiplex) Wait() {
	// Every time it is called with mutexWeight (1) and thus limits the number
	// of goroutines in the critical section to the total weight provided
	m.sem.Acquire(context.Background(), mutexWeight)
}

func (m *Multiplex) Signal() {
	m.sem.Release(mutexWeight)
}

func NewMultiplex(weight int64) *Multiplex {
	return &Multiplex{
		sem: semaphore.NewWeighted(weight),
	}
}
