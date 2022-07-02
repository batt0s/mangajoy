package models

import (
	"context"
	"errors"
	"fmt"
	"net/mail"
	"time"

	"github.com/batt0s/mangajoy/database"
	"github.com/uptrace/bun"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	bun.BaseModel `bun:"table:users"`
	ID            int64      `json:"id" bun:"id,pk,autoincrement"`
	CreatedAt     time.Time  `json:"created_at" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt     time.Time  `json:"updated_at" bun:",nullzero,notnull,default:current_timestamp"`
	LastLogin     time.Time  `json:"last_login"`
	IsAdmin       bool       `json:"is_admin" form:"is_admin"`
	IsStaff       bool       `json:"is_staff" form:"is_staff"`
	Username      string     `form:"username" json:"username" bun:",unique"`
	Email         string     `form:"email" json:"email" bun:",unique"`
	Password      string     `form:"password" json:"password"`
	Uploads       []*Chapter `bun:"rel:has-many,join:uploader_id"`
}

// This should be done before using the model for the first time
func (u User) Init() {
	database.DB.RegisterModel((*User)(nil))
}

func (u User) String() string {
	return fmt.Sprintf("User<%d, %s, %s, %v, %v, %v, %v - %v>",
		u.ID, u.Username, u.Email, u.IsAdmin, u.IsStaff, u.LastLogin, u.CreatedAt, u.UpdatedAt)
}

func (u *User) Create() error {
	ctx := context.Background()
	if !u.IsValid() {
		return errors.New("user not valid")
	}
	pwdHash, err := createHash(u.Password)
	if err != nil {
		return err
	}
	u.Password = pwdHash
	if _, err := database.DB.NewInsert().Model(u).Exec(ctx); err != nil {
		return err
	}
	return nil
}

// DO NOT USE THIS TO UPDATE THE PASSWORD. USE SetPassword() INSTEAD
func (u *User) Update() error {
	ctx := context.Background()
	if !u.IsValid() {
		return errors.New("user not valid")
	}
	_, err := database.DB.NewUpdate().Model(u).WherePK().Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) SetPassword(newPass string) error {
	ctx := context.Background()
	newHash, err := createHash(newPass)
	if err != nil {
		return err
	}
	u.Password = newHash
	_, err = database.DB.NewUpdate().Model(u).Column("password").Where("id = ?", u.ID).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) IsValid() bool {
	if len(u.Password) < 6 {
		return false
	}
	if len(u.Username) < 4 {
		return false
	}
	if !isValidEmail(u.Email) {
		return false
	}
	return true
}

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func Authenticate(email, password string) (*User, error) {
	user := &User{}
	var err error
	ctx := context.Background()
	err = database.DB.NewSelect().Model(user).Where("email = ?", email).Scan(ctx)
	if err != nil {
		return nil, err
	}
	err = checkPassword(user.Password, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func createHash(pwd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func checkPassword(hash, givenPwd string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(givenPwd))
	return err
}
