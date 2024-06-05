package routes

import (
	"github.com/gin-gonic/gin"
	"go-web-new/api"
)

func V1RouterInit(r *gin.RouterGroup) {
	router := r.Group("api/v1")
	var ApiV1Router = api.ApiGroupApp.V1Group
	{

		router.GET("/user/:id/", ApiV1Router.GetUserInfo)

		router.POST("user/", ApiV1Router.AddUser)
		router.POST("login/", ApiV1Router.Login)

	}
}
