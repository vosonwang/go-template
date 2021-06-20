// @title RICN Smart IoT Platform API
// @version 1.0
// @description 卫川物联网平台API
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host iot.ricnsmart.com
// @BasePath /api/v1

package router

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"my-project-name/config"
	_ "my-project-name/docs"
	Middleware "my-project-name/router/middleware"
	"my-project-name/router/pprof"
	"my-project-name/router/v1"
	"net/http"
	"regexp"
	"strings"
)

// jwtCustomClaims are custom claims extending default ones.
type jwtCustomClaims struct {
	UserID   primitive.ObjectID `json:"user_id"`
	Username string             `json:"username"`
	RoleID   primitive.ObjectID `json:"role_id"`
	jwt.StandardClaims
}

func Init() *echo.Echo {
	e := echo.New()

	// 开发环境下直接panic，让bug查找更加清晰
	if !config.IsDev() {
		e.Use(middleware.Recover())
	}

	e.Use(middleware.Logger())

	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Skipper: func(c echo.Context) bool {
			// 解决使用了Gzip后打开Swagger页面空白
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

	// prefix必须尾部加/，否则如果找不到路由，会提示jwt缺失，而不是not found
	api := e.Group("/api/")

	api.Use(Middleware.JWTWithConfig(Middleware.JWTConfig{
		Skipper: func(c echo.Context) bool {
			// 登录、登出、获取验证码、验证码登录不需要token
			return regexp.MustCompile("login|logout|captcha").Match([]byte(c.Path()))
		},
		Claims:      &jwtCustomClaims{},
		TokenLookup: "cookie:token",
		SigningKey:  []byte(config.ProjectName()),
	}))

	v1.Register(api)

	address := fmt.Sprintf(":%v", config.HttpPort())

	go func() {
		if err := e.Start(address); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	return e
}
