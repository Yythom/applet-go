package database

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NewDBHandler 创建一个新的DBHandler实例
func NewDBHandler(connectionString string) (*DBHandler, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, fmt.Errorf("error connecting to MongoDB: %v", err)
	}
	return &DBHandler{
		client: client,
	}, nil
}

// Close 关闭数据库连接
func (db *DBHandler) Close() error {
	return db.client.Disconnect(context.Background())
}

// Collection 获取指定集合
func (db *DBHandler) Collection(dbName, collectionName string) *mongo.Collection {
	return db.client.Database(dbName).Collection(collectionName)
}

// InsertDocument 插入文档
func (db *DBHandler) InsertDocument(dbName, collectionName string, document interface{}) error {
	collection := db.Collection(dbName, collectionName)
	_, err := collection.InsertOne(context.Background(), document)
	if err != nil {
		return fmt.Errorf("error inserting document: %v", err)
	}
	return nil
}

// GetDocument 根据查询条件获取单个文档
func (db *DBHandler) GetDocument(dbName, collectionName string, filter bson.D, result interface{}) error {
	collection := db.Collection(dbName, collectionName)
	err := collection.FindOne(context.Background(), filter).Decode(result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return fmt.Errorf("document not found")
		}
		return fmt.Errorf("error querying document: %v", err)
	}
	return nil
}

// UpdateDocument 更新文档
func (db *DBHandler) UpdateDocument(dbName, collectionName string, filter bson.D, update bson.D) error {
	collection := db.Collection(dbName, collectionName)
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return fmt.Errorf("error updating document: %v", err)
	}
	return nil
}

// DeleteDocument 删除文档
func (db *DBHandler) DeleteDocument(dbName, collectionName string, filter bson.D) error {
	collection := db.Collection(dbName, collectionName)
	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("error deleting document: %v", err)
	}
	return nil
}
