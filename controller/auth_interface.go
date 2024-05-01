package controller

import "github.com/gin-gonic/gin"

type AuthControllerInterface interface {
	Register(*gin.Context)
	Login(*gin.Context)
}
