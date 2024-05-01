package app

import (
	"database/sql"

	auth "github.com/Adekabang/social-cat/controller"
	user "github.com/Adekabang/social-cat/controller"
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
	controller := user.NewUserController(a.DB)
	r.POST("/users", controller.InsertUser)
	r.GET("/users", controller.GetAllUsers)
	r.GET("/users/:id", controller.GetOneUser)
	r.PUT("/users/:id", controller.UpdateUser)
	r.DELETE("/users/:id", controller.DeleteUser)

	controllerAuth := auth.NewAuthController(a.DB)
	public := r.Group("/v1/user")
	public.POST("/register", controllerAuth.Register)

	a.Router = r
}

func (a *App) Run() {
	a.Router.Run(":8080")
}
