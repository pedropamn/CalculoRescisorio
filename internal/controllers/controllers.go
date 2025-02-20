package controllers

import (
	"recisao/internal/entity"
	"time"
)

type Usecase interface {
	CalculateContractDetails(info entity.EmploymentContract, month time.Month, year int) (entity.ContractCalculationResults, error)
	CalculateTimeOfWork(info entity.EmploymentContract) (int, int, int, error)
}