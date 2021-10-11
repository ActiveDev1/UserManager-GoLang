package db

import (
	"fmt"

	"UserManager/config"
	"UserManager/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase(logger *logger.Logger, conf *config.Config) *gorm.DB {
	logger.GetZapLogger().Infof("Try database connection")
	dsn := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", conf.Database.Username, conf.Database.Password, conf.Database.Host, conf.Database.Dbname)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		logger.GetZapLogger().Errorf("Failure database connection")
	}
	logger.GetZapLogger().Infof("Success database connection, %s:%s", conf.Database.Host, conf.Database.Port)
	return db
}
