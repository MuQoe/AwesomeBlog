package global

import (
	"time"
)

type ServerSettingStructure struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSettingStructure struct {
	DefaultPageSize int
	MaxPageSize     int
	LogMaxSize      int
	LogMaxAge       int
	LogMaxBackups   int
	LogSavePath     string
	LogFileName     string
	LogFileExt      string
	LogLevel        string
}
type DatabaseSettingStructure struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    string
	MaxIdleConns int
	MaxOpenConns int
}

var (
	ServerSetting   *ServerSettingStructure
	AppSetting      *AppSettingStructure
	DatabaseSetting *DatabaseSettingStructure
)
