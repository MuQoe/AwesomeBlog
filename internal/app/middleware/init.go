package middleware

import (
	"AwesomeBlog/pkg/logger"
	"github.com/gin-gonic/gin"
)

func MiddleInit(g *gin.Engine) {
	g.Use(logger.GinLogger())
	g.Use(logger.GinRecovery(true))
}
