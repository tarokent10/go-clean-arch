package gateway

import (
	"context"
	"picture-go-app/adapter/gateway/connector"
	"picture-go-app/domain/models"
	"picture-go-app/domain/repository"
	"picture-go-app/infrastructure/database/entities"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type UserRepository struct {
	Con connector.DBConnector
}

func NewUserRepository(c connector.DBConnector) repository.UserRepository {
	return &UserRepository{
		Con: c,
	}
}

func (r UserRepository) FindAllWithContext(ctx context.Context) ([]*models.User, error) {
	es, err := entities.Users().All(ctx, r.Con.GetDB())
	if err != nil {
		// TODO error handle
		return nil, err
	}
	return r.fromEntitys(es), nil
}

func (r UserRepository) FindWithContext(ctx context.Context, userID string) (*models.User, error) {
	e, err := entities.Users(entities.UserWhere.UserID.EQ(userID)).One(ctx, r.Con.GetDB())
	if err != nil {
		// TODO error handle
		return nil, err
	}
	return r.fromEntity(e), nil
}

func (r UserRepository) InsertWithContext(ctx context.Context, m *models.User) error {
	entity := r.toEntity(m)
	return entity.Insert(context.Background(), r.Con.GetDB(), boil.Infer())
}

func (r UserRepository) fromEntity(e *entities.User) *models.User {
	return &models.User{
		ID:       e.ID,
		UserID:   e.UserID,
		Name:     e.Name,
		Password: e.Password,
	}
}

func (r UserRepository) fromEntitys(es []*entities.User) []*models.User {
	ms := []*models.User{}
	for _, e := range es {
		m := &models.User{
			ID:       e.ID,
			UserID:   e.UserID,
			Name:     e.Name,
			Password: e.Password,
		}
		ms = append(ms, m)
	}
	return ms
}

func (r UserRepository) toEntity(m *models.User) *entities.User {
	entity := &entities.User{
		UserID:   m.UserID,
		Name:     m.Name,
		Password: m.Password,
	}
	return entity
}
