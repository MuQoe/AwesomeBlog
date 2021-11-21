package configs

import (
	global "AwesomeBlog/globals"
	"time"

	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")    // 指定配置文件名称 (不需要带后缀)
	vp.SetConfigType("yaml")      // 指定配置文件类型
	vp.AddConfigPath("./configs") // 指定查找配置文件的路径 (相对路径)
	err := vp.ReadInConfig()      // 读取配置信息
	if err != nil {
		return nil, err
	}
	return &Setting{vp}, nil
}

func Init() error {
	setting, err := NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}
