package main

import (

	// "std/middlewares"
	"std/pkg/handler"
	initPkg "std/pkg/init"
	"std/pkg/repository"
	"std/pkg/service"

	// "std/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// err := godotenv.Load(".env")

	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	r := gin.Default()

	// r.Static("/pictures", "./pictures")
	// r.LoadHTMLGlob("templates/*.html")
	// r.Use(static.Serve("/pictures", static.LocalFile("/pictures", true)))

	auth := r.Group("/auth")

	user := r.Group("/user")

	// auth.POST("/login", routes.Login)

	// protected := r.Group("/user")

	// protected.Use(middlewares.JwtAuthMiddleware())

	dbConn, _ := initPkg.InitDb()
	userRepo := repository.NewUserRepository(dbConn)
	repository := repository.Repository{
		UserRepository: userRepo,
	}
	userService := service.NewUserService(&repository)
	handler := handler.NewHandler(userService)
	handler.AuthenUserRoutes(auth)
	handler.ProfileUserRoutes(user)

	r.Run(":8080")
}
