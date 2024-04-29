package controller

import "github.com/gin-gonic/gin"

type UserControllerInterface interface {
	InsertUser(*gin.Context)
	GetAllUsers(*gin.Context)
	GetOneUser(*gin.Context)
	UpdateUser(*gin.Context)
	DeleteUser(*gin.Context)
}
