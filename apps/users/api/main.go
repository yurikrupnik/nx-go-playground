package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	go_fiber_helpers "nx-go-playground/libs/go/fiber-helpers"
	go_mongodb "nx-go-playground/libs/go/mongodb"
	"nx-go-playground/libs/go/myutils"
	"time"
)

//var mongoUrl string = "mongodb+srv://yurikrupnik:T4eXKj1RBI4VnszC@cluster0.rdmew.mongodb.net/"
//
////MongoInstance contains the Mongo client and database objects
//type MongoInstance struct {
//	Client *mongo.Client
//	Db     *mongo.Database
//}
//
//var mg MongoInstance
//
//// Connect configures the MongoDB client and initializes the database connection.
//// Source: https://www.mongodb.com/blog/post/quick-start-golang--mongodb--starting-and-setup
//func Connect() error {
//	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUrl))
//
//	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
//	defer cancel()
//
//	err = client.Connect(ctx)
//	db := client.Database("test")
//
//	if err != nil {
//		return err
//	}
//
//	mg = MongoInstance{
//		Client: client,
//		Db:     db,
//	}
//
//	return nil
//}

func UserRoute(app *fiber.App) {
	//other routes goes here
	co := app.Group("/api/v1", func(ctx *fiber.Ctx) error {
		ctx.Set("Version", "v1")
		return ctx.Next()
	})
	co.Get("/nir/:name?", func(c *fiber.Ctx) error {
		//return c.SendString("All good with nir api")
		//return c.SendString(c.Params("*"))
		return c.SendString(c.Params("name"))
	})

	// Second way
	//app.Route("/api", func(router fiber.Router) {
	//	router.Get("/shit", func(ctx *fiber.Ctx) error {
	//		return ctx.SendString("Hello, from shit!")
	//	})
	//})
	////app.Use("/api", func(c *fiber.Ctx) {
	////	//router.Get("/shit1", func(ctx *fiber.Ctx) error {
	////	return c.SendString("Hello, World!")
	////	//})
	////})
	//app.Get("/api/users", func(c *fiber.Ctx) error {
	//	// Pass error to Fiber
	//	return c.SendString("Hello, World!")
	//	//return c.SendFile("file-does-not-exist")
	//	//return c.JSON()
	//}) //add this
}

func FakeGroup(api fiber.Router) {
	//DoIt()
	router := api.Group("/users")
	router.Get("", func(ctx *fiber.Ctx) error {
		//return ctx.SendString("all good in new api 2")
		return ctx.SendString("some string")
	})
}

// User struct
type User struct {
	ID    string `json:"id,omitempty" bson:"_id,omitempty"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Role  string `json:"role"`
	//Role float64 `json:"salary"`
	//Age    float64 `json:"age"`
}
type UserResponse struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    *fiber.Map `json:"data"`
}

// @securityDefinitions.basic  BasicAuth
func main() {
	// Connect to the database
	if err := go_mongodb.Connect(); err != nil {
		log.Fatal(err)
	}
	app := fiber.New(fiber.Config{
		ErrorHandler: go_fiber_helpers.DefaultErrorHandler,
	})
	//client := mongoDb(mongoUrl)
	//mongoConfig.ConnectDB()
	app.Use(recover.New())
	// Default middleware config
	app.Use(logger.New())
	app.Use(csrf.New())
	app.Use(cors.New())
	apiGroup := app.Group("/api")

	UserRoute(app)
	FakeGroup(apiGroup)

	app.Get("/", func(c *fiber.Ctx) error {
		// get all records as a cursor
		query := bson.D{{}}
		cursor, err := go_mongodb.Mg.Db.Collection("users").Find(c.Context(), query)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		var users []User = make([]User, 0)

		// iterate the cursor and decode each item into an Employee
		if err := cursor.All(c.Context(), &users); err != nil {
			return c.Status(500).SendString(err.Error())

		}
		// return employees list in JSON format
		return c.JSON(users)
	})

	app.Get("/:id", func(c *fiber.Ctx) error {
		// get all records as a cursor
		query := bson.D{{}}
		cursor, err := go_mongodb.Mg.Db.Collection("users").Find(c.Context(), query)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		var users []User = make([]User, 0)

		// iterate the cursor and decode each item into an Employee
		if err := cursor.All(c.Context(), &users); err != nil {
			return c.Status(500).SendString(err.Error())

		}
		// return employees list in JSON format
		return c.JSON(users)
	})

	app.Put("/:id", func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Params("id")
		var user User
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(userId)

		//validate the request body
		if err := c.BodyParser(&user); err != nil {
			return c.Status(http.StatusBadRequest).JSON(UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}

		//use the validator library to validate required fields
		//if validationErr := validate.Struct(&user); validationErr != nil {
		//  return c.Status(http.StatusBadRequest).JSON(UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
		//}

		update := bson.M{"name": user.Name, "email": user.Email, "role": user.Role}

		result, err := go_mongodb.Mg.Db.Collection("users").UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}
		//get updated user details
		var updatedUser User
		if result.MatchedCount == 1 {
			err := go_mongodb.Mg.Db.Collection("users").FindOne(ctx, bson.M{"id": objId}).Decode(&updatedUser)

			if err != nil {
				return c.Status(http.StatusInternalServerError).JSON(UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
			}
		}

		return c.Status(http.StatusOK).JSON(UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": updatedUser}})
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
	})

	app.Delete("/:id", func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Params("id")
		fmt.Println("user id>>", userId)
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(userId)

		//result, err := userCollection.DeleteOne(ctx, bson.M{"id": objId})
		result, err := go_mongodb.Mg.Db.Collection("users").DeleteOne(ctx, bson.M{"id": objId})
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}

		if result.DeletedCount < 1 {
			return c.Status(http.StatusNotFound).JSON(
				UserResponse{Status: http.StatusNotFound, Message: "error", Data: &fiber.Map{"data": "User with specified ID not found!"}},
			)
		}

		return c.Status(http.StatusOK).JSON(
			UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": "User successfully deleted!"}},
		)
	})

	app.Get("/dashboard", monitor.New())

	port := go_myutils.Getenv("PORT", "8080")
	host := go_myutils.Getenv("HOST", "0.0.0.0")
	result := fmt.Sprintf("%s:%s", host, port)
	log.Panic(app.Listen(result))

}
