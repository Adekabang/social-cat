package app

import (
	"database/sql"

	"github.com/Adekabang/social-cat/controller"
	"github.com/Adekabang/social-cat/db"
	"github.com/gin-gonic/gin"
)

type App struct {
	DB     *sql.DB
	Router *gin.Engine
}

func (a *App) CreateConnection() {
	db := db.Connectdb()
	a.DB = db
}

func (a *App) Routes() {
	r := gin.Default()
	controller := controller.NewUserController(a.DB)
	r.POST("/users", controller.InsertUser)
	r.GET("/users", controller.GetAllUsers)
	r.GET("/users/:id", controller.GetOneUser)
	r.PUT("/users/:id", controller.UpdateUser)
	r.DELETE("/users/:id", controller.DeleteUser)
	a.Router = r
}

func (a *App) Run() {
	a.Router.Run(":8080")
}
