package app

import (
	ping "github.com/lekkalraja/users-api/controllers/ping"
	users "github.com/lekkalraja/users-api/controllers/users"
)

func mapUrls() {
	// PING Mappings
	router.GET("/ping", ping.Ping)

	// USER Mappings
	router.POST("/user", users.CreateUser)

	router.GET("/users", users.GetUsers)
	router.GET("/user/:userId", users.FindUser)
	router.DELETE("/user/:userId", users.DeleteUser)
}
