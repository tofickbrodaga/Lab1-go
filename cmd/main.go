package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"s0709-22/internal/service"
)

func main() {
	api := fiber.New()
	svc := service.New()

	contact := api.Group("/api/v1/contact")
	contact.Post("/", svc.ContactCreate)
	contact.Get("/", svc.ContactRead)
	contact.Put("/", svc.ContactUpdate)
	contact.Delete("/", svc.ContactDelete)

	group := api.Group("/api/v1/group")
	group.Post("/", svc.GroupCreate)
	group.Get("/", svc.GroupRead)
	group.Put("/", svc.GroupUpdate)
	group.Delete("/", svc.GroupDelete)

	if err := api.Listen(":6080"); err != nil {
		log.Fatalln(err)
	}
}
