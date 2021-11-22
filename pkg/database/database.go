package database

import (
	global "AwesomeBlog/globals"
	"AwesomeBlog/internal/app/model"
	"fmt"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func NewDBEngine(databaseSetting *global.DatabaseSettingStructure) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local",
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset)

	config := &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	config.Logger = logger.Default.LogMode(logger.Info)
	config.NamingStrategy = schema.NamingStrategy{
		TablePrefix:   "blog_",
		SingularTable: true,
	}
	db, err := gorm.Open(mysql.Open(dsn), config)

	if err != nil {
		zap.S().Warn("Database Connection Failed!")
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		zap.S().Warn("Database.DB() Connection Failed!")
		panic(err)
	}
	sqlDB.SetMaxIdleConns(databaseSetting.MaxIdleConns)
	sqlDB.SetMaxOpenConns(databaseSetting.MaxOpenConns)

	return db, nil
}

func initDataBase() {
	needMigrate := []interface{}{
		&model.User{},
	}
	for _, v := range needMigrate {
		err := global.DBEngine.AutoMigrate(v)
		if err != nil {
			zap.S().Warn(err)
			return
		}
	}
}
func initRedisConnection() {

}

func Init() (err error) {
	global.DBEngine, err = NewDBEngine(global.DatabaseSetting)
	// 如果有需要初始化数据库
	initDataBase()
	// 初始化Redis链接
	// initRedisConnection()
	return err
}
