package storage

import "sync"

type ForwardConfiguration struct {
	UserID int64
	ChatID int64
}

type ForwardStorage interface {
	Save(config ForwardConfiguration) error
	FindByUser(userID int64) (*ForwardConfiguration, error)
}

type MemoryForwardStorage struct {
	data map[int64]ForwardConfiguration
	mu   sync.RWMutex
}

func (m *MemoryForwardStorage) Save(config ForwardConfiguration) error {
	m.mu.Lock()
	m.data[config.UserID] = config
	m.mu.Unlock()

	return nil
}

func (m *MemoryForwardStorage) FindByUser(userID int64) (*ForwardConfiguration, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	config, ok := m.data[userID]
	if !ok {
		return nil, nil
	}
	return &config, nil
}

func NewMemoryForwardStorage() *MemoryForwardStorage {
	return &MemoryForwardStorage{data: make(map[int64]ForwardConfiguration)}
}
