package repository

import "ulxng/forwarder-bot/db"

type forwardConfigRepository struct {
	q *db.Queries
}

func (m *forwardConfigRepository) Save(config ForwardConfig) error {

	return nil
}

func (m *forwardConfigRepository) Update(config ForwardConfig) error {
	return nil
}

func (m *forwardConfigRepository) FindByUser(userID int64) (*ForwardConfig, error) {
	return nil, nil
}

func NewForwardConfigRepository(q *db.Queries) ForwardConfigRepository {
	return &forwardConfigRepository{q: q}
}
