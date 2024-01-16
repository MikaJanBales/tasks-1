package counter

import "sync"

// счетчик с использование мьютексов
type MutexCounter struct {
	counter int64
	mutex   *sync.RWMutex
}

func NewMutexStruct() *MutexCounter {
	return &MutexCounter{
		counter: 0,
		mutex:   &sync.RWMutex{},
	}
}

func (m *MutexCounter) Add() {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.counter++
}

func (m *MutexCounter) Get() int64 {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.counter
}
