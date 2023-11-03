package storage

import "sync"

type MemoryStorage struct {
	data map[string]int64

	mu sync.Mutex
}

func (m *MemoryStorage) Get(key string) (int64, error) {
	return m.data[key], nil
}

func (m *MemoryStorage) Set(key string, value int64) error {
	m.mu.TryLock()
	defer m.mu.Unlock()
	m.data[key] = value
	return nil
}

func NewMemoryStorage() MemoryStorage {
	return MemoryStorage{data: map[string]int64{}}
}
