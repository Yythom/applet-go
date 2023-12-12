package tools

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
)

func ConvertStructToBsonD(data interface{}) (bson.D, error) {
	value := reflect.ValueOf(data)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if value.Kind() != reflect.Struct {
		return nil, fmt.Errorf("input is not a struct")
	}

	var bsonD bson.D

	for i := 0; i < value.NumField(); i++ {
		field := value.Type().Field(i).Name
		fieldValue := value.Field(i).Interface()

		// 打印字段名和字段值，用于调试
		fmt.Printf("Field: %s, Value: %v\n", field, fieldValue)

		// 如果你的结构体字段的类型是自定义的，你可能需要添加更多的类型检查和处理逻辑
		// 在这个例子中，我们简单地将字段名和字段值添加到bson.D中
		bsonD = append(bsonD, bson.E{Key: field, Value: fieldValue})
	}

	return bsonD, nil
}
