package lock

import "sync"

type wrlock struct {
	nreader int
	guard   sync.Mutex
	lock    sync.Mutex
}

func NewWrlock() wrlock {
	return wrlock{
		nreader: 0,
	}
}
func (wr wrlock) writerLock() {
	wr.lock.Lock()
}
func (wr wrlock) writerUnlock() {
	wr.lock.Unlock()
}

func (wr wrlock) readerLock() {
	wr.guard.Lock()
	wr.nreader++
	if wr.nreader == 1 {
		wr.lock.Lock()
	}
	wr.guard.Unlock()
}

func (wr wrlock) readerUnlock() {
	wr.guard.Lock()
	wr.nreader--
	if wr.nreader == 0 {
		wr.lock.Unlock()
	}
	wr.guard.Unlock()
}
