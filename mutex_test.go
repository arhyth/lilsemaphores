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
	var counter int
	var m = lilsem.NewMutex()
	for i := 0; i < count; i++ {
		w.Add(1)
		go func(w *sync.WaitGroup, m *lilsem.Mutex) {
			m.Wait()
			counter++
			m.Signal()

			w.Done()
		}(w, m)
	}
	w.Wait()

	if counter != count {
		// if a data race exists, counter may be less than count
		t.Fatalf("rendezvous, expected: %v, got: %v", count, counter)
	}
}
