package dto

import (
	"UserManager/model"

	"github.com/go-playground/validator/v10"
)

const (
	required string = "required"
	max      string = "max"
	min      string = "min"
	length   string = "len"
)

type UserDto struct {
	Username string `validate:"required,min=3,max=10" json:"username"`
	Password string `validate:"required,min=3,max=50" json:"password"`
	Age      uint8  `validate:"required" json:"age"`
}

func NewUserDto() *UserDto {
	return &UserDto{}
}

func (u *UserDto) Validate() map[string]string {
	return validateDto(u)
}

func (u *UserDto) Create() *model.User {
	return model.NewUser(u.Username, u.Password, u.Age)
}

func validateDto(b interface{}) map[string]string {
	result := make(map[string]string)
	err := validator.New().Struct(b)

	if err != nil {
		errors := err.(validator.ValidationErrors)
		if len(errors) != 0 {
			for i := range errors {
				switch errors[i].StructField() {
				case "Username":
					switch errors[i].Tag() {
					case required, min, max:
						result["username"] = "یوزرنیم وارد شده نا معتبر می‌باشد."
					}
				case "Password":
					switch errors[i].Tag() {
					case required, min, max:
						result["passsword"] = "رمز عبور وارد شده نا معتبر می‌باشد."
					}
				case "Age":
					switch errors[i].Tag() {
					case required:
						result["age"] = "سن وارد شده نا معتبر می‌باشد."
					}
				}
			}
		}
		return result
	}

	return nil
}
