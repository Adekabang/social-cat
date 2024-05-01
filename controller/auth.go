package controller

import (
	"database/sql"

	"github.com/Adekabang/social-cat/model"
	"github.com/Adekabang/social-cat/repository"
	"github.com/Adekabang/social-cat/utils"
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

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}

	errEmail := utils.ValidateEmail(user.Email)
	errName := utils.ValidateName(user.Name)
	errPassword := utils.ValidatePassword(user.Password)
	if !errEmail || !errName || !errPassword {
		c.JSON(400, gin.H{"status": "failed", "msg": "email:not null, can't be duplicate email, should be in email format, name:not null, minLength 5, maxLength 50, name can be duplicate with others, password:not null, minLength 5, maxLength 15"})
		return
	}

	repository := repository.NewAuthRepository(DB)
	insert := repository.Register(user)

	if insert == "success" {
		c.JSON(201, gin.H{"status": "success", "msg": "User Registered"})
		return
	} else if insert == "23505" {
		c.JSON(409, gin.H{"status": "failed", "msg": "User Already Registered"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "msg": "server error"})
		return
	}
}
func (m *AuthController) Login(c *gin.Context) {
	DB := m.Db
	var input model.LoginUser

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}

	errEmail := utils.ValidateEmail(input.Email)
	errPassword := utils.ValidatePassword(input.Password)
	if !errEmail || !errPassword {
		c.JSON(400, gin.H{"status": "failed", "msg": "email:not null, should be in email format, password:not null, minLength 5, maxLength 15"})
		return
	}

	repository := repository.NewAuthRepository(DB)
	check := repository.Login(input)

	if check.Status == "success" {
		c.JSON(200, gin.H{"status": "success", "msg": "User Logon", "accessToken": check.Msg})
		return
	} else if check.Msg == "user not found" {
		c.JSON(404, gin.H{"status": "failed", "msg": "username not found."})
		return
	} else if check.Msg == "wrong password" {
		c.JSON(400, gin.H{"status": "failed", "msg": "wrong password."})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "msg": "server error"})
		return
	}
}
