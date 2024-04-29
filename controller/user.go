package controller

import (
	"database/sql"

	"github.com/Adekabang/social-cat/model"
	"github.com/Adekabang/social-cat/repository"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	Db *sql.DB
}

func NewUserController(db *sql.DB) UserControllerInterface {
	return &UserController{Db: db}
}

// DeleteUser implements UserControllerInterface
func (m *UserController) DeleteUser(c *gin.Context) {
	DB := m.Db
	var uri model.UserUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewUserRepository(DB)
	delete := repository.DeleteUser(uri.ID)
	if delete {
		c.JSON(200, gin.H{"status": "success", "msg": "delete user successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "msg": "delete user failed"})
		return
	}
}

// GetAllUsers implements UserControllerInterface
func (m *UserController) GetAllUsers(c *gin.Context) {
	DB := m.Db
	repository := repository.NewUserRepository(DB)
	get := repository.GetAllUsers()
	if get != nil {
		c.JSON(200, gin.H{"status": "success", "data": get, "msg": "get users successfully"})
		return
	} else {
		c.JSON(200, gin.H{"status": "success", "data": nil, "msg": "users not found"})
		return
	}
}

// GetOneUser implements UserControllerInterface
func (m *UserController) GetOneUser(c *gin.Context) {
	DB := m.Db
	var uri model.UserUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewUserRepository(DB)
	get := repository.GetOneUser(uri.ID)
	if (get != model.GetUser{}) {
		c.JSON(200, gin.H{"status": "success", "data": get, "msg": "get user successfully"})
		return
	} else {
		c.JSON(200, gin.H{"status": "success", "data": nil, "msg": "user not found"})
		return
	}
}

// InsertUser implements UserControllerInterface
func (m *UserController) InsertUser(c *gin.Context) {
	DB := m.Db
	var post model.PostUser
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewUserRepository(DB)
	insert := repository.InsertUser(post)
	if insert {
		c.JSON(200, gin.H{"status": "success", "msg": "insert user successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "msg": "insert user failed"})
		return
	}
}

// UpdateUser implements UserControllerInterface
func (m *UserController) UpdateUser(c *gin.Context) {
	DB := m.Db
	var post model.PostUser
	var uri model.UserUri
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewUserRepository(DB)
	update := repository.UpdateUser(uri.ID, post)
	if (update != model.GetUser{}) {
		c.JSON(200, gin.H{"status": "success", "data": update, "msg": "update user successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "data": nil, "msg": "update user failed"})
		return
	}
}
