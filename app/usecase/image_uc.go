package usecase

import (
	"context"
	"fmt"
	"picture-go-app/domain/models"
	"picture-go-app/domain/repository"
	"picture-go-app/infrastructure/log"
	"regexp"
	"strconv"

	"github.com/volatiletech/null/v8"
)

type ImageUsecaseIF interface {
	FindAll(ctx context.Context) ([]*ImageContainer, error)
	Insert(ctx context.Context, m *models.Image) error
}

// ImageUsecase is image usecase
type ImageUsecase struct {
	rep repository.ImageRepository
}

type ImageContainer struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	Data           string `json:"data"`
	UpdateDateTime string `json:"updateDateTime"`
	Tagid          string `json:"tagid"`
	Description    string `json:"description"`
}

// NewImageUsecase is factory method
func NewImageUsecase(r repository.ImageRepository) *ImageUsecase {
	return &ImageUsecase{
		rep: r,
	}
}

// FindAll get all data
func (uc ImageUsecase) FindAll(ctx context.Context) ([]*ImageContainer, error) {
	imgs := make([]*ImageContainer, 0, 0)
	mimgs, err := uc.rep.FindAllWithContext(ctx)
	if err != nil {
		return nil, err
	}
	for _, mimg := range mimgs {
		img := &ImageContainer{
			ID:             mimg.ID,
			Name:           mimg.Name,
			Data:           mimg.Data,
			UpdateDateTime: mimg.UpdateDateTime,
			Tagid:          strconv.Itoa(mimg.Tagid.Int),
			Description:    mimg.Description.String,
		}
		imgs = append(imgs, img)
	}
	return imgs, nil
}

// Insert regist data
func (uc ImageUsecase) Insert(ctx context.Context, img *ImageContainer) error {
	var itagid int
	var err error
	rep := regexp.MustCompile("^data:.*;base64,")
	if len(img.Tagid) != 0 {
		itagid, err = strconv.Atoi(img.Tagid)
	}
	if err != nil {
		log.Warn(fmt.Sprintf("failed to atoi at tagid:%s,%v", img.Tagid, err))
	}
	mimg := &models.Image{
		Name:        img.Name,
		Data:        rep.ReplaceAllString(img.Data, ""),
		Tagid:       null.IntFrom(itagid),
		Description: null.StringFrom(img.Description),
	}
	return uc.rep.InsertWithContext(ctx, mimg)
}
