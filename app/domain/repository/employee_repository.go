package repository

import (
	"context"
	"picture-go-app/domain/models"
)

type EmployeeRepository interface {
	FindAllWithContext(ctx context.Context) ([]*models.Employee, error)
	InsertWithContext(ctx context.Context, m *models.Employee) error
}
