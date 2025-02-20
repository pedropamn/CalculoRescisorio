package adapter

import (
	"recisao/internal/controllers"
	"recisao/internal/entity"
	"time"

	"github.com/gofiber/fiber/v2"
)

type ContractHandler struct {
	usecase controllers.Usecase
}

func (c ContractHandler) HandlerPostContractDetails(ctx *fiber.Ctx) error {
	var contractInfo entity.EmploymentContract

	err := ctx.BodyParser(&contractInfo)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error":   "invalid request body",
			"message": err.Error(),
		})
	}

	month, years, _, err := c.usecase.CalculateTimeOfWork(contractInfo)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error":   "Invalid time calculation",
			"message": err.Error(),
		})
	}

	contract, err := c.usecase.CalculateContractDetails(contractInfo, time.Month(month), years)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error":   "Internal server error",
			"message": err.Error(),
		})
	}
	return ctx.Status(200).JSON(contract)
}

func NewHandler(u controllers.Usecase) ContractHandler {
	return ContractHandler{usecase: u}
}
