package users

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func New(r fiber.Router, url string) {
	s, err := NewMongoStore[User]()
	if err != nil {
		//return errors.Wrap(err, "unable to connect to db")
	}
	srv := NewUserSvc[User](s)
	log.Print("all good")
	handlers := NewHandler[User](srv)
	Routes(r, url, handlers)
}
