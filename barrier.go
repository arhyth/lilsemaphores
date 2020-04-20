package lilsem

import (
	"context"
	"time"

	"golang.org/x/sync/semaphore"
)

// Barrier is a rendezvous followed by a critical region.
// For simplicity, only a single semaphore that acts as
// a mutex is used to protect all race prone operations
// of the data structure
type Barrier struct {
	blocks int
	mode   barrierMode
	sem    *semaphore.Weighted
}

type barrierMode int

const (
	rendezvousBM barrierMode = iota
	mutexBM
)

func (b *Barrier) Add(n int) {
	b.add(n)
}

func (b *Barrier) add(n int) {
	b.sem.Acquire(context.Background(), mutexWeight)
	b.blocks = b.blocks + n
	b.sem.Release(mutexWeight)
}

func (b *Barrier) Wait() {
	b.add(-1)

	for {
		if b.blocks == 0 {
			b.mode = mutexBM
			break
		}

		// this perf hit is a tradeoff to minimize busywaiting
		time.Sleep(200 * time.Nanosecond)
	}

	b.sem.Acquire(context.Background(), mutexWeight)
}

func (b *Barrier) Signal() {
	b.sem.Release(mutexWeight)
}

func NewBarrier(b int) *Barrier {
	return &Barrier{
		blocks: b,
		mode:   rendezvousBM,
		sem:    semaphore.NewWeighted(mutexWeight),
	}
}
