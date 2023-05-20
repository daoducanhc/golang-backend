package main

import (
	"log"
	"std/middlewares"
	"std/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	r.Static("/pictures", "./pictures")
	r.LoadHTMLGlob("templates/*.html")
	// r.Use(static.Serve("/pictures", static.LocalFile("/pictures", true)))

	auth := r.Group("/auth")

	auth.POST("/login", routes.Login)

	protected := r.Group("/user")

	protected.Use(middlewares.JwtAuthMiddleware())

	protected.POST("/:username/nickname/edit", routes.ChangeNickname)
	protected.GET("/nickname/edit", routes.EditNickname)
	protected.GET("/profile", routes.GetProfile)
	protected.PUT("/picture", routes.ChangePicture)

	r.Run(":8080")
}
