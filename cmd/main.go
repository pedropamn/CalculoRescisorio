package main

import (
	"recisao/internal/adapter"
	"recisao/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

func main() {

	usecaseContract := usecase.NewUsecaseContract()
	handler := adapter.NewHandler(usecaseContract)

	app := fiber.New()
	app.Post("/", handler.HandlerPostContractDetails)

	app.Listen(":8080")
}
