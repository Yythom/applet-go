package repository

import (
	"applet/core/mongo"
	"applet/domain"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type permissionRepository struct {
	database   mongo.Database
	collection string
}

func (p permissionRepository) Create(c context.Context, permission domain.CreatePermissionForRepository) error {
	_, err := p.database.Collection(p.collection).InsertOne(c, permission)
	if err != nil {
		return err
	}
	return nil
}

func (p permissionRepository) Fetch(c context.Context) ([]domain.Permission, error) {
	var permissions []domain.Permission

	cursor, err := p.database.Collection(p.collection).Find(c, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)

	if err := cursor.All(c, &permissions); err != nil {
		return nil, err
	}

	return permissions, nil
}

func (p permissionRepository) GetByID(c context.Context, id string) (*domain.Permission, error) {
	var permission *domain.Permission

	// 将传入的字符串形式的 ObjectID 转换为 primitive.ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return permission, err
	}

	filter := bson.M{"_id": objectID}
	err = p.database.Collection(p.collection).FindOne(c, filter).Decode(&permission)

	if err != nil {
		return permission, err
	}

	return permission, nil
}

func (p permissionRepository) GetByIDs(c context.Context, ids []string) ([]domain.Permission, error) {
	var permissions []domain.Permission

	objectIDs := make([]primitive.ObjectID, len(ids))

	for i, id := range ids {
		objectID, _ := primitive.ObjectIDFromHex(id)

		objectIDs[i] = objectID
	}

	filter := bson.M{"_id": objectIDs}
	cursor, err := p.database.Collection(p.collection).Find(c, filter)

	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)

	if err := cursor.All(c, &permissions); err != nil {
		return nil, err
	}

	if err != nil {
		return []domain.Permission{}, err
	}

	return permissions, nil
}

func (p permissionRepository) Update(c context.Context, permission domain.UpdatePermissionRequest) (*domain.Permission, error) {
	var newPermission *domain.Permission
	// 将传入的字符串形式的 ObjectID 转换为 primitive.ObjectID
	objectID, err := primitive.ObjectIDFromHex(permission.ID)
	if err != nil {
		return newPermission, err
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": permission}

	_, err2 := p.database.Collection(p.collection).UpdateOne(c, filter, update)
	if err2 != nil {
		return newPermission, err2
	}

	result, err3 := p.GetByID(c, permission.ID)

	if err3 != nil {
		return newPermission, err3
	}

	return result, nil
}

func (p permissionRepository) DeleteByID(c context.Context, id string) error {
	// 将字符串 ID 转换为 ObjectID（假设你的 ID 是 MongoDB 的 ObjectID）
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	// 根据 ID 删除权限
	filter := bson.M{"_id": objectID}
	result, err := p.database.Collection(p.collection).DeleteOne(c, filter)
	if err != nil {
		return err
	}

	if result == 0 {
		return fmt.Errorf("permission not found with ID: %s", id)
	}

	return nil
}

func (p permissionRepository) GetByName(c context.Context, name string) (*domain.Permission, error) {
	// 根据名称查询权限
	filter := bson.M{"permission_name": name}
	result := p.database.Collection(p.collection).FindOne(c, filter)

	// 解码结果到权限结构体
	var permission domain.Permission
	err := result.Decode(&permission)
	if err != nil {
		return nil, err
	}

	return &permission, nil
}

func NewPermissionRepository(db mongo.Database, collection string) domain.PermissionRepository {
	return &permissionRepository{
		database:   db,
		collection: collection,
	}
}
