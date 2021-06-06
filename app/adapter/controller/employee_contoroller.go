package controller

import (
	"net/http"
	"picture-go-app/usecase"

	"picture-go-app/infrastructure/log"

	"github.com/gin-gonic/gin"
)

type EmployeeContorollerIF interface {
	FineAll(ctx *gin.Context)
	Regist(ctx *gin.Context)
}

// EmployeeContoroller is a handler of image tag
type EmployeeContoroller struct {
	employeehUC usecase.EmployeeUsecaseIF
}

func NewEmployeeContoroller(euc usecase.EmployeeUsecaseIF) EmployeeContorollerIF {
	return &EmployeeContoroller{
		employeehUC: euc,
	}
}
func (c EmployeeContoroller) Regist(ctx *gin.Context) {
	user := new(usecase.EmployeeContainer)
	if err := ctx.Bind(user); err != nil {
		log.Err(err.Error())
		ctx.String(http.StatusBadRequest, "bad request")
		return
	}
	err := c.employeehUC.Regist(ctx, user)
	if err != nil {
		log.Err(err.Error())
		ctx.String(http.StatusInternalServerError, "internal server error")
		return
	}
	ctx.JSON(http.StatusCreated, "")

}
func (c EmployeeContoroller) FineAll(ctx *gin.Context) {
	records, err := c.employeehUC.FineAll(ctx)
	if err != nil {
		log.Err(err.Error())
		ctx.String(http.StatusInternalServerError, "internal server error")
		return
	}
	ctx.JSON(http.StatusOK, records)

}
