package database

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	var dsn = fmt.Sprintf("root:@tcp(localhost:3306)/grasinet?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Error("Database : Connect to MySQL Successfully")
	return db
}
