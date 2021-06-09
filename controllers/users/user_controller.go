package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement Me!!!")
}

func CreateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement Me!!")
}

func FindUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement Me!!!")
}
