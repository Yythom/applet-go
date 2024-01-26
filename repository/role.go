package repository

import (
	"applet/core/mongo"
	"applet/domain"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type roleRepository struct {
	database   mongo.Database
	collection string
}

func (r roleRepository) Create(c context.Context, role domain.Role) error {
	_, err := r.database.Collection(r.collection).InsertOne(c, role)
	if err != nil {
		return err
	}
	return nil
}

func (r roleRepository) Fetch(c context.Context) ([]domain.Role, error) {
	var roles []domain.Role

	cursor, err := r.database.Collection(r.collection).Find(c, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)

	if err := cursor.All(c, &roles); err != nil {
		return nil, err
	}

	return roles, nil
}

func (r roleRepository) GetByID(c context.Context, id string) (domain.Role, error) {
	var role domain.Role

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.Role{}, err
	}

	filter := bson.M{"_id": objectId}
	err = r.database.Collection(r.collection).FindOne(c, filter).Decode(&role)

	if err != nil {
		return domain.Role{}, err
	}

	return role, nil
}

func (r roleRepository) Update(c context.Context, id string, role domain.Role) (domain.Role, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.Role{}, err
	}

	filter := bson.M{"_id": objectId}
	update := bson.M{"$set": role}

	_, err2 := r.database.Collection(r.collection).UpdateOne(c, filter, update)
	if err2 != nil {
		return domain.Role{}, err2
	}

	result, err3 := r.GetByID(c, id)

	if err3 != nil {
		return domain.Role{}, err3
	}

	return result, nil
}

func (r roleRepository) DeleteByID(ctx context.Context, id string) error {
	// 将角色ID转换为 ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	// 构建删除的过滤条件
	filter := bson.M{"_id": objectID}

	// 使用 DeleteOne 方法删除匹配条件的文档
	result, err := r.database.Collection(r.collection).DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	// 检查是否成功删除了文档
	if result == 0 {
		return errors.New("role not found")
	}

	return nil
}

func NewRoleRepository(db mongo.Database, collection string) domain.RoleRepository {
	return &roleRepository{
		database:   db,
		collection: collection,
	}
}
