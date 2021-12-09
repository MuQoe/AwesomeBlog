package routers

import (
	v1 "AwesomeBlog/internal/app/api/v1"
	"AwesomeBlog/internal/app/middleware"

	"github.com/gin-gonic/gin"
)

func IndexRouter(e *gin.Engine) {
	const base = Version + "index"
	Index := v1.NewIndex()
	authRoute := e.Group("/", middleware.Auth)
	authRoute.GET(base, Index.Get)
}
