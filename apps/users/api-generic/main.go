package main

import (
	//"fiber-mongo/app"
	//"fiber-mongo/db"
	//ihttp "fiber-mongo/http"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"nx-go-playground/apps/users/api-generic/users"
	go_myutils "nx-go-playground/libs/go/myutils"
	"os"
)

func main() {
	args := os.Args
	op := "servers"
	if len(args) > 1 {
		op = args[0]
	}
	if err := run(op); err != nil {
		fmt.Println(fmt.Errorf("error - server failed to start. err: %v", err))
	}
}

func run(op string) error {
	//d, err := db.NewMongoStore()
	//if err != nil {
	//  return errors.Wrap(err, "unable to connect to db")
	//}
	//svc := app.NewUserSvc(d)
	//
	application := fiber.New(fiber.Config{
		//JSONEncoder: json.Marshal,
		//JSONDecoder: json.Unmarshal,
	})
	application.Use(logger.New())
	//h := ihttp.NewHandler(svc)
	apiGroup := application.Group("api")

	//users.Routes(apiGroup, "user")
	//ihttp.Routes(apiGroup, h)
	users.New(apiGroup, "users")
	port := go_myutils.Getenv("PORT", "8080")
	log.Println("port", port)
	host := go_myutils.Getenv("HOST", "0.0.0.0")
	result := fmt.Sprintf("%s:%s", host, port)
	return application.Listen(result)
}
