package lilsem

import (
	"sync"
)

// NuMtx is a counter guarded by a mutex for testing purposes
type NuMtx struct {
	counter int
	mtx     *sync.Mutex
}

// Inc increments counter safely
func (n *NuMtx) Inc() int {
	var prev int

	n.mtx.Lock()
	prev = n.counter
	n.counter = n.counter + 1
	n.mtx.Unlock()

	return prev
}

// NewNuMtx ...
func NewNuMtx() *NuMtx {
	return &NuMtx{
		mtx: &sync.Mutex{},
	}
}
