package container

import (
	"UserManager/config"
	"UserManager/logger"

	"gorm.io/gorm"
)

type Container interface {
	GetDatabase() *gorm.DB
	GetLogger() *logger.Logger
	GetConfig() *config.Config
	GetEnv() string
}

type container struct {
	db     *gorm.DB
	logger *logger.Logger
	config *config.Config
	env    string
}

func NewContainer(db *gorm.DB, logger *logger.Logger, config *config.Config, env string) Container {
	return &container{db: db, logger: logger, config: config, env: env}
}

func (c *container) GetDatabase() *gorm.DB {
	return c.db
}

func (c *container) GetLogger() *logger.Logger {
	return c.logger
}

func (c *container) GetConfig() *config.Config {
	return c.config
}

func (c *container) GetEnv() string {
	return c.env
}
