package gateway

import (
	"context"
	"picture-go-app/adapter/gateway/connector"
	"picture-go-app/domain/models"
	"picture-go-app/domain/repository"
	"picture-go-app/infrastructure/database/entities"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type EmployeeRepository struct {
	Con connector.DBConnector
}

func NewEmployeeRepository(c connector.DBConnector) repository.EmployeeRepository {
	return &EmployeeRepository{
		Con: c,
	}
}

func (r EmployeeRepository) FindAllWithContext(ctx context.Context) ([]*models.Employee, error) {
	es, err := entities.Employees().All(ctx, r.Con.GetDB())
	if err != nil {
		return nil, err
	}
	ret := make([]*models.Employee, 0)
	for _, m := range es {
		ret = append(ret, r.fromEntity(m))
	}
	return ret, nil
}

func (r EmployeeRepository) InsertWithContext(ctx context.Context, m *models.Employee) error {
	entity := r.toEntity(m)
	return entity.Insert(context.Background(), r.Con.GetDB(), boil.Infer())
}

func (r EmployeeRepository) fromEntity(e *entities.Employee) *models.Employee {
	return &models.Employee{
		ID:             e.ID,
		Name:           e.Name,
		Picture:        e.Picture,
		UpdateDateTime: e.UpdateDateTime,
		Department:     e.Department.String,
		Remarks:        e.Remarks.String,
	}
}

func (r EmployeeRepository) toEntity(m *models.Employee) *entities.Employee {
	entity := &entities.Employee{
		ID:             m.ID,
		Name:           m.Name,
		Picture:        m.Picture,
		UpdateDateTime: m.UpdateDateTime,
		Department:     null.StringFromPtr(&m.Department),
		Remarks:        null.StringFromPtr(&m.Remarks),
	}
	return entity
}
