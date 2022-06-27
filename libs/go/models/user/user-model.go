package go_models_user

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	go_mongodb "nx-go-playground/libs/go/mongodb"
	"time"
)

type circle struct {
	width  float64
	height float64
}

func (u User) getOne() string {
	//u.
	return u.Name
}

func (c circle) area() float64 {
	return c.width + c.height
}

type ApiMethods struct {
	create  func()
	getById func()
	getList func()
	remove  func()
	update  func()
}

type User struct {
	ID    string `json:"id,omitempty" bson:"_id,omitempty"`
	Email string `json:"email,omitempty" bson:"email,omitempty"`
	Name  string `json:"name"`
	Role  string `json:"role"`
	//Role float64 `json:"salary"`
	//Age    float64 `json:"age"`
}

type UserStore interface {
	Insert(item interface{}) error
	Get(id int) error
}
type UserService struct {
	store UserStore
}

// Accepting interface here!
func NewUserService(s UserStore) *UserService {
	return &UserService{
		store: s,
	}
}

//func (u *UserService) CreateUser()              {}
//func (u *UserService) RetrieveUser(id int) User {}

func updateOneMove(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"watched": true}

	result, err := go_mongodb.Mg.Db.Collection("users").UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("modified count:", result.ModifiedCount)
}

func CreateFakeGroup(api fiber.Router, name string) {
	router := api.Group(name)
	router.Post("", CreateHandlerCreate[User]())
	router.Get("", CreateHandlerList())
	router.Get("/:id", CreateHandleGetById())
	router.Delete("/:id", CreateHandleDeleteById())
	router.Put("/:id", CreateHandleUpdate())
}

func CreateHandlerCreate[T interface{}]() fiber.Handler {

	return func(c *fiber.Ctx) error {
		var user User
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		//validate the request body
		if err := c.BodyParser(&user); err != nil {
			return c.Status(http.StatusBadRequest).JSON(err.Error())
		}
		fmt.Println(">>>>>>user", user)

		//use the validator library to validate required fields
		//if validationErr := validate.Struct(&user); validationErr != nil {
		//  return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
		//}

		newUser := User{
			Name:  user.Name,
			Email: user.Email,
			Role:  user.Role,
		}
		result, err := go_mongodb.Mg.Db.Collection("users").InsertOne(ctx, newUser)
		if err != nil {
			//return c.Status(http.StatusInternalServerError).JSON(UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}
		fmt.Println("result", result)
		fmt.Println(">>>>>>user", user)
		//err := go_mongodb.Mg.Db.Collection("users").FindOne(ctx, bson.M{"_id": result.InsertedID}).Decode(&returnUser)
		//return c.Status(http.StatusCreated).JSON(UserResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
		return c.Status(http.StatusCreated).JSON(result)
	}
}

func CreateHandlerList() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// get all records as a cursor
		//query := bson.D{{}}
		var usr User
		//var myStrings SensorReading
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		//todo handle query
		data := c.Request().URI().QueryString()
		//args := c.Request().URI().QueryArgs()
		//limitVal, _ := strconv.Atoi(c.Query("limit", "10"))
		//DecodeJSON(data, &usr)
		//if err != nil {
		//	fmt.Println("error")
		//}
		//jsons, err := json.Marshal(data)
		//err := json.Unmarshal([]byte(data), &usr)
		//if err != nil {
		//	fmt.Println("error parsing json")
		//}
		fmt.Println("aris")
		fmt.Println("usr => ", usr)
		//fmt.Println("das => ", das)
		//fmt.Println("jsons => ", jsons)
		fmt.Println("data => ", data)
		fmt.Println("String => ", string(data))
		//fmt.Println(map)
		query := bson.M{} // todo handle query
		//query := bson.M{} // todo handle query
		fmt.Println(c.Query("email"))
		fmt.Println(c.Query("*"))
		//query := bson.M{"email": "a@a.com"}
		cursor, err := go_mongodb.Mg.Db.Collection("users").Find(ctx, query)
		if err != nil {
			return c.Status(http.StatusInternalServerError).SendString(err.Error())
		}
		//defer cursor.Close(ctx)
		var users []User = make([]User, 0)

		// iterate the cursor and decode each item into an Employee
		//if err := cursor.All(c.Context(), &users); err != nil {
		//  return c.Status(http.StatusInternalServerError).SendString(err.Error())
		//}

		for cursor.Next(ctx) {
			var user User
			err := cursor.Decode(&user)
			if err != nil {
				fmt.Println(err)
				//return err
				//log.Fatal("fatal")
				//log.Panic("panic")
			}
			users = append(users, user)
		}
		// return employees list in JSON format
		return c.Status(http.StatusOK).JSON(users)
	}
}

func CreateHandleGetById() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		id := c.Params("id")
		var user User
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(id)

		err := go_mongodb.Mg.Db.Collection("users").FindOne(ctx, bson.M{"_id": objId}).Decode(&user)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(err.Error())
		}

		//return c.Status(http.StatusOK).JSON(UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": user}})
		return c.Status(fiber.StatusOK).JSON(user)
	}
}

func CreateHandleDeleteById() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Params("id")
		fmt.Println("user id>>", userId)
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(userId)

		//result, err := userCollection.DeleteOne(ctx, bson.M{"id": objId})
		result, err := go_mongodb.Mg.Db.Collection("users").DeleteOne(ctx, bson.M{"_id": objId})
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(err.Error())
		}

		if result.DeletedCount < 1 {
			return c.Status(http.StatusNotFound).JSON(
				&fiber.Map{"data": "User with specified ID not found!"},
			)
		}

		return c.Status(http.StatusOK).JSON(
			&fiber.Map{"data": "User successfully deleted!"},
		)
	}
}

func CreateHandleUpdate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		id := c.Params("id")
		var user User
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(id)
		fmt.Println(objId)
		//validate the request body
		if err := c.BodyParser(&user); err != nil {
			return c.Status(http.StatusBadRequest).JSON(err.Error())
		}

		//use the validator library to validate required fields
		//if validationErr := validate.Struct(&user); validationErr != nil {
		//  return c.Status(http.StatusBadRequest).JSON(UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
		//}

		// map
		update := bson.M{"name": user.Name, "email": user.Email, "role": user.Role}

		fmt.Println(update)
		result, err := go_mongodb.Mg.Db.Collection("users").UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})

		if err != nil {
			fmt.Println(err)
			return c.Status(http.StatusInternalServerError).JSON(err.Error())
		}
		//get updated user details
		var updatedUser User
		//fmt.Println("result", result.)
		if result.MatchedCount == 1 {
			fmt.Println("dmmm")
			err := go_mongodb.Mg.Db.Collection("users").FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedUser)

			if err != nil {
				fmt.Println("dmmm11")
				return c.Status(http.StatusInternalServerError).JSON(err.Error())
			}
		}

		return c.Status(http.StatusOK).JSON(updatedUser)
		//result, err := userCollection.DeleteOne(ctx, bson.M{"id": objId})
		//result, err := go_mongodb.Mg.Db.Collection("users").DeleteOne(ctx, bson.M{"id": objId})
		//if err != nil {
		//  return c.Status(http.StatusInternalServerError).JSON(UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		//}
		//
		//if result.DeletedCount < 1 {
		//  return c.Status(http.StatusNotFound).JSON(
		//    UserResponse{Status: http.StatusNotFound, Message: "error", Data: &fiber.Map{"data": "User with specified ID not found!"}},
		//  )
		//}
		//
		//return c.Status(http.StatusOK).JSON(
		//  UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": "User successfully deleted!"}},
		//)
	}
}
