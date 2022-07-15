package users

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type mongoStore[T any] struct {
	Client *mongo.Client
}

func NewMongoStore[T any]() (UserDB[T], error) {
	//uri := os.Getenv("mongo_uri")
	uri := "mongodb+srv://yurikrupnik:T4eXKj1RBI4VnszC@cluster0.rdmew.mongodb.net/"
	//uri := "mongodb://db/profiles"
	//uri := "http://host.docker.internal:27017"
	//uri := "mongodb://localhost/db" // compose local gow run main.go
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Println("error connecting mongo")
		return nil, err
	}
	return mongoStore[T]{Client: client}, nil
}

func (ms mongoStore[T]) Get(id primitive.ObjectID) (*T, error) {
	col := ms.Client.Database("users").Collection("users")
	filter := bson.M{"_id": id}
	var a T
	err := col.FindOne(context.TODO(), filter).Decode(&a)
	if err != nil {
		return nil, err
	}
	return &a, nil
}

func (ms mongoStore[T]) List(u *T) ([]*T, error) {
	col := ms.Client.Database("users").Collection("users")
	log.Println("u", u)
	filter := bson.M{}
	log.Println("filter", filter)
	data := []*T{}
	cursor, err := col.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	if err := cursor.All(context.TODO(), &data); err != nil {
		return nil, err
	}
	return data, nil
}

func (ms mongoStore[T]) Delete(id primitive.ObjectID) error {
	col := ms.Client.Database("users").Collection("users")
	filter := bson.M{"_id": id}
	_, err := col.DeleteOne(context.TODO(), filter)
	return err
}

func (ms mongoStore[T]) Create(user *T) (*mongo.InsertOneResult, error) {
	col := ms.Client.Database("users").Collection("users")
	result, err := col.InsertOne(context.TODO(), user)
	return result, err
}

func (ms mongoStore[T]) Update(id primitive.ObjectID, user *T) error {
	col := ms.Client.Database("users").Collection("users")
	//update := bson.M{"name": user.Name, "email": user.Email, "role": user.Role}
	//fmt.Println(update)
	_, err := col.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": user})
	//log.Println("UpsertedID", result.UpsertedID)
	//_, err := col.UpdateOne(context.TODO(), user)
	return err
}
