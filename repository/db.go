package repository

import (
	"context"
	"ulxng/forwarder-bot/db"
)

type forwardConfigRepository struct {
	q *db.Queries
}

func (r *forwardConfigRepository) Save(ctx context.Context, config ForwardConfig) error {
	return r.q.SaveConfig(ctx, db.SaveConfigParams{
		UserID: config.UserID,
		ChatID: config.ChatID,
	})
}

func (r *forwardConfigRepository) FindChatByUserID(ctx context.Context, userID int64) (int64, error) {
	id, err := r.q.GetChatByUserID(ctx, userID)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func NewForwardConfigRepository(q *db.Queries) ForwardConfigRepository {
	return &forwardConfigRepository{q: q}
}
