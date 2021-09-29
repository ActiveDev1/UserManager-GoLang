package db

import (
	"fmt"

	"UserManager/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase(conf *config.Config) *gorm.DB {
	fmt.Println("Try database connection")
	dsn := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", conf.Database.Username, conf.Database.Password, conf.Database.Host, conf.Database.Dbname)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic("Failure database connection")
	}
	fmt.Printf("Success database connection, %s:%s", conf.Database.Host, conf.Database.Port)
	return db
}
