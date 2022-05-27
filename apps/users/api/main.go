package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
	go_fiber_helpers "nx-go-playground/libs/go/fiber-helpers"
	"nx-go-playground/libs/go/models/user"
	go_mongodb "nx-go-playground/libs/go/mongodb"
	"nx-go-playground/libs/go/myutils"
)

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
	go_models_user.CreateFakeGroup(apiGroup, "users")

	app.Get("/dashboard", monitor.New())

	port := go_myutils.Getenv("PORT", "8080")
	host := go_myutils.Getenv("HOST", "0.0.0.0")
	result := fmt.Sprintf("%s:%s", host, port)
	log.Panic(app.Listen(result))
}
