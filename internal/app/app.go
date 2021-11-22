package app

import (
	"AwesomeBlog/configs"
	global "AwesomeBlog/globals"
	"AwesomeBlog/internal/app/middleware"
	"AwesomeBlog/internal/app/routers"
	"AwesomeBlog/pkg/database"
	"AwesomeBlog/pkg/logger"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var GinEngine *gin.Engine

// Init 初始化
func Init() {
	// 1. 初始化配置项
	if err := configs.Init(); err != nil {
		fmt.Printf("init configs failed, err:%v\n", err)
		return
	}
	// 2. 初始化日志
	if err := logger.Init(); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	zap.S().Debug("Logger init Success...")

	// 3. 初始化Mysql连接
	if err := database.Init(); err != nil {
		fmt.Printf("init database failed, err:%v\n", err)
		return
	}
	// 4. 初始化Redis连接

	// 关闭debug模式
	gin.SetMode(global.ServerSetting.RunMode)
	//初始化gin框架
	GinEngine = gin.New()
	//初始化中间件
	middleware.MiddleInit(GinEngine)
	//设置随机数种子
	rand.Seed(time.Now().Unix())

	// 5. 注册路由 初始化路由
	routers.Init(GinEngine)
	//启动项目,这里我们使用http2.0版本

	address := ":" + global.ServerSetting.HttpPort
	fmt.Println(address)
	s := &http.Server{
		Addr:           address,
		Handler:        GinEngine,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	zap.S().Infof("Listening and serving HTTP on %s", address)
	err := s.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		zap.S().Errorf("%v\n", err)
		return
	}

	/*
		if err := GinEngine.Run(":" + viper.GetString("app.port")); err != nil {
			fmt.Println(err)
			zap.L().Error(err.Error())
		}
	*/
	// 6. 启动服务 (优雅关机)
}
