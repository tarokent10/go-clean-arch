package session

import (
	"fmt"
	"net/http"
	"picture-go-app/infrastructure/log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type SessionData struct {
	Userid   string `json:"userid,string"`
	UserName string `json:"username,string"`
}

const (
	SsnKey  string = "session-key"
	DataKey string = "session-data-key"
	authURL string = "/v1/auth/login/"
)

func AuthRequired(c *gin.Context) {
	if c.Request.URL.Path == authURL {
		log.Debug("skip auth")
		c.Next()
		return
	}
	s := sessions.Default(c)
	sessionDataJson := s.Get(SsnKey)
	log.Debug(fmt.Sprintf("get session data => %v", sessionDataJson))
	// TODO 必要なら構造体化してコンテキストに詰める
	if sessionDataJson != nil {
		c.AbortWithStatus(http.StatusUnauthorized) // not login
	}
	c.Next()
}
