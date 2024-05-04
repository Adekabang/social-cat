package controller

import "github.com/gin-gonic/gin"

type MatchControllerInterface interface {
	RequestMatch(*gin.Context)
	GetMatchRequest(*gin.Context)
	ApproveMatch(*gin.Context)
	RejectMatch(*gin.Context)
	DeleteRequestMatch(*gin.Context)
}
