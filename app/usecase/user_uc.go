package usecase

import (
	"encoding/json"
	"fmt"
	"picture-go-app/domain/models"
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
	writeSessionStore(ctx, m)
	return nil
}
func (uc AuthUsecase) Logout(ctx *gin.Context) error {
	s := sessions.Default(ctx)
	s.Clear()
	s.Options(sessions.Options{
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1, // delete session from store
	})
	if err := s.Save(); err != nil {
		log.Err(fmt.Sprintf("failed to delete session. %v", err))
	}
	log.Debug("session deleted")
	return nil
}
func writeSessionStore(c *gin.Context, m *models.User) {
	s := sessions.Default(c)
	data := &session.SessionData{
		Userid:   m.UserID,
		UserName: m.Name,
	}
	sdata, err := json.Marshal(data)
	if err != nil {
		log.Err(fmt.Sprintf("failed to merchal session. %v", err))
	}
	log.Debug(string(sdata))
	s.Set(session.DataKey, string(sdata))
	s.Options(sessions.Options{
		Path:     "/",
		HttpOnly: true,
		MaxAge:   2000, //sec
	})
	if err = s.Save(); err != nil {
		log.Err(fmt.Sprintf("failed to save session. %v", err))
	}
	log.Debug("session saved")
}
