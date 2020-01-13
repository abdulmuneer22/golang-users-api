// used from transfer data from persistence layer to application layer

package users

import (
	"fmt"
	"golang-users-api/utils/errors"
)

var (
	userDB = make(map[int64]*User)
)

func (user *User) Get()  *errors.RestErr{
	result := userDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("User does not exist %d", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return  nil
}


func (user *User) Save() *errors.RestErr{
	current := userDB[user.Id]
	if  current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError("This email address is already registered")
		}
		return  errors.NewBadRequestError(fmt.Sprintf("User already exists %d" , user.Id))
	}
	userDB[user.Id] = user
	return  nil
}
