package server

import (
	"fmt"
	"picture-go-app/adapter/controller"
	"picture-go-app/infrastructure/log"
	"picture-go-app/infrastructure/session"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func Run(c *dig.Container) {
	log.Debug("start run")
	defer log.Debug("end run")
	// set routing table for resource handler
	r := gin.Default()
	store, err := redis.NewStore(10, "tcp", "redis:6379", "", []byte("secret"))
	// store := cookie.NewStore([]byte("secret"))
	if err != nil {
		log.Fatal(fmt.Sprintf("failed to new redis session. error => %v", err))
	}
	r.Use(sessions.Sessions("mysession", store))
	r.Use(session.AuthRequired)
	if err := c.Invoke(func(
		imageController controller.ImageContorollerIF,
		userController controller.UserContorollerIF,
	) {
		r.GET("/v1/images/", imageController.GetAll)
		r.POST("/v1/users/", userController.Create)
		r.POST("/v1/auth/login/", userController.Login)
		r.POST("/v1/images/", imageController.Create)
		r.GET("/v1/test/", func(ctx *gin.Context) {
			s := sessions.Default(ctx)
			fmt.Printf("get: %v¥n", s.Get("key"))
			s.Set("key", "testtesttest")
			if err := s.Save(); err != nil {
				println(err.Error())
			}
		})
	}); err != nil {
		log.Fatal(fmt.Sprintf("error at dig invoke-> %s", err.Error()))
	}
	log.Info("now linetennig...")
	if err := r.Run(":3000"); err != nil {
		log.Fatal(fmt.Sprintf("error at gin run. error-> %s", err.Error()))
	}
}
