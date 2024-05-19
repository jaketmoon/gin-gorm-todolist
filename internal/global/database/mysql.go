package database

import (
	"fmt"
	"gin/configs"
	"gin/internal/global/log"
	"gin/internal/model"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	DB *gorm.DB
)

func Connect() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%s&loc=%s",
		configs.DbSettings.Root,
		configs.DbSettings.Password,
		configs.DbSettings.Host,
		configs.DbSettings.Port,
		configs.DbSettings.Dbname,
		configs.DbSettings.Charset,
		configs.DbSettings.ParseTime,
		configs.DbSettings.Loc,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) //gorm.Open方法尝试使用构建的DSN（数据源名称）和MySQL驱动打开一个新的数据库连接。
	if err != nil {
		log.SugarLogger.Error(err)
		return nil
	}
	fmt.Println("连接数据库成功")

	sqlDB, err3 := db.DB()
	if err != nil {
		log.SugarLogger.Error(err3)
		return nil
	}
	sqlDB.SetMaxIdleConns(20)           // 设置最大空闲连接数
	sqlDB.SetMaxOpenConns(200)          // 设置最大打开连接数
	sqlDB.SetConnMaxLifetime(time.Hour) // 设置了连接可复用的最大时间
	err1 := db.AutoMigrate(&model.User{})
	err2 := db.AutoMigrate(&model.Todo{})
	if err1 != nil {
		log.SugarLogger.Error(err1)
		return nil
	}
	fmt.Println("user数据库迁移成功")
	if err2 != nil {
		log.SugarLogger.Error(err2)
		return nil
	}

	fmt.Println("todo数据库迁移成功")
	return db
}
func Init() {
	DB = Connect()
}
