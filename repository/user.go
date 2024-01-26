package repository

import (
	"applet/core/mongo"
	"applet/domain"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepository struct {
	database   mongo.Database
	collection string
}

func (ur *userRepository) Update(c context.Context, id string, user *domain.User) (domain.User, error) {
	// 将传入的字符串形式的 ObjectID 转换为 primitive.ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.User{}, err
	}

	// 构建用于查询的 filter
	filter := bson.M{"_id": objectID}

	// 构建用于更新的 update
	update := bson.M{
		"$set": bson.M{
			"name":     user.Name,
			"email":    user.Email,
			"role_ids": user.RoleIds,
		},
	}

	// 执行更新操作
	result, err := ur.database.Collection(ur.collection).UpdateOne(c, filter, update)
	if err != nil {
		return domain.User{}, err
	}

	// 检查更新是否成功
	if result.ModifiedCount == 0 {
		// 如果 ModifiedCount 为 0，表示未找到匹配的文档需要更新
		return domain.User{}, errors.New("user not found or no fields to update")
	}

	// 查询更新后的用户信息
	updatedUser, err := ur.GetByID(c, id)
	if err != nil {
		return domain.User{}, err
	}

	return updatedUser, nil
}

func (ur *userRepository) Create(c context.Context, user *domain.User) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, user)

	return err
}

func (ur *userRepository) Fetch(c context.Context) ([]domain.User, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var users []domain.User

	err = cursor.All(c, &users)
	if users == nil {
		return []domain.User{}, err
	}

	return users, err
}

func (ur *userRepository) GetByName(c context.Context, name string) (domain.User, error) {
	collection := ur.database.Collection(ur.collection)
	var user domain.User
	err := collection.FindOne(c, bson.M{"name": name}).Decode(&user)
	return user, err
}

func (ur *userRepository) GetByID(c context.Context, id string) (domain.User, error) {
	collection := ur.database.Collection(ur.collection)

	var user domain.User

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&user)
	return user, err
}

func NewUserRepository(db mongo.Database, collection string) domain.UserRepository {
	return &userRepository{
		database:   db,
		collection: collection,
	}
}
