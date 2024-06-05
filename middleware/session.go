package middleware

import (
	"context"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go-web-new/global"
	"go-web-new/utils/errmsg"

	"time"
)

var code int

type UserInfo struct {
	Username string `json:"username"`

	// 其他用户相关信息
}

// 生成token
func SetToken(c *gin.Context, username string) (string, int) {

	userInfo := UserInfo{
		username,
	}
	session := sessions.Default(c)
	sessionId := uuid.New().String()
	session.Set(sessionId, username)
	sessionTTL := 24 * time.Hour
	global.RedisClient.Set(context.Background(), sessionId, userInfo, sessionTTL)
	c.SetCookie(sessionId, username, int(sessionTTL), "/", "localhost", false, true)

	return sessionId, errmsg.SUCCSE
}

// 验证token
func CheckToken(c *gin.Context, token string) (string, int) {
	data, err := c.Cookie(token)
	if err != nil {
		return "", errmsg.ERROR
	} else {
		ctx := context.Background()
		userinfo, _ := global.RedisClient.Get(ctx, data).Result()
		if err != nil {
			return "", errmsg.ERROR
		}
		return userinfo, errmsg.SUCCSE
	}

}

// jwt管理员中间件
//func JwtToken() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		// 接收参数
//		tokenHeader := c.Request.Header.Get("Authorization")
//		// 如果不存在
//		if tokenHeader == "" {
//			code = errmsg.ERROR_TOKEN_EXIST
//			response.Result(code, "", c)
//			c.Abort()
//			return
//		}
//		// 分割，验证格式
//		checkToken := strings.SplitN(tokenHeader, " ", 2)
//		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
//			code = errmsg.ERROR_TOKEN_TYPE_WRONG
//			response.Result(code, "", c)
//			c.Abort()
//			return
//		}
//
//		// 验证token
//		key, Tcode := CheckToken(checkToken[1])
//		if Tcode == errmsg.ERROR {
//			code = errmsg.ERROR_TOKEN_WRONG
//			response.Result(code, "", c)
//			c.Abort()
//			return
//		}
//		// 验证权限
//		var user model.User
//		global.Db.Where("username=?", key.Username).First(&user)
//		if user.Role != 1 {
//			response.Result(code, "", c)
//			c.Abort()
//			return
//		}
//		// 过期情况
//		if time.Now().Unix() > key.ExpiresAt {
//			code = errmsg.ERROR_TOKEN_RUNTIME
//			response.Result(code, "", c)
//			c.Abort()
//			return
//		}
//		c.Set("username", key.Username)
//		c.Next()
//	}
//}
