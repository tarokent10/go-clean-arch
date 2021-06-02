package gateway

import (
	"context"
	"encoding/base64"
	"picture-go-app/adapter/gateway/connector"
	"picture-go-app/domain/models"
	"picture-go-app/domain/repository"
	"picture-go-app/infrastructure/database/entities"
	"time"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type ImageRepository struct {
	Con connector.DBConnector
}

func NewImageRepository(c connector.DBConnector) repository.ImageRepository {
	return &ImageRepository{
		Con: c,
	}
}

func (r ImageRepository) FindAllWithContext(ctx context.Context) ([]*models.Image, error) {
	es, err := entities.ImageItems().All(ctx, r.Con.GetDB())
	if err != nil {
		// TODO error handle
		return nil, err
	}
	return r.fromEntity(es), nil
}

func (r ImageRepository) InsertWithContext(ctx context.Context, m *models.Image) error {
	entity, err := r.toEntity(m)
	if err != nil {
		return err
	}
	return entity.Insert(context.Background(), r.Con.GetDB(), boil.Infer())
}

func (r ImageRepository) fromEntity(es []*entities.ImageItem) []*models.Image {
	ms := []*models.Image{}
	for _, e := range es {
		m := &models.Image{
			ID:             e.ID,
			Name:           e.Name,
			Data:           base64.StdEncoding.EncodeToString(e.Data),
			UpdateDateTime: e.UpdateDateTime.Format("2006/01/02 15:04:05"),
			Tagid:          e.TagID,
			Description:    e.Description,
		}
		ms = append(ms, m)
	}
	return ms
}

func (r ImageRepository) toEntity(m *models.Image) (*entities.ImageItem, error) {
	data, e := base64.StdEncoding.DecodeString(m.Data)
	if e != nil {
		return nil, e
	}

	entity := &entities.ImageItem{
		Name:           m.Name,
		Data:           data,
		UpdateDateTime: time.Now(),
		TagID:          m.Tagid,
		Description:    m.Description,
	}
	return entity, nil
}
