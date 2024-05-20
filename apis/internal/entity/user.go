package entity

import (
	"github.com/becardine/learning-go/apis/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       entity.ID `json:"id"` // value object for id
	Name     string    `json:"name"`
	Password string    `json:"-"`
	Email    string    `json:"email"`
}

func NewUser(name, password, email string) (*User, error) {
	// Hash the password before storing it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:       entity.NewID(),
		Name:     name,
		Password: string(hashedPassword),
		Email:    email,
	}, nil
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
