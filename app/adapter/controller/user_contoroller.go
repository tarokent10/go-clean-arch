package controller

import (
	"net/http"
	"picture-go-app/usecase"

	"picture-go-app/infrastructure/log"

	"github.com/gin-gonic/gin"
)

type UserContorollerIF interface {
	Login(ctx *gin.Context)
	Create(ctx *gin.Context)
}

// UserContoroller is a handler of image tag
type UserContoroller struct {
	userUC *usecase.UserUsecase
}

func NewUserContoroller(iuc *usecase.UserUsecase) UserContorollerIF {
	return &UserContoroller{
		userUC: iuc,
	}
}

func (c UserContoroller) Login(ctx *gin.Context) {
	user := new(usecase.UserContainer)
	if err := ctx.Bind(user); err != nil {
		log.Err(err.Error())
		ctx.String(http.StatusUnauthorized, "ID or Password is uncorrect")
		return
	}
	err := c.userUC.Login(ctx, user.UserID, user.Password)
	if err != nil {
		log.Err(err.Error())
		ctx.String(http.StatusUnauthorized, "ID or Password is uncorrect")
		return
	}
	ctx.JSON(http.StatusOK, "")

}
func (c UserContoroller) Create(ctx *gin.Context) {
	user := new(usecase.UserContainer)
	if err := ctx.Bind(user); err != nil {
		log.Err(err.Error())
		ctx.String(http.StatusInternalServerError, "internal server error")
		return
	}
	if err := c.userUC.Insert(ctx.Request.Context(), user); err != nil {
		log.Err(err.Error())
		ctx.String(http.StatusInternalServerError, "internal server error")
		return
	}
	ctx.String(http.StatusCreated, "")
}
