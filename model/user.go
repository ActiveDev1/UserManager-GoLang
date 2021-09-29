package model

import (
	"github.com/godruoyi/go-snowflake"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uint64 `gorm:"primarykey;autoIncrement:false"`
	Username  string `gorm:"uniqueIndex;size:36;not null"`
	Password  string
	Age       uint8
	CreatedAt int32
	UpdatedAt int32
}

func (User) TableName() string {
	return "users"
}

func NewUser(Username string, Password string, Age uint8) *User {
	return &User{ID: snowflake.ID(), Username: Username, Password: Password, Age: Age}
}

func (u *User) HashPassword() {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password = string(hashed)
}

func (u *User) Create(db *gorm.DB) (*User, error) {
	if error := db.Create(u).Error; error != nil {
		return nil, error
	}
	return u, nil
}

func (user *User) FindAll(db *gorm.DB) (*[]User, error) {
	var users []User
	result := db.Select("id", "username", "age", "created_at", "updated_at").Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return &users, nil
}

func (u *User) Update(db *gorm.DB, id uint64) (*User, error) {
	if error := db.Model(&u).Where("id = ?", id).Updates(u).Error; error != nil {
		return nil, error
	}
	return u, nil
}

func (u *User) FindByID(db *gorm.DB, id uint64) (*User, error) {
	if error := db.Where("id = ?", id).Find(&u).Error; error != nil {
		return nil, error
	}
	return u, nil
}

func (u *User) DeleteByID(db *gorm.DB, id uint64) (*User, error) {
	if error := db.Delete(&User{}, id).Error; error != nil {
		return nil, error
	}
	return u, nil
}
