package main

import (
	"AwesomeBlog/internal/app"
	"github.com/kardianos/service"
	"go.uber.org/zap"
	"log"
	"time"
)

// the main entry point
// Go Web 个人博客开发

// Author: David.R

// Go Web 个人博客开发

func SerStart() {
	// 注册插件
	// plugin.ScanAllPlugins()
	// 启动定时任务
	// cron.Init()
	// 启动项目
	app.Init()
}

// 当前应用程序的结构体
type program struct{}

// Start 服务启动
func (p *program) Start(s service.Service) error {
	log.Println("开始服务")
	go p.run()
	return nil
}

// Stop 停止程序
func (p *program) Stop(s service.Service) error {
	log.Println("停止服务")
	return nil
}

// 启动程序
func (p *program) run() {
	//API初始化
	SerStart()
}

// @title 博客系统
// @version 1.0
// @description 练手项目
func main() {
	//服务的配置信息
	cfg := &service.Config{
		Name:        "博客系统",
		DisplayName: "阿呆博客",
		Description: "前后端分离模式 练手用博客系统项目",
	}
	// 程序的接口
	prg := &program{}
	// 构建服务对象
	s, err := service.New(prg, cfg)
	if err != nil {
		log.Fatal(err)
	}
	err = s.Run()
	if err != nil {
		time.Sleep(time.Second * 5)
		zap.S().Errorf("%v\n", err)
	}
}
