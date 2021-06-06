package usecase

import (
	"encoding/base64"
	"picture-go-app/domain/models"
	"picture-go-app/domain/repository"
	"time"

	"github.com/gin-gonic/gin"
)

type EmployeeUsecaseIF interface {
	Regist(ctx *gin.Context, container *EmployeeContainer) error
	FineAll(ctx *gin.Context) ([]*EmployeeContainer, error)
}

// AuthUsecase is image usecase
type EmployeeUsecase struct {
	rep repository.EmployeeRepository
}
type EmployeeContainer struct {
	Name           string `json:"name"`
	Picture        string `json:"picture"`
	UpdateDateTime string `json:"updateDateTime"`
	Department     string `json:"department"`
	Remarks        string `json:"remarks"`
}

// NewEmployeeUsecase is factory method
func NewEmployeeUsecase(r repository.EmployeeRepository) EmployeeUsecaseIF {
	return &EmployeeUsecase{
		rep: r,
	}
}
func (uc EmployeeUsecase) Regist(ctx *gin.Context, container *EmployeeContainer) error {
	model, err := uc.toModel(container)
	if err != nil {
		return err
	}
	return uc.rep.InsertWithContext(ctx, model)
}
func (uc EmployeeUsecase) FineAll(ctx *gin.Context) ([]*EmployeeContainer, error) {
	ms, err := uc.rep.FindAllWithContext(ctx)
	if err != nil {
		return nil, err
	}
	ret := make([]*EmployeeContainer, 0)
	for _, m := range ms {
		ret = append(ret, uc.fromModel(m))
	}
	return ret, nil
}
func (uc EmployeeUsecase) fromModel(m *models.Employee) *EmployeeContainer {
	c := &EmployeeContainer{
		Name:           m.Name,
		Picture:        base64.StdEncoding.EncodeToString(m.Picture),
		UpdateDateTime: m.UpdateDateTime.Format("2006/1/2 15:04:05"),
		Department:     m.Department,
		Remarks:        m.Remarks,
	}
	return c
}
func (uc EmployeeUsecase) toModel(c *EmployeeContainer) (*models.Employee, error) {
	pictureByte, err := base64.StdEncoding.DecodeString(c.Picture)
	if err != nil {
		return nil, err
	}
	model := &models.Employee{
		Name:           c.Name,
		Picture:        pictureByte,
		UpdateDateTime: time.Now(),
		Department:     c.Department,
		Remarks:        c.Remarks,
	}
	return model, nil
}
