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

func Create(c *gin.Context) {
	user, err := getRequestBody(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewBadRequest(err.Error()))
	}

	savedUser, savedErr := service.CreateUser(*user)
	if savedErr != nil {
		c.JSON(http.StatusInternalServerError, savedErr)
		return
	}

	c.JSON(http.StatusCreated, savedUser)
}

func Update(c *gin.Context) {
	user, err := getRequestBody(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewBadRequest(err.Error()))
	}
	id, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewBadRequest("Invalid User Id"))
		return
	}
	updatedRows, updateErr := service.UpdateUser(int64(id), *user)
	if updateErr != nil {
		c.JSON(http.StatusInternalServerError, updateErr)
		return
	}

	c.JSON(http.StatusCreated, map[string]string{"message": fmt.Sprintf("%d Rows Updated", updatedRows)})
}

func Find(c *gin.Context) {
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

func GetAll(c *gin.Context) {
	res, restErr := service.GetUsers()
	if restErr != nil {
		c.JSON(http.StatusInternalServerError, restErr)
		return
	}

	c.JSON(http.StatusOK, res)
}

func Delete(c *gin.Context) {
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

func getRequestBody(c *gin.Context) (*users.User, error) {
	user := &users.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		return nil, err
	}
	user.Format()
	return user, nil
}
