// database/db.go

package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

var (
	globalDB     *DBHandler
	globalDBOnce sync.Once
)

// initializeDB 初始化数据库连接
func initializeDB() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(DB_URL))
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return
	}

	globalDB = &DBHandler{
		client: client,
	}
}

// GetGlobalDB 获取全局数据库连接
func GetGlobalDB() *DBHandler {
	globalDBOnce.Do(initializeDB)
	return globalDB
}

// DBHandler 包含与数据库交互的方法
type DBHandler struct {
	client *mongo.Client
}
