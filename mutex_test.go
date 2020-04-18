package lilsem_test

import (
	"math/rand"
	"sync"
	"testing"

	"github.com/arhyth/lilsem"
)

func TestMutex(t *testing.T) {
	count := rand.Intn(1000)

	// wait for all goroutines to finish
	w := &sync.WaitGroup{}
	n := &numtx{
		counter: 0,
		mtx:     lilsem.NewMutex(),
	}
	for i := 0; i < count; i++ {
		w.Add(1)
		go func(w *sync.WaitGroup, n *numtx) {
			n.inc()
			w.Done()
		}(w, n)
	}
	w.Wait()

	if n.counter != count {
		// if some data race happens, counter will be less than count
		t.Fatalf("rendezvous, expected: %v, got: %v", count, n.counter)
	}
}

// numtx is a counter guarded by a mutex
type numtx struct {
	counter int
	mtx     *lilsem.Mutex
}

// inc increments numtx counter
func (n *numtx) inc() int {
	var prev int
	n.mtx.Wait()
	prev = n.counter
	n.counter = n.counter + 1
	n.mtx.Signal()
	return prev
}
