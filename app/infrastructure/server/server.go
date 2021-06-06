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
	if err != nil {
		log.Fatal(fmt.Sprintf("failed to new redis session. error => %v", err))
	}
	r.Use(sessions.Sessions("mysession", store))
	r.Use(session.AuthRequired)
	if err := c.Invoke(func(
		authController controller.AuthContorollerIF,
		employeeController controller.EmployeeContorollerIF,
	) {
		r.POST("/v1/auth/login/", authController.Login)
		r.POST("/v1/auth/logout/", authController.Logout)
		r.GET("/v1/employees/", employeeController.FineAll)
		r.POST("/v1/employee/", employeeController.Regist)
	}); err != nil {
		log.Fatal(fmt.Sprintf("error at dig invoke-> %s", err.Error()))
	}
	log.Info("now linetennig...")
	if err := r.Run(":3000"); err != nil {
		log.Fatal(fmt.Sprintf("error at gin run. error-> %s", err.Error()))
	}
}
