package config

import (
	"fmt"
	"go-gql/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var mysqlDB *gorm.DB

type Mysql struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	UserName string `toml:"username"`
	Password string `toml:"password"`
	Database string `toml:"database"`
}

func initMysql() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		Config.Mysql.UserName, Config.Mysql.Password, Config.Mysql.Host,
		Config.Mysql.Port, Config.Mysql.Database,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic("init mysql error")
	}
	_ = db.AutoMigrate(&model.User{})
	log.Println("init mysql success")
	mysqlDB = db
}

func GetDB() *gorm.DB {
	return mysqlDB
}
