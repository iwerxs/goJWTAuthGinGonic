package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func CheckUserType(c *gin.Context, role string) (err error){
	userType := c.GetString("user_type")
	err = nil
	if userType != role {
		err = errors.New("unathorized to access this resource")
		return err
	}
	return err
}

func MatchUserTypeToUid(c *gin.Context, userId string) (err error) {
	userType := c.GetString("user_type")
	uid := c.GetString("uid")
	err = nil

	if userType == "USER" && uid != userId {
		err = errors.New("unauthorized to access this resource")
		return err
	}
	// add CAdmin, PAdmin via if conditionals

	// helper function
	err = CheckUserType(c, userType)
	return err
}