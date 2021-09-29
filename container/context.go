package container

import (
	"UserManager/config"

	"gorm.io/gorm"
)

type Container interface {
	GetDatabase() *gorm.DB
	GetConfig() *config.Config
	GetEnv() string
}

type container struct {
	db     *gorm.DB
	config *config.Config
	env    string
}

func NewContainer(db *gorm.DB, config *config.Config, env string) Container {
	return &container{db: db, config: config, env: env}
}

func (c *container) GetDatabase() *gorm.DB {
	return c.db
}

func (c *container) GetConfig() *config.Config {
	return c.config
}

func (c *container) GetEnv() string {
	return c.env
}
