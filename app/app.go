package app

import (
	"database/sql"

	auth "github.com/Adekabang/social-cat/controller"
	cat "github.com/Adekabang/social-cat/controller"
	user "github.com/Adekabang/social-cat/controller"
	"github.com/Adekabang/social-cat/db"
	"github.com/Adekabang/social-cat/middleware"
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
	public.POST("/login", controllerAuth.Login)

	protected := r.Group("/admin")
	protected.Use(middleware.JwtAuthMiddleware())
	protected.GET("/user", controller.GetAllUsers)

	controllerCat := cat.NewCatController(a.DB)
	publicCat := r.Group("/v1/cat")
	publicCat.Use(middleware.JwtAuthMiddleware())
	publicCat.POST("/", controllerCat.InsertCat)
	publicCat.GET("/", controllerCat.GetAllCats)
	publicCat.GET("/:id", controllerCat.GetAllCats)
	publicCat.PUT("/:id", controllerCat.UpdateCat)
	publicCat.DELETE("/:id", controllerCat.DeleteCat)

	a.Router = r
}

func (a *App) Run() {
	a.Router.Run(":8080")
}
