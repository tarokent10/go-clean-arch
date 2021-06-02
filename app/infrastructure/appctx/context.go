package appctx

import "github.com/labstack/echo"

type AppContext struct {
	echo.Context
	requestID string
}
