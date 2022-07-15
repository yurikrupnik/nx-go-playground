package users

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type Handler[T any] struct {
	Svc UserSvc[T]
}

func NewHandler[T any](svc UserSvc[T]) *Handler[T] {
	return &Handler[T]{
		Svc: svc,
	}
}

func (h *Handler[T]) GetById(c *fiber.Ctx) error {
	id := c.Params("id")
	objId, _ := primitive.ObjectIDFromHex(id)
	data, err := h.Svc.Get(objId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(200).JSON(data)
}

func (h *Handler[T]) List(c *fiber.Ctx) error {
	u := new(T)
	if err := c.QueryParser(u); err != nil {
		return err
	}
	result, err := h.Svc.List(u)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

func (h *Handler[T]) Create(c *fiber.Ctx) error {
	var user T
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	result, err := h.Svc.Create(&user)
	if err != nil {
		fmt.Println(fmt.Errorf("error - adding new user detail to db failed, err : %v", err))
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"id": result.InsertedID})
}

func (h *Handler[T]) DeleteById(c *fiber.Ctx) error {
	id := c.Params("id")

	objId, _ := primitive.ObjectIDFromHex(id)
	log.Println("objId", objId)
	err := h.Svc.Delete(objId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(objId)
}

func (h *Handler[T]) UpdateById(c *fiber.Ctx) error {
	id := c.Params("id")
	objId, _ := primitive.ObjectIDFromHex(id)
	var user T
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	err := h.Svc.Update(objId, &user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	return c.SendString("updated")
}
