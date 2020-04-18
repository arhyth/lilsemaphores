package lilsem_test

import (
	"math/rand"
	"testing"

	"github.com/arhyth/lilsem"
)

func TestRendezvous(t *testing.T) {
	count := rand.Intn(101)
	r := lilsem.Rendezvous(count)

	var collected int
	for i := 1; i <= count; i++ {
		go func() {
			collected++ // data race
			// adding a Println here or even inside the critical section of rendezvous
			// does not help at all, as Println seems an asynchronous write to stdout
			r.Done()
			return
		}()
	}
	r.Wait()
	if collected != count {
		// this is proven to be an unreliable test due to the tagged data race
		t.Fatalf("rendezvous, expected: %v, got: %v", count, collected)
	}
}
