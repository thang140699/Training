package services

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"mongo-with-golang/models"
	_ "mongo-with-golang/uploadfile"
)

type UserServiceImpl struct {
	usercollection *mongo.Collection
	ctx            context.Context
}

func NewUserService(usercollection *mongo.Collection, ctx context.Context) UserService {
	return &UserServiceImpl{
		usercollection: usercollection,
		ctx:            ctx,
	}
}
func (u *UserServiceImpl) CreateUser(user *models.SetTime) error {
	_, err := u.usercollection.InsertOne(u.ctx, user)
	return err
}

// db.collections.finc((name:0)) , get Domain
func (u *UserServiceImpl) GetDomain(name *string) (*models.SetTime, error) {
	var user *models.SetTime
	query := bson.D{bson.E{Key: "name", Value: name}}
	err := u.usercollection.FindOne(u.ctx, query).Decode(&user)
	return user, err
}

func (u *UserServiceImpl) GetAll() ([]*models.SetTime, error) {
	var users []*models.SetTime
	cursor, err := u.usercollection.Find(u.ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(u.ctx) {
		var user models.SetTime
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)

	}
	if err := cursor.Err(); err != nil {
		return nil, err

	}
	cursor.Close(u.ctx)
	if len(users) == 0 {
		return nil, errors.New("Document not found")
	}
	return users, nil

}
