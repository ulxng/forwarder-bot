package repository

import (
	"context"
	"sync"
)

type MemoryForwardStorage struct {
	data map[int64]ForwardConfig
	mu   sync.RWMutex
}

func (m *MemoryForwardStorage) Save(ctx context.Context, config ForwardConfig) error {
	m.mu.Lock()
	m.data[config.UserID] = config
	m.mu.Unlock()

	return nil
}

func (m *MemoryForwardStorage) FindChatByUserID(ctx context.Context, userID int64) (int64, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	config, ok := m.data[userID]
	if !ok {
		return 0, nil
	}
	return config.ChatID, nil
}

func NewMemoryForwardStorage() *MemoryForwardStorage {
	return &MemoryForwardStorage{data: make(map[int64]ForwardConfig)}
}
