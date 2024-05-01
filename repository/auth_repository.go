package repository

import "github.com/Adekabang/social-cat/model"

type AuthRepositoryInterface interface {
	Register(model.RegisterUser) model.ResponseMessage
	Login(model.LoginUser) model.ResponseMessage
}
