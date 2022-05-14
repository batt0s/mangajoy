package models

import (
	"context"
	"fmt"
	"mangajoy/database"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	LastLogin time.Time
	IsAdmin   bool   `json:"is_admin" form:"is_admin"`
	IsStaff   bool   `json:"is_staff" form:"is_staff"`
	Username  string `form:"username" json:"username"`
	Email     string `form:"email" json:"email"`
	Password  string `form:"password" json:"password"`
}

func (u User) String() string {
	return fmt.Sprintf("User<%d, %s, %s, %v, %v, %v, %v - %v>",
		u.ID, u.Username, u.Email, u.IsAdmin, u.IsStaff, u.LastLogin, u.CreatedAt, u.UpdatedAt)
}

func (u *User) Save() error {
	ctx := context.Background()
	pwdHash, err := CreateHash(u.Password)
	if err != nil {
		return err
	}
	u.Password = pwdHash
	if _, err := database.DB.NewInsert().Model(u).Exec(ctx); err != nil {
		return err
	}
	return nil
}

func Authenticate(email, password string) (*User, error) {
	user := &User{}
	var err error
	ctx := context.Background()
	err = database.DB.NewSelect().Model(user).Where("email = ?", email).Scan(ctx)
	if err != nil {
		return nil, err
	}
	err = CheckPassword(user.Password, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func CreateHash(pwd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func CheckPassword(hash, givenPwd string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(givenPwd))
	return err
}
