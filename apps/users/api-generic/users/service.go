package users

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userSvc[T any] struct {
	DB UserDB[T]
}

func NewUserSvc[T any](db UserDB[T]) UserSvc[T] {
	return userSvc[T]{
		DB: db,
	}
}

func (us userSvc[T]) Get(id primitive.ObjectID) (*T, error) {
	return us.DB.Get(id)
}

func (us userSvc[T]) List(u *T) ([]*T, error) {
	return us.DB.List(u)
}

func (us userSvc[T]) Create(user *T) (*mongo.InsertOneResult, error) {
	return us.DB.Create(user)
}

func (us userSvc[T]) Delete(id primitive.ObjectID) error {
	return us.DB.Delete(id)
}

func (us userSvc[T]) Update(id primitive.ObjectID, update *T) error {
	return us.DB.Update(id, update)
}
