package v1

import (
	"github.com/gin-gonic/gin"
	"go-web-new/model"
	"go-web-new/model/response"
	"go-web-new/utils/errmsg"
	"go-web-new/utils/validator"
	"net/http"
	"strconv"
)

type UserApi struct {
}

// 添加用户
func (u *UserApi) AddUser(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	msg, code := validator.Validate(&data)
	if code != errmsg.SUCCSE {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"data": data,
			"msg":  msg,
		})
		c.Abort()
		return
	}
	code = model.CheckUser(data.Username)
	if code == errmsg.SUCCSE {
		data.CreateUser()
	}
	if code == errmsg.ERROR_USERNAME_USED {
		code = errmsg.ERROR_USERNAME_USED
	}
	response.Result(code, data, c)
}

// 查询用户
func (u *UserApi) GetUserInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := model.User{}
	code := data.GetUser(id)
	response.Result(code, data, c)
}
