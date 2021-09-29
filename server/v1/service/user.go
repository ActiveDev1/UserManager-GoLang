package service

import (
	"UserManager/container"
	"UserManager/model"
	"UserManager/model/dto"
)

type UserService struct {
	container container.Container
}

func NewUserService(container container.Container) *UserService {
	return &UserService{container: container}
}

func (u *UserService) CreateNewUser(dto *dto.UserDto) (map[string]interface{}, map[string]string) {
	errors := dto.Validate()
	var err error

	if errors != nil {
		return nil, errors
	}

	db := u.container.GetDatabase()
	var result *model.User

	user := dto.Create()
	user.HashPassword()

	if result, err = user.Create(db); err != nil {
		return nil, map[string]string{"error": "Failed to the create user"}
	}
	newUser := map[string]interface{}{"ID": result.ID, "Username": result.Username, "Age": result.Age, "CreatedAt": result.CreatedAt, "UpdatedAt": result.UpdatedAt}
	return newUser, nil
}

func (u *UserService) FindAllUsers() (*[]model.User, error) {
	db := u.container.GetDatabase()
	user := model.User{}
	result, err := user.FindAll(db)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *UserService) UpdateUser(dto *dto.UserDto, id uint64) (map[string]interface{}, map[string]string) {
	errors := dto.Validate()
	if errors != nil {
		return nil, errors
	}

	var err error
	db := u.container.GetDatabase()
	var user *model.User

	if user, err = user.FindByID(db, id); err != nil {
		return nil, map[string]string{"error": "Failed to retrieve user"}
	}
	user.HashPassword()
	user.Username = dto.Username
	user.Age = dto.Age

	var result *model.User

	if result, err = user.Update(db, id); err != nil {
		return nil, map[string]string{"error": "Failed to update user"}
	}
	userData := map[string]interface{}{"ID": result.ID, "Username": result.Username, "Age": result.Age, "CreatedAt": result.CreatedAt, "UpdatedAt": result.UpdatedAt}

	return userData, errors
}

func (u *UserService) DeleteUser(id uint64) map[string]string {
	var err error
	db := u.container.GetDatabase()
	var user *model.User

	if _, err = user.DeleteByID(db, id); err != nil {
		return map[string]string{"error": "Failed to delete user"}
	}

	return nil
}
