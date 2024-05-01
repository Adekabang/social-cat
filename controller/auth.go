package controller

import (
	"database/sql"

	"github.com/Adekabang/social-cat/model"
	"github.com/Adekabang/social-cat/repository"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	Db *sql.DB
}

func NewAuthController(db *sql.DB) AuthControllerInterface {
	return &AuthController{Db: db}
}

func (m *AuthController) Register(c *gin.Context) {
	DB := m.Db
	var user model.RegisterUser

	// if err := c.ShouldBindJSON(&user); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewAuthRepository(DB)
	insert := repository.Register(user)

	if insert == "success" {
		c.JSON(200, gin.H{"status": "success", "msg": "User Registered"})
		return
	} else if insert == "23505" {
		c.JSON(409, gin.H{"status": "failed", "msg": "User Already Registered"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "msg": "server error"})
		return
	}
}
