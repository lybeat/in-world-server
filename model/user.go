package model

import (
	"in-world-server/pkg/setting"

	"github.com/jinzhu/gorm"
)

type User struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Age      int    `json:"age"`
	Sex      int    `json:"sex"`
	Avatar   string `json:"avatar"`
}

func (u *User) Register() error {
	user := User{
		Username: u.Username,
		Password: u.Password,
		Age:      u.Age,
		Sex:      u.Sex,
		Avatar:   u.Avatar,
	}
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func Login(username string, password string) (*User, error) {
	var user User
	err := db.Where("username = ? AND password = ?", username, password).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &user, nil
}

func ExistUserByID(id int) (bool, error) {
	var user User
	err := db.Select("id").Where("id = ?", id).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if user.ID > 0 {
		return true, nil
	}

	return false, nil
}

func GetUserTotal() (int, error) {
	var count int
	if err := db.Model(&User{}).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func GetUser(id int) (*User, error) {
	var user User
	err := db.Where("id = ?", id).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &user, nil
}

func GetUsers(pageNum int) ([]*User, error) {
	var users []*User
	pageSize := setting.AppSetting.PageSize
	err := db.Offset(pageNum).Limit(pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return users, nil
}
