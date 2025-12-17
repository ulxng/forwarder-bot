package repository

import "sync"

type MemoryForwardStorage struct {
	data map[int64]ForwardConfig
	mu   sync.RWMutex
}

func (m *MemoryForwardStorage) Save(config ForwardConfig) error {
	m.mu.Lock()
	m.data[config.UserID] = config
	m.mu.Unlock()

	return nil
}

func (m *MemoryForwardStorage) Update(config ForwardConfig) error {
	m.mu.Lock()
	m.data[config.UserID] = config
	m.mu.Unlock()

	return nil
}

func (m *MemoryForwardStorage) FindByUser(userID int64) (*ForwardConfig, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	config, ok := m.data[userID]
	if !ok {
		return nil, nil
	}
	return &config, nil
}

func NewMemoryForwardStorage() *MemoryForwardStorage {
	return &MemoryForwardStorage{data: make(map[int64]ForwardConfig)}
}
