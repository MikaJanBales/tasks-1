package counter

import "sync/atomic"

// счетчик с использованием атомарных операций
type AtomicCounter struct {
	counter int64
}

func NewAtomicStruct() *AtomicCounter {
	return &AtomicCounter{
		counter: 0,
	}
}

func (a *AtomicCounter) Add() {
	atomic.AddInt64(&a.counter, 1)
}

func (a *AtomicCounter) Get() int64 {
	return atomic.LoadInt64(&a.counter)
}
