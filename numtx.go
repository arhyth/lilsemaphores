package lilsem

import (
	"sync"
)

// NuMtx is a counter guarded by a mutex for testing purposes
type NuMtx struct {
	counter int
	mtx     *sync.RWMutex
}

// Inc increments counter safely
func (n *NuMtx) Inc() int {
	return n.add(1)
}

// Dec decrements counter safely
func (n *NuMtx) Dec() int {
	return n.add(-1)
}

// Val returns current counter value
func (n *NuMtx) Val() int {
	var val int

	n.mtx.RLock()
	val = n.counter
	n.mtx.RUnlock()

	return val
}

func (n *NuMtx) add(a int) int {
	var prev int

	n.mtx.Lock()
	prev = n.counter
	n.counter = n.counter + a
	n.mtx.Unlock()

	return prev
}

// NewNuMtx ...
func NewNuMtx() *NuMtx {
	return &NuMtx{
		mtx: &sync.RWMutex{},
	}
}
