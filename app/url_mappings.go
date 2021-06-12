package app

import (
	ping "github.com/lekkalraja/users-api/controllers/ping"
	users "github.com/lekkalraja/users-api/controllers/users"
)

func mapUrls() {
	// PING Mappings
	router.GET("/ping", ping.Ping)

	// USER Mappings
	router.POST("/user", users.Create)

	router.GET("/users", users.GetAll)
	router.GET("/user/:userId", users.Find)
	router.DELETE("/user/:userId", users.Delete)
	router.PUT("/user/:userId", users.Update)
}
