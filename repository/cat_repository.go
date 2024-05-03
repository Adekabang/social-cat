package repository

import (
	"github.com/Adekabang/social-cat/model"
)

type CatRepositoryInterface interface {
	InsertCat(model.PostCat) bool
	GetAllCats(model.GetCat) []model.Cat
	UpdateCat(string, model.PostCat) bool
	DeleteCat(string) bool
}
