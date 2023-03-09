package repository

import (
	"context"
	"errors"

	"github.com/brunoseba/event-management/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepositoryImp struct {
	userCollection *mongo.Collection
	ctx            context.Context
}

func NewUserRepository(userCollection *mongo.Collection, ctx context.Context) UserRepositoryImp {
	return UserRepositoryImp{userCollection, ctx}
}

func (repo *UserRepositoryImp) CreateUser(user *entity.User) (*entity.UserResponse, error) {
	res, err := repo.userCollection.InsertOne(repo.ctx, user)
	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return nil, errors.New("the user is already exist")
		}
		return nil, errors.New("nose que paso")
	}
	var newUser *entity.UserResponse

	query := bson.M{"_id": res.InsertedID}
	if err = repo.userCollection.FindOne(repo.ctx, query).Decode(&newUser); err != nil {
		return nil, err
	}

	return newUser, nil

}

func (repo *UserRepositoryImp) GetUserRepo(name string) (*entity.UserResponse, error) {
	result := &entity.UserResponse{}
	query := bson.M{"nickname": name}

	err := repo.userCollection.FindOne(repo.ctx, query).Decode(result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &entity.UserResponse{}, err
		}
		return nil, err
	}

	return result, nil
}

func (repo *UserRepositoryImp) GetUserById(id string) (*entity.UserResponse, error) {
	oid, _ := primitive.ObjectIDFromHex(id)
	result := &entity.UserResponse{}
	query := bson.M{"_id": oid}

	err := repo.userCollection.FindOne(repo.ctx, query).Decode(result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &entity.UserResponse{}, err
		}
		return nil, err
	}

	return result, nil
}

func (repo *UserRepositoryImp) UpdateUser(id string, user *entity.UserUpdate) (*entity.UserResponse, error) {
	userId, _ := primitive.ObjectIDFromHex(id)
	query := bson.D{{Key: "_id", Value: userId}}
	update := bson.D{{Key: "$set", Value: user}}

	res := repo.userCollection.FindOneAndUpdate(repo.ctx, query, update, options.FindOneAndUpdate().SetReturnDocument(options.After))

	var updateUser *entity.UserResponse
	if err := res.Decode(&updateUser); err != nil {
		return nil, errors.New("could not update event")
	}
	return updateUser, nil
}

func (repo *UserRepositoryImp) DeleteUser(id string) error {

	userId, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": userId}

	_, err := repo.userCollection.DeleteOne(repo.ctx, query)
	if err != nil {
		return err
	}

	return nil
}
