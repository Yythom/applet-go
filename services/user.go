package services

import (
	"fmt"
	"test/database"
	"test/model"
	"test/tools"
)

func RegisterUser(params map[string]string) *model.UserInfoType {
	data := &model.UserInfoType{
		Username:         params["username"],
		Password:         params["password"],
		Gender:           "",
		Birthday:         "",
		Address:          "",
		LastLogin:        "",
		AccountLocked:    false,
		RegistrationTime: "",
		UserType:         "",
	}

	bsonD, _ := tools.ConvertStructToBsonD(data)
	fmt.Println(bsonD, data)
	database.GetGlobalDB().InsertDocument(database.DB_NAME, "user", bsonD)
	return data
}
