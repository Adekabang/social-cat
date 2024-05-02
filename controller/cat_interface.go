package controller

import "github.com/gin-gonic/gin"

type CatControllerInterface interface {
	InsertCat(*gin.Context)
	GetAllCats(*gin.Context)
	// GetOneCat(*gin.Context)
	// UpdateCat(*gin.Context)
	// DeleteCat(*gin.Context)
}
