package app

import (
	"github.com/gin-gonic/gin"
)

var GinEngine *gin.Engine


const (
	ADMINISTRATOR = iota
	USER
)
// Init 初始化
func Init() {
	// 1. 初始化配置项

	// 2. 初始化日志

	// 3. 初始化Mysql连接

	// 4. 初始化Redis连接

	// 关闭debug模式

	//初始化gin框架

	//初始化中间件

	//设置随机数种子

	// 5. 注册路由 初始化路由
}
