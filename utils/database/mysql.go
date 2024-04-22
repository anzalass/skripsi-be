package database

import (
	"fmt"
	"testskripsi/config"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	var dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.InitConfig().DBUser, config.InitConfig().DBPass, config.InitConfig().DBHost, config.InitConfig().DBPort, config.InitConfig().DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Error("Database : Connect to MySQL Successfully")
	return db
}
