package gateway

import (
	"context"
	"picture-go-app/adapter/gateway/connector"
	"picture-go-app/domain/models"
	"picture-go-app/domain/repository"
	"picture-go-app/infrastructure/database/entities"
)

type ImagetagRepository struct {
	Con connector.DBConnector
}

func NewImagetagRepository(c connector.DBConnector) repository.ImagetagRepository {
	return &ImagetagRepository{
		Con: c,
	}
}

func (r ImagetagRepository) FindAllWithContext(ctx context.Context) []*models.ImageTag {
	es, err := entities.ImageTags().All(ctx, r.Con.GetDB())
	if err != nil {
		// TODO error handle
		panic(err)
	}
	return r.fromEntity(es)
}

func (r ImagetagRepository) fromEntity(es []*entities.ImageTag) []*models.ImageTag {
	ms := []*models.ImageTag{}
	for _, e := range es {
		m := &models.ImageTag{ID: e.ID, Name: e.Name}
		ms = append(ms, m)
	}
	return ms
}
