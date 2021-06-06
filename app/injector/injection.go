package injector

import (
	"fmt"
	"picture-go-app/adapter/controller"
	"picture-go-app/adapter/gateway"
	"picture-go-app/adapter/gateway/connector"
	"picture-go-app/infrastructure/log"
	"picture-go-app/usecase"
	"strings"

	"github.com/pkg/errors"
	"go.uber.org/dig"
)

// InitDIContainer init container
func InitDIContainer(conn connector.DBConnector) (*dig.Container, error) {
	// TODO database層の抽象化をしてdig.provideでdi. sql.openは別のところでやる。
	log.Debug("init di container start")
	defer log.Debug("init di container end")

	errors := make([]error, 0, 0)
	// dependency injection
	c := dig.New()
	// repository
	errors = appendIfExists(c.Provide(func() connector.DBConnector {
		return conn
	}), errors)
	errors = appendIfExists(c.Provide(gateway.NewUserRepository), errors)
	errors = appendIfExists(c.Provide(gateway.NewEmployeeRepository), errors)
	// usecase
	errors = appendIfExists(c.Provide(usecase.NewAuthUsecase), errors)
	errors = appendIfExists(c.Provide(usecase.NewEmployeeUsecase), errors)
	// contoroller
	errors = appendIfExists(c.Provide(controller.NewAuthContoroller), errors)
	errors = appendIfExists(c.Provide(controller.NewEmployeeContoroller), errors)

	return c, mergeError(errors)
}
func appendIfExists(e error, errors []error) []error {
	if e != nil {
		errors = append(errors, e)
	}
	return errors
}
func mergeError(es []error) error {
	if len(es) == 0 {
		return nil
	}
	str := "dependency injection faild: "
	for _, e := range es {
		str += fmt.Sprintf("%s, ", e.Error())
	}
	str = strings.TrimRight(str, ",")
	return errors.New(str)
}
