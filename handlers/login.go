package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/munaycacao/chocode/libs"
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

	if isTrue := libs.CheckPasswordHash(userInput.Password, dbUser.Password); isTrue {
		token := libs.GenerateToken(dbUser.ID)
		c.JSON(http.StatusOK, gin.H{"msg": "You are logged in", "token": token})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{"msg": "Password not matched"})
	return

}
