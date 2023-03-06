package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/munaycacao/chocode/libs"
	"github.com/munaycacao/chocode/repository"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func (b *BaseHandler) CreateUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashPass, err := libs.HashPassword(user.Password)

	createUserParams := repository.CreateUserParams{
		Username: user.Username,
		Password: hashPass,
		Role:     user.Role,
	}

	_, err = b.Repo.Queries.CreateUser(b.Ctx, createUserParams)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dbUser, err := b.Repo.Queries.GetUserByName(b.Ctx, user.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dbUser.Password = ""

	c.JSON(http.StatusCreated, dbUser)
}