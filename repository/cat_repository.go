package repository

import (
	"github.com/Adekabang/social-cat/model"
)

type CatRepositoryInterface interface {
	InsertCat(model.PostCat) bool
	GetAllCats(model.GetCat) []model.Cat
	// GetOneCat(string) model.GetCat
	// UpdateCat(string, model.PostCat) model.GetCat
	// DeleteCat(string) bool
}
