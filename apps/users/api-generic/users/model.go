package users

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name  string             `json:"name" bson:"name,omitempty" validate:"required,min=3,max=36"`
	Role  string             `json:"role" bson:"role,omitempty"`
	Age   int                `json:"age" bson:"age,omitempty" validate:"required,numeric" query:"age"`
	Email string             `json:"email" bson:"email,omitempty" validate:"required,email" query:"email"`
}

type UserSvc[T any] interface {
	Get(id primitive.ObjectID) (*T, error)
	List(item *T) ([]*T, error)
	Create(item *T) (*mongo.InsertOneResult, error)
	Delete(id primitive.ObjectID) error
	Update(id primitive.ObjectID, item *T) error
}

type UserDB[T any] interface {
	Get(id primitive.ObjectID) (*T, error)
	List(item *T) ([]*T, error)
	Create(item *T) (*mongo.InsertOneResult, error)
	Delete(id primitive.ObjectID) error
	Update(id primitive.ObjectID, item *T) error
}
