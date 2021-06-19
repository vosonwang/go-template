package v1

import (
	"github.com/labstack/echo/v4"
)

func Register(g *echo.Group) {

	v2 := g.Group("/v1")

	v2.POST("/login", login)
}
