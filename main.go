package main

import (
	"net/http"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"

	"github.com/munaycacao/chocode/handlers"
	"github.com/munaycacao/chocode/repository"
	_ "github.com/go-sql-driver/mysql"
)



func main() {

	// Load .env file if it exists
	godotenv.Load()

	// Init repo and handlers
	repo := repository.NewRepo()
	h := handlers.NewHandler(repo)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/login", h.DoLogIn)
	r.POST("/user", h.CreateUser)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}