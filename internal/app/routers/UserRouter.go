package routers

import (
	v1 "AwesomeBlog/internal/app/api/v1"

	"github.com/gin-gonic/gin"
)

func UserRouter(e *gin.Engine) {
	const base = Version + "user"
	UserAPI := v1.NewUserAPI()

	// 用户注册
	e.POST(base, UserAPI.Create)
}
