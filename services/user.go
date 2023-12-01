package services

import (
	"fmt"
	"test/database"
	"test/model"
	"test/tools"
)

func RegisterUser(params map[string]string) *model.UserRegisterParams {
	data := &model.UserRegisterParams{
		Username:        params["username"],
		Password:        params["password"],
		PasswordConfirm: params["password"],
	}
	bsonD, _ := tools.ConvertStructToBsonD(data)
	fmt.Println(bsonD, data)
	database.GetGlobalDB().InsertDocument(database.DB_NAME, "user", bsonD)
	return data
}
