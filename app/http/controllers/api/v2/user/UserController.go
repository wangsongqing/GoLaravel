package user

import (
	v2 "GoLaravel/app/http/controllers/api/v2"
	"GoLaravel/app/models/user"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserController struct {
	v2.BaseAPIController
}

func (uc *UserController) GetUser(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	userInfo := user.GetUser(userId)

	c.JSON(http.StatusOK, gin.H{
		"ID":     userInfo,
		"UserId": c.Param("id"),
	})
}
