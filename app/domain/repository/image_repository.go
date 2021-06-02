package repository

import (
	"context"
	"picture-go-app/domain/models"
)

type ImageRepository interface {
	FindAllWithContext(ctx context.Context) ([]*models.Image, error)
	InsertWithContext(ctx context.Context, m *models.Image) error
}
