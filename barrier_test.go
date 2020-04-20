package lilsem_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/arhyth/lilsem"
)

func TestBarrier(t *testing.T) {
	count := rand.Intn(1001)

	b := lilsem.NewBarrier(0)
	var counter int

	for i := 0; i < count; i++ {
		b.Add(1)
		go func(b *lilsem.Barrier) {
			b.Wait()
			counter++
			b.Signal()
		}(b)
	}

	// this is a workaround to keep the "main" goroutine alive
	// long enough for other goroutines to finish work
	// since the pattern does not indicate a second "rendezvous"
	time.Sleep(time.Second)

	if count != counter {
		t.Fatalf("barrier failed, expected: %v, got: %v", count, counter)
	}
}
