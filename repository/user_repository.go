package repository

import "github.com/Adekabang/social-cat/model"

type UserRepositoryInterface interface {
	InsertUser(model.PostUser) bool
	GetAllUsers() []model.GetUser
	GetOneUser(string) model.GetUser
	UpdateUser(string, model.PostUser) model.GetUser
	DeleteUser(string) bool
}
