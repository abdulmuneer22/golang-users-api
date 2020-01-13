package users

import (
	"github.com/gin-gonic/gin"
	"golang-users-api/domain/users"
	"golang-users-api/services"
	"golang-users-api/utils/errors"
	"net/http"
	"strconv"
)



func CreateUser(c *gin.Context){
	var user users.User
	if err :=  c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid JSON Body")
		c.JSON(restErr.Status,restErr)
		return
	}
	result , saveError := services.CreateUser(user)
	if saveError != nil {
		c.JSON(saveError.Status,saveError)
		return
	}

	c.JSON(http.StatusCreated,result)
}

func GetUser(c *gin.Context){

	userId , userErr := strconv.ParseInt(c.Param("user_id"),10,64)

	if userErr != nil {
		err := errors.NewBadRequestError("Invalid User id, should be a number")
		c.JSON(err.Status , err)
	}

	user , getError := services.GetUser(userId)
	if getError != nil {
		c.JSON(getError.Status,getError)
		return
	}

	c.JSON(http.StatusOK,user)
}

func FindUser(c *gin.Context){
	c.String(http.StatusNotImplemented,"not implemented")
}
