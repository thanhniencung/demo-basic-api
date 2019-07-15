package repository

import "basic-api/model"

type UserRepository interface {
	Insert(user model.User) error
}
