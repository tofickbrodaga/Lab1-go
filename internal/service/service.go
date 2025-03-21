package service

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type Service struct{}

func New() *Service {
	return new(Service)
}

func (s *Service) RootHandler(ctx *fiber.Ctx) error {
	log.Println("hello")

	return ctx.SendString("Hello!!!")
}

func (s *Service) UserRead(ctx *fiber.Ctx) error {
	log.Println("user")

	return ctx.SendString("User Data")
}

// ProfileRead
func (s *Service) ProfileRead(ctx *fiber.Ctx) error {
	log.Println("profile")

	return ctx.SendString("Profile Data")
}

func (s *Service) ContactCreate(ctx *fiber.Ctx) error {
	var contact Contact
	if err := ctx.BodyParser(&contact); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusCreated).JSON(Contact{})
}

func (s *Service) ContactRead(ctx *fiber.Ctx) error {
	return ctx.JSON(Contact{})
}

func (s *Service) ContactUpdate(ctx *fiber.Ctx) error {
	var contact Contact
	if err := ctx.BodyParser(&contact); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(Contact{})
}

func (s *Service) ContactDelete(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusNoContent)
}

func (s *Service) GroupCreate(ctx *fiber.Ctx) error {
	var group Group
	if err := ctx.BodyParser(&group); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusCreated).JSON(Group{})
}

func (s *Service) GroupRead(ctx *fiber.Ctx) error {
	return ctx.JSON(Group{})
}

func (s *Service) GroupUpdate(ctx *fiber.Ctx) error {
	var group Group
	if err := ctx.BodyParser(&group); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(Group{})
}

func (s *Service) GroupDelete(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusNoContent)
}
