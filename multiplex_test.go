package lilsem_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/arhyth/lilsem"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func TestMultiplex(t *testing.T) {
	// in this test we simulate a random number of goroutines contending access
	// to a mutex protected counter that records the maximum number of goroutines
	// inside the holding region (the one before the critical region where the actual
	// write happens) and fails if it exceeds it
	limit := int64(5)
	count := rand.Intn(1000)
	holding := lilsem.NewMultiplex(limit)

	counter := lilsem.NewNuMtx()
	for i := 0; i < count; i++ {
		holding.Wait()
		go func(m *lilsem.Multiplex, c *lilsem.NuMtx) {
			u := int64(c.Inc())
			if u > limit-1 {
				c.Dec()
				holding.Signal()
				t.Fatalf("multiplex exceeded: %v", u)
			}
			c.Dec()
			holding.Signal()
		}(holding, counter)
	}

	// t.Fatalf("multiplex failure, counter expected: %v, got: %v", count, counter.Val())
}
