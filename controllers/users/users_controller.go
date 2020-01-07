package users

import (
	"bookstore_users-api/domain/users"
	"bookstore_users-api/services"
	"github.com/gin-gonic/gin"
	"net/http"
)



func CreateUser(c *gin.Context){
	var user users.User
	if err :=  c.ShouldBindJSON(&user); err != nil {
		// TODO Error handler here
	}
	result , saveError := services.CreateUser(user)
	if saveError != nil {
		// Handle User Creation Error
		return
	}

	c.JSON(http.StatusCreated,result)
}

func GetUser(c *gin.Context){
	c.String(http.StatusNotImplemented,"not implemented")
}

func FindUser(c *gin.Context){
	c.String(http.StatusNotImplemented,"not implemented")
}