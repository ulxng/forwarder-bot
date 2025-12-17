package repository

import "context"

type ForwardConfig struct {
	UserID int64
	ChatID int64
}

type ForwardConfigRepository interface {
	Save(ctx context.Context, config ForwardConfig) error
	FindChatByUserID(ctx context.Context, userID int64) (int64, error)
}
