package storage

import "sync"

type MemoryStorage struct {
	data map[string]map[string]int64

	mu sync.Mutex
}

func (m *MemoryStorage) Get(network string, key string) (int64, error) {
	return m.data[network][key], nil
}

func (m *MemoryStorage) Set(network string, key string, value int64) error {
	m.mu.TryLock()
	defer m.mu.Unlock()
	if m.data[network] == nil {
		m.data[network] = make(map[string]int64)
	}
	m.data[network][key] = value
	return nil
}

func NewMemoryStorage() MemoryStorage {
	return MemoryStorage{data: make(map[string]map[string]int64)}
}
