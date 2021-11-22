package routers

import (
	"github.com/gin-gonic/gin"
)

const Version = "/api/v1/"
const AdminVersion = "/api/v1/admin/"

func Init(g *gin.Engine) {
	// 注册各个板块的路由
	IndexRouter(g)
	UserRouter(g)
	// 注册Admin 面板路由
	// TODO: dev admin panel
}
