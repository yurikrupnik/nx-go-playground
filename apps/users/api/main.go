package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"nx-go-playground/apps/users/api-generic/users"
	go_fiber_helpers "nx-go-playground/libs/go/fiber-helpers"
	"nx-go-playground/libs/go/models/user"
	go_mongodb "nx-go-playground/libs/go/mongodb"
	"nx-go-playground/libs/go/myutils"
	"os"
)

type UserResponseM struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    *fiber.Map `json:"data"`
}

type Project struct {
	ID   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name string             `json:"name" bson:"name,omitempty" validate:"required,min=3,max=36"`
}

// @securityDefinitions.basic  BasicAuth
func main() {
	// Connect to the database
	if err := go_mongodb.Connect(); err != nil {
		log.Fatal(err)
	}
	//if err := go_mongodb.NewDB(); err != nil {
	//	log.Fatal(err)
	//}
	app := fiber.New(fiber.Config{
		ErrorHandler: go_fiber_helpers.DefaultErrorHandler,
	})

	app.Use(recover.New())
	// Default middleware config
	app.Use(logger.New())
	//app.Use(csrf.New()) // todo check it - forbidden post events
	app.Use(cors.New())
	apiGroup := app.Group("api")
	apiGroup1 := app.Group("dss")
	apiGroup1.Get("/aris", func(ctx *fiber.Ctx) error {
		return ctx.SendString("ssks")
	})
	apiGroup1.Get("/dom", func(ctx *fiber.Ctx) error {
		return ctx.SendString("dosmd-s")
	})
	go_models_user.CreateFakeGroup[users.User](apiGroup, "users")
	go_models_user.CreateFakeGroup[Project](apiGroup, "projects")

	app.Get("/dashboard", monitor.New())
	ds := os.Getenv("MONGO_URI")
	fmt.Println("ds", ds)
	port := go_myutils.Getenv("PORT", "8080")
	log.Println("port", port)
	host := go_myutils.Getenv("HOST", "0.0.0.0")
	result := fmt.Sprintf("%s:%s", host, port)
	log.Panic(app.Listen(result))
}
