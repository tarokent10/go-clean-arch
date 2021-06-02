package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"picture-go-app/domain/models"
	"picture-go-app/domain/repository"
	"picture-go-app/infrastructure/log"
	"picture-go-app/infrastructure/session"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type UserUsecaseIF interface {
	Login(ctx *gin.Context, userID, password string) error
	Insert(ctx context.Context, m *models.User) error
}

// UserUsecase is image usecase
type UserUsecase struct {
	rep repository.UserRepository
}
type UserContainer struct {
	UserID   string `json:"userID"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

// NewUserUsecase is factory method
func NewUserUsecase(r repository.UserRepository) *UserUsecase {
	return &UserUsecase{
		rep: r,
	}
}

func (uc UserUsecase) Login(ctx *gin.Context, userID, password string) error {
	// TODO ユーザユースケースにログインがあるのはおかしいのでどっかに移す
	m, err := uc.rep.FindWithContext(ctx, userID)
	if err != nil {
		return err
	}
	if password != m.Password {
		// password is not correct
		return fmt.Errorf("uncorrect auth info")
	}
	logined(ctx)
	return nil
}

// Insert regist data
func (uc UserUsecase) Insert(ctx context.Context, u *UserContainer) error {
	m := &models.User{
		UserID:   u.UserID,
		Name:     u.Name,
		Password: u.Password,
	}
	return uc.rep.InsertWithContext(ctx, m)
}

func logined(c *gin.Context) {
	// TODO SetCoockeされない。ミドルウェアでレスポンスに書き込めない可能性がありそう
	s := sessions.Default(c)
	data := &session.SessionData{ // TODO ユーザ情報が詰められない
		Userid:   "test",
		UserName: "test",
	}
	sdata, err := json.Marshal(data)
	if err != nil {
		log.Err(fmt.Sprintf("failed to merchal session. %v", err))
	}
	println(string(sdata))
	s.Set("skey", string(sdata))
	s.Options(*&sessions.Options{
		HttpOnly: true,
		MaxAge:   3600, //sec
	})
	if err = s.Save(); err != nil {
		log.Err(fmt.Sprintf("failed to save session. %v", err))
	}
	log.Debug("session saved")
}
