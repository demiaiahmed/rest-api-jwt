package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"html"
	"rest-api-jwt/utils/token"
	"strings"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"password"`
}

func (u *User) SaveUser() (*User, error) {
	var err error
	err = DB.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) BeforeSave() error {
	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	//Remove Spaces in Password
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	return nil
}
func LoginCheck(username string, password string) (string, error) {

	var err error

	u := User{}

	err = DB.Model(User{}).Where("username = ?", username).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateToken(u.ID)
	if err != nil {
		return "", nil
	}
	return token, nil
}
func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
func GetUserById(uid uint) (User, error) {
	var u User
	if err := DB.First(&u, uid).Error; err != nil {
		return u, errors.New("User Not Found!")

	}
	u.PrepareGive()
	return u, nil
}

func (u *User) PrepareGive() {
	u.Password = ""
}
