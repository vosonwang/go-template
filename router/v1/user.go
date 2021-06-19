package v1

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"my-project-name/db/mongodb"
	"my-project-name/model"
	"my-project-name/router/response"
)

// login godoc
// @Summary 登录
// @Description 通过用户名、密码登录
// @Accept  json
// @Produce  json
// @Param username body string true "用户名"
// @Param password body string true "密码"
// @Success 200 {object} response.Style{data=model.User}
// @Failure default {object} response.Style
// @Router /login [post]
func login(c echo.Context) error {
	var user model.User

	if err := c.Bind(&user); err != nil {
		return response.BadRequest(c, err)
	}

	user, err := mongodb.FindUserByName(user.Username)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return response.Fail(c, "user not found")
		}
		return response.Fail(c, err)
	}

	return response.OK(c, user)
}
