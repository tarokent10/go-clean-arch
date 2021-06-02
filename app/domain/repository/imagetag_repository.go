package repository

import (
	"context"
	"picture-go-app/domain/models"
)

type ImagetagRepository interface {
	FindAllWithContext(ctx context.Context) []*models.ImageTag
}
