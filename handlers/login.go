package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserLoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *BaseHandler) DoLogIn(c *gin.Context) {
	var userInput UserLoginInput
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dbUser, err := h.Repo.Queries.GetUserByName(h.Ctx, userInput.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(dbUser)

	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}
