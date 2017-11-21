package example

import (
	"gopkg.in/go-playground/validator.v9"
)

type Server interface {
	CreateUser(user *User) error
}

type server struct {
	userStore UserStore
}

func NewServer(userStore UserStore) *server {
	return &server{
		userStore: userStore,
	}
}

func (s server) CreateUser(user *User) error {
	validate := validator.New()

	err := validate.Struct(user)
	if err != nil {
		return err
	}
	return nil
}
