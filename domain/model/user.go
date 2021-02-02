package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	Base  `valid:"required"`
	Name  string `json:"name" valid:"required"`
	Email string `json:"name" valid:"required"`
}

func (user *User) isValid() error {
	_, err := govalidator.ValidateStruct(user)

	if err != nil {
		return err
	}

	return nil
}

func NewUser(name string, email string) (*User, error) {
	user := User{
		Name:  name,
		Email: email,
	}

	user.ID = uuid.NewV4().String()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	e := user.isValid()

	if e != nil {
		return nil, e
	}

	return &user, nil
}
