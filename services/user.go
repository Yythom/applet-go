package services

import (
	"fmt"
	"test/database"
	"test/domain"
	"test/tools"
)

func RegisterUser(params map[string]string) *domain.UserRegisterParams {
	data := &domain.UserRegisterParams{
		Username: params["username"],
		Password: params["password"],
	}

	bsonD, _ := tools.ConvertStructToBsonD(data)
	fmt.Println(bsonD, data)
	database.GetGlobalDB().InsertDocument(database.DB_NAME, "user", bsonD)
	return data
}
