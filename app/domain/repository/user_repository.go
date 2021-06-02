package repository

import (
	"context"
	"picture-go-app/domain/models"
)

type UserRepository interface {
	FindWithContext(ctx context.Context, userID string) (*models.User, error)
	FindAllWithContext(ctx context.Context) ([]*models.User, error)
	InsertWithContext(ctx context.Context, m *models.User) error
}
