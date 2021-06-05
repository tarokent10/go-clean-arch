package usecase

import (
	"encoding/base64"
	"picture-go-app/domain/models"
	"picture-go-app/domain/repository"

	"github.com/gin-gonic/gin"
)

type EmployeeUsecaseIF interface {
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

func (uc EmployeeUsecase) FineAll(ctx *gin.Context) ([]*EmployeeContainer, error) {
	ms, err := uc.rep.FindAllWithContext(ctx)
	if err != nil {
		return nil, err
	}
	ret := make([]*EmployeeContainer, 0)
	for _, m := range ms {
		ret = append(ret, uc.toContainer(m))
	}
	return ret, nil
}
func (uc EmployeeUsecase) toContainer(m *models.Employee) *EmployeeContainer {
	entity := &EmployeeContainer{
		Name:           m.Name,
		Picture:        base64.StdEncoding.EncodeToString(m.Picture),
		UpdateDateTime: m.UpdateDateTime.Format("2006/1/2 15:04:05"),
		Department:     m.Department,
		Remarks:        m.Remarks,
	}
	return entity
}
