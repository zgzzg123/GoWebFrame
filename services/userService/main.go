package userService

import (
	"goDoc/models/userModel"
)

func Create(params map[string]interface{}) interface{} {
	//return userModel.CreateUnique(params.(userModel.DataProfiles))
	return nil
}

func GetList(params map[string]interface{}) interface{} {
	result := userModel.GetList(params)

	return result
}