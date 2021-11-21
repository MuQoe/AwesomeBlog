package routers

import (
	v1 "AwesomeBlog/internal/app/api/v1"
	"github.com/gin-gonic/gin"
)

func IndexRouter(e *gin.Engine) {
	const base = Version + "index"
	Index := v1.NewIndex()
	e.GET(base, Index.Get)
}
