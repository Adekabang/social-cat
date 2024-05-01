package repository

import "github.com/Adekabang/social-cat/model"

type AuthRepositoryInterface interface {
	Register(model.RegisterUser) string
	Login(model.LoginUser) model.ResponseMessage
}
