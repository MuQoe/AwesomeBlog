package middleware

import (
	global "AwesomeBlog/globals"
	"AwesomeBlog/internal/app/model"
	"AwesomeBlog/pkg/Response"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
)

type Claims struct {
	ID    int64 `json:"ID"`
	Exp   int64 `json:"exp"`
	Admin bool  `json:"admin"`
	jwt.StandardClaims
}

// 判断用户是否有权限访问
func Auth(c *gin.Context) {
	tokenString, _ := c.Cookie("token")
	if tokenString == "" {
		Response.GlobalResponse.ResponseUnauthorized(c, "Auth Needed")
		c.Abort()
		return
	}
	claims := &Claims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.AppSetting.TokenSecret), nil
	})
	if err != nil {
		zap.S().Warnf("Auth Middleware ParseWithClaims Error %v", err)
		Response.GlobalResponse.ResponseUnauthorized(c, "Internal error")
		c.Abort()
		return
	}
	if claims.Exp <= time.Now().Unix() {

	}
	// TODO: 权限判断
	user, err := model.User().FindByID(claims.ID)
	if err != nil {
		Response.GlobalResponse.ResponseUnauthorized(c, "UserID not found 1")
		c.Abort()
		return
	}
	if user.IsEmpty() {
		Response.GlobalResponse.ResponseUnauthorized(c, "UserID not found 2")
		c.Abort()
		return
	}
	return
}
