package v1

import (
	"github.com/labstack/echo/v4"
)

func Register(g *echo.Group) {

	v1 := g.Group("v1/")

	v1.POST("login", login)
}
