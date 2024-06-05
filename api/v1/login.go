package v1

import (
	"github.com/gin-gonic/gin"
	"go-web-new/middleware"
	"go-web-new/model"
	"go-web-new/model/response"
	"go-web-new/model/schemas"
	"go-web-new/utils/errmsg"
)

type LoginApi struct {
}

func (l *LoginApi) Login(c *gin.Context) {
	var (
		formData schemas.Login
		token    string
		code     int
		user     model.User
	)
	_ = c.ShouldBindJSON(&formData)

	// 验证用户名密码
	code = user.CheckLogin(formData)
	if code == errmsg.SUCCSE {
		// 成功则签发token
		token, code = middleware.SetToken(c, user.Username)

	}
	response.Result(code, token, c)
	return

}
