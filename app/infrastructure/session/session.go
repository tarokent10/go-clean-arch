package session

import (
	"encoding/json"
	"fmt"
	"picture-go-app/infrastructure/log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type SessionData struct {
	Userid   string `json:"userid,string"`
	UserName string `json:"username,string"`
}

const (
	skey          string = "skey"
	authURL       string = "/v1/auth/login/"
	SessionCtxKey string = "session"
)

func AuthRequired(c *gin.Context) {
	if c.Request.URL.Path == authURL {
		log.Debug("skip auth")
		c.Next()
		if c.Writer.Status() == 200 {
			// なんかアンチパターン実装な気がするが自作なんでまぁヨシ
			// logined(c)
		}
		return
	}
	s := sessions.Default(c)
	sessionDataJson := s.Get(skey)
	log.Debug(fmt.Sprintf("get session data => %v", sessionDataJson))
	// TODO 必要なら構造体化してコンテキストに詰める
	c.Next()
}

func logined(c *gin.Context) {
	// TODO SetCoockeされない。ミドルウェアでレスポンスに書き込めない可能性がありそう
	s := sessions.Default(c)
	data := &SessionData{ // TODO ユーザ情報が詰められない
		Userid:   "test",
		UserName: "test",
	}
	sdata, err := json.Marshal(data)
	if err != nil {
		log.Err(fmt.Sprintf("failed to merchal session. %v", err))
	}
	println(string(sdata))
	s.Set(skey, string(sdata))
	s.Options(*&sessions.Options{
		HttpOnly: true,
		MaxAge:   3600, //sec
	})
	if err = s.Save(); err != nil {
		log.Err(fmt.Sprintf("failed to save session. %v", err))
	}
	log.Debug("session saved")
}
