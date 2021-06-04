package usecase

import (
	"encoding/json"
	"fmt"
	"picture-go-app/domain/repository"
	"picture-go-app/infrastructure/log"
	"picture-go-app/infrastructure/session"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type AuthUsecaseIF interface {
	Login(ctx *gin.Context, userID, password string) error
}

// AuthUsecase is image usecase
type AuthUsecase struct {
	rep repository.UserRepository
}
type AuthContainer struct {
	UserID   string `json:"userID"`
	Password string `json:"password"`
}

// NewAuthUsecase is factory method
func NewAuthUsecase(r repository.UserRepository) *AuthUsecase {
	return &AuthUsecase{
		rep: r,
	}
}

func (uc AuthUsecase) Login(ctx *gin.Context, userID, password string) error {
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
