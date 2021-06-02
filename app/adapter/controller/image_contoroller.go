package controller

import (
	"net/http"
	"picture-go-app/usecase"

	"picture-go-app/infrastructure/log"

	"github.com/gin-gonic/gin"
)

type ImageContorollerIF interface {
	GetAll(ctx *gin.Context)
	Create(ctx *gin.Context)
}

// ImageContoroller is a handler of image tag
type ImageContoroller struct {
	imageUC *usecase.ImageUsecase
}

func NewImageContoroller(iuc *usecase.ImageUsecase) ImageContorollerIF {
	return &ImageContoroller{
		imageUC: iuc,
	}
}

func (c ImageContoroller) GetAll(ctx *gin.Context) {
	imgs, err := c.imageUC.FindAll(ctx.Request.Context())
	if err != nil {
		log.Err(err.Error())
		ctx.String(http.StatusInternalServerError, "internal server error")
		return
	}
	ctx.JSON(http.StatusOK, imgs)

}
func (c ImageContoroller) Create(ctx *gin.Context) {
	img := new(usecase.ImageContainer)
	if err := ctx.Bind(img); err != nil {
		log.Err(err.Error())
		ctx.String(http.StatusInternalServerError, "internal server error")
		return
	}
	if err := c.imageUC.Insert(ctx.Request.Context(), img); err != nil {
		log.Err(err.Error())
		ctx.String(http.StatusInternalServerError, "internal server error")
		return
	}
	ctx.String(http.StatusCreated, "")
}
