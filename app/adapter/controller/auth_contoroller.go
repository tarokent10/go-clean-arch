package controller

import (
	"net/http"
	"picture-go-app/usecase"

	"picture-go-app/infrastructure/log"

	"github.com/gin-gonic/gin"
)

type AuthContorollerIF interface {
	Login(ctx *gin.Context)
}

// UserContoroller is a handler of image tag
type AuthContoroller struct {
	authUC *usecase.AuthUsecase
}

func NewAuthContoroller(auc *usecase.AuthUsecase) AuthContorollerIF {
	return &AuthContoroller{
		authUC: auc,
	}
}

func (c AuthContoroller) Login(ctx *gin.Context) {
	user := new(usecase.AuthContainer)
	if err := ctx.Bind(user); err != nil {
		log.Err(err.Error())
		ctx.String(http.StatusUnauthorized, "ID or Password is uncorrect")
		return
	}
	err := c.authUC.Login(ctx, user.UserID, user.Password)
	if err != nil {
		log.Err(err.Error())
		ctx.String(http.StatusUnauthorized, "ID or Password is uncorrect")
		return
	}
	ctx.JSON(http.StatusOK, "")

}
