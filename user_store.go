package example

import (
	"github.com/jmoiron/sqlx"
)

type UserStore interface {
	GetByUsername(username string) (*User, error)
	Create(user *User) error
}

type userStore struct {
	db *sqlx.DB
}

func NewUserStore(db *sqlx.DB) *userStore {
	return &userStore{
		db: db,
	}
}

func (s userStore) GetByUsername(username string) (*User, error) {
	return nil, nil
}

func (s userStore) Create(user *User) error {
	res, err := s.db.NamedExec(
		"INSERT INTO users (name, age) VALUES (:name, :age)",
		user,
	)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = id
	return nil
}
