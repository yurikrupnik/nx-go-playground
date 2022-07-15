package users

import (
	"github.com/gofiber/fiber/v2"
)

type IError struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

//var Validator = validator.New()

//type Op func(v *Validate) (s interface{}, fields string) error

//func CreateAndUpdateValidation(partial bool) fiber.Handler {
//	return func(c *fiber.Ctx) error {
//		var errors []*IError
//		body := new(User)
//		c.BodyParser(&body)
//		if partial {
//			err := Validator.StructPartial(body)
//			if err != nil {
//				for _, err := range err.(validator.ValidationErrors) {
//					var el IError
//					el.Field = err.Field()
//					el.Tag = err.Tag()
//					el.Value = err.Param()
//					log.Println(err)
//					errors = append(errors, &el)
//				}
//				return c.Status(fiber.StatusBadRequest).JSON(errors)
//			}
//		} else {
//			err := Validator.Struct(body)
//			if err != nil {
//				for _, err := range err.(validator.ValidationErrors) {
//					var el IError
//					el.Field = err.Field()
//					el.Tag = err.Tag()
//					el.Value = err.Param()
//					log.Println(err)
//					errors = append(errors, &el)
//				}
//				return c.Status(fiber.StatusBadRequest).JSON(errors)
//			}
//		}
//		return c.Next()
//	}
//}

func Routes[T any](api fiber.Router, name string, h *Handler[T]) {
	//r.Route("/users", func(r fiber.Router) {
	router := api.Group(name)
	router.Post("/", h.Create)
	router.Get("", h.List)
	router.Get("/:id", h.GetById)
	router.Delete("/:id", h.DeleteById)
	router.Put("/:id", h.UpdateById)
	//r.With(ValidateQueryParam(ListPetsQuery{})).Get("/", h.ListPets)
	//r.With(ValidateBody(domain.Pet{})).Post("/", h.AddPet)
	//r.Route("/{id}", func(r chi.Router) {
	//	r.Use(ValidateURLParam("id"))
	//})
	//})
}
