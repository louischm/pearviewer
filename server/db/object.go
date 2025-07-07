package db

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string `gorm:"size:255"`
}

func (u *User) String() string {
	return u.Username + "\r" + u.Password
}
