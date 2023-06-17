package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func CheckUserType(c *gin.Context, role string) (err error){
	// getting the role type from the func below then processing it
	userType := c.GetString("user_type")
	err = nil
	if userType != role{
		err = errors.New("Nah you cant acess this one boi")
		return err
	}
	// this function will or will not return error
	return err
}

func MatchUserTypeToUid(c *gin.Context, userId string) (err error){
	userType := c.GetString("user_type")
	uid := c.GetString("uid")
	err = nil

	// Basically we are letting the user to access only his own data
	if userType == "USER" && uid != userId {
		err = errors.New("Unauthorized access")
		return err
	}
	// This is a hirachical way of checking user type idk what happening here i thought i would be able to hold my senses a little long
	err = CheckUserType(c, userType)
	// If the above function doesnt return any err this function will not return any error
	// basically saying that you have the passport and all the things require you are legit 
	return err
}
