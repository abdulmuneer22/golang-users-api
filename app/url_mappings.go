package app

import (
	"golang-users-api/controllers/ping"
	"golang-users-api/controllers/users"
)

func mapUrls()  {
	router.GET("/ping", ping.Ping)

	router.POST("/users", users.CreateUser)

	router.GET("/users/:user_id", users.GetUser)

	//router.GET("/users/search",controllers.FindUser)
}
