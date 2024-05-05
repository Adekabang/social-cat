package repository

import (
	"github.com/Adekabang/social-cat/model"
)

type CatRepositoryInterface interface {
	InsertCat(model.PostCat) model.CatResponseMessage
	GetAllCats(model.GetCat) []model.Cat
	UpdateCat(string, model.PostCat) int
	DeleteCat(string, string) bool
}
