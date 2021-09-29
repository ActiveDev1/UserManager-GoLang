package migration

import (
	"UserManager/container"
	"UserManager/model"
)

// CreateDatabase creates the tables used in this application.
func CreateDatabase(container container.Container) {
	if container.GetConfig().Database.Migration {
		db := container.GetDatabase()

		db.AutoMigrate(&model.User{})
	}
}
