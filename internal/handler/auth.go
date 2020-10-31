package handler

import (
	"github.com/gin-gonic/gin"
	"github/trad3r/go_temp.git/internal/models"
	"net/http"
)

func (h Handler) Signin(c *gin.Context) {
	var user *models.User

	if err := c.BindJSON(&user); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "Not valid json data")
		return
	}

	if err := h.service.CreateUser(user); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": user.Id,
	})
}

func (h Handler) Login(c *gin.Context) {
	var user *models.User

	if err := c.BindJSON(&user); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "Not valid json data")
		return
	}

	token, err := h.service.GenerateToken(user)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
