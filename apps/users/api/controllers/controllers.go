package controllers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"nx-go-playground/libs/go/models/user"
	go_mongodb "nx-go-playground/libs/go/mongodb"
	"time"
)

func Create(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user go_models_user.User
	//var returnUser User
	defer cancel()

	//validate the request body
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	//use the validator library to validate required fields
	//if validationErr := validate.Struct(&user); validationErr != nil {
	//  return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	//}

	newUser := go_models_user.User{
		//ID:    primitive.NewObjectID(),
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}

	result, err := go_mongodb.Mg.Db.Collection("users").InsertOne(ctx, newUser)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	//err := go_mongodb.Mg.Db.Collection("users").FindOne(ctx, bson.M{"_id": result.InsertedID}).Decode(&returnUser)
	//return c.Status(http.StatusCreated).JSON(UserResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
	return c.Status(http.StatusCreated).JSON(result)

}
