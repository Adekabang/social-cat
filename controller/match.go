package controller

import (
	"database/sql"

	"github.com/Adekabang/social-cat/model"
	"github.com/Adekabang/social-cat/repository"
	"github.com/Adekabang/social-cat/utils"
	"github.com/gin-gonic/gin"
)

type MatchController struct {
	Db *sql.DB
}

func NewMatchController(db *sql.DB) MatchControllerInterface {
	return &MatchController{Db: db}
}

// RequestMatch implements MatchControllerInterface
func (m *MatchController) RequestMatch(c *gin.Context) {
	DB := m.Db
	var requestMatch model.RequestMatch
	if err := c.ShouldBind(&requestMatch); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}

	userId, err := utils.GetUserId(c.GetHeader(("Authorization")))
	if err != nil {
		c.JSON(500, gin.H{"message": "failed", "msg": "failed to get user id"})
		return
	}

	requestMatch.IssuedBy = userId

	repository := repository.NewMatchRepository(DB)
	request := repository.RequestMatch(requestMatch)
	if request {
		c.JSON(201, gin.H{"message": "success", "data": gin.H{"id": request, "createdAt": request}})
		return
	} else {
		c.JSON(500, gin.H{"message": "failed", "msg": "request match failed"})
		return
	}
}

// GetMatchRequest implements MatchControllerInterface
func (m *MatchController) GetMatchRequest(c *gin.Context) {

	userId, err := utils.GetUserId(c.GetHeader(("Authorization")))
	if err != nil {
		return
	}

	DB := m.Db
	repository := repository.NewMatchRepository(DB)
	get := repository.GetMatchRequest(userId)
	if get != nil {
		c.JSON(200, gin.H{"status": "success", "data": get, "msg": "get match successfully"})
		return
	} else {
		c.JSON(200, gin.H{"status": "success", "data": make([]string, 0), "msg": "cats not found"})
		return
	}
}

// DeleteRequestMatch implements MatchControllerInterface
func (m *MatchController) DeleteRequestMatch(c *gin.Context) {
	DB := m.Db
	var uri model.MatchUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}

	userId, err := utils.GetUserId(c.GetHeader(("Authorization")))
	if err != nil {
		c.JSON(500, gin.H{"message": "failed", "msg": "failed to get user id"})
		return
	}

	repository := repository.NewMatchRepository(DB)
	delete := repository.DeleteRequestMatch(uri.ID, userId)
	if delete {
		c.JSON(200, gin.H{"status": "success", "msg": "delete request successfully"})
		return
	} else {
		c.JSON(404, gin.H{"data": make([]string, 0)})
		return
	}
}

func (m *MatchController) ApproveMatch(c *gin.Context) {
	DB := m.Db
	var uri model.PostApproveReject
	if err := c.ShouldBind(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewMatchRepository(DB)
	approve := repository.ApproveMatch(uri.MatchId)
	if approve {
		c.JSON(200, gin.H{"status": "success", "msg": "approve request successfully"})
		return
	} else {
		c.JSON(404, gin.H{"data": make([]string, 0)})
		return
	}
}

func (m *MatchController) RejectMatch(c *gin.Context) {
	DB := m.Db
	var uri model.PostApproveReject
	if err := c.ShouldBind(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewMatchRepository(DB)
	approve := repository.RejectMatch(uri.MatchId)
	if approve {
		c.JSON(200, gin.H{"status": "success", "msg": "reject request successfully"})
		return
	} else {
		c.JSON(404, gin.H{"data": make([]string, 0)})
		return
	}
}