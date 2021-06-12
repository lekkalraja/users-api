package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lekkalraja/users-api/domain/users"
	"github.com/lekkalraja/users-api/service"
	"github.com/lekkalraja/users-api/utils"
)

func GetUsers(c *gin.Context) {
	res, restErr := service.GetUsers()
	if restErr != nil {
		c.JSON(http.StatusInternalServerError, restErr)
		return
	}

	c.JSON(http.StatusOK, res)
}

func CreateUser(c *gin.Context) {
	user := &users.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewBadRequest(err.Error()))
		return
	}

	if vError := user.Validate(); vError != nil {
		c.JSON(http.StatusBadRequest, vError)
		return
	}

	savedUser, err := service.CreateUser(*user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, savedUser)
}

func FindUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewBadRequest("Invalid User Id"))
		return
	}
	user, restErr := service.FindUser(int64(id))
	if restErr != nil {
		c.JSON(http.StatusInternalServerError, restErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewBadRequest("Invalid User Id"))
		return
	}
	user, restErr := service.DeleteUser(int64(id))
	if restErr != nil {
		c.JSON(http.StatusInternalServerError, restErr)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": fmt.Sprintf("%d Rows Deleted", user)})
}
