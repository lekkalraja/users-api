package app

import controllers "github.com/lekkalraja/users-api/controllers/ping"

func mapUrls() {
	router.GET("/ping", controllers.Ping)
}
