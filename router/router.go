// @title RICN Smart IoT Platform API
// @version 1.0
// @description
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host iot.ricnsmart.com
// @BasePath /api/v1
// @query.collection.format multi

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationurl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationurl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

// @x-extension-openapi {"example": "value on a json format"}
package router

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger"
	"my-project-name/config"
	_ "my-project-name/docs"
	"my-project-name/router/pprof"
	"my-project-name/router/v1"
	"net/http"
	"strings"
)

func Init() *echo.Echo {
	e := echo.New()

	// 开发环境下直接panic，让bug查找更加清晰
	if !config.IsDev() {
		e.Use(middleware.Recover())
	}

	e.Use(middleware.Logger())

	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Skipper: func(c echo.Context) bool {
			if strings.Contains(c.Request().URL.Path, "swagger") {
				return true
			}
			return false
		},
	}))

	pprof.Register(e)

	e.GET("/swagger/*", echoSwagger.EchoWrapHandler(
		func(c *echoSwagger.Config) {
			// 必须按如下设置，否则index.html页面上的json文件打开后是空白页
			c.URL = "./swagger/doc.json"
		}),
	)

	api := e.Group("/api")

	v1.Register(api)

	address := fmt.Sprintf(":%v", config.HttpPort())

	go func() {
		if err := e.Start(address); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	return e
}
