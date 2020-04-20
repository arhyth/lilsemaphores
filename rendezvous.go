package lilsem

import (
	"sync"
)

// rendezvous in principle works similar to a sync.WaitGroup
type rendezvous struct {
	*sync.Mutex

	count int
}

func (r *rendezvous) Done() {
	r.Add(-1)
}

func (r *rendezvous) Add(p int) {
	r.Lock()
	if r.count == 0 && p < 0 {
		panic("cannot decrement a zero Rendezvous")
	}
	r.count = r.count + p
	r.Unlock()
}

func (r *rendezvous) Wait() {
	// This busywaiting is probably very inefficient
	// but I have no idea how else to implement.
	for {
		if r.count != 0 {
			continue
		}
		break
	}

	return
}

func Rendezvous(count ...int) *rendezvous {
	var c int
	if count != nil {
		c = count[0]
	}

	return &rendezvous{&sync.Mutex{}, c}
}
