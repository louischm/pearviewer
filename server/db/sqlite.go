package db

import (
	"github.com/louischm/pkg/logger"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var log = logger.NewLog()

func OpenDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("./data/pearviewer.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&User{})
	return db, nil
}

func Migrate(db *gorm.DB) {

}

func AddUser(db *gorm.DB, user *User) {
	log.Info("Creating new user: %s", user.String())
	user.Password = encryptPassword(user.Password)
	db.Create(user)
}

func encryptPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func GetUserByUserName(db *gorm.DB, userName string) User {
	var user User

	db.First(&user, "Username = ?", userName)
	return user
}
