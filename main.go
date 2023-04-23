package main

import (
	"Book-Inventory/app"
	"Book-Inventory/auth"
	"Book-Inventory/db"
	"Book-Inventory/middleware"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	conn := db.InitDB()

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	handler := app.New(conn)

	// home
	router.GET("/", auth.HomeHandler)

	// login
	router.GET("/login", auth.LoginGetHandler)
	router.POST("/login", auth.LoginPostHandler)

	// get all Books
	router.GET("/books", middleware.AuthValid, handler.GetBooks)

	router.Run()
}
