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
	//s.
	srv := NewUserSvc[User](s)
	//srv.
	log.Print("all good")
	handlers := NewHandler[User](srv)
	//handlers.
	Routes[User](r, url, handlers)
}
