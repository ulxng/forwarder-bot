package repository

type ForwardConfig struct {
	UserID int64
	ChatID int64
}

type ForwardConfigRepository interface {
	Save(config ForwardConfig) error
	Update(config ForwardConfig) error
	FindByUser(userID int64) (*ForwardConfig, error)
}
