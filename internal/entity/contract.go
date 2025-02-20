package entity

import (
	"errors"
	"time"
)

type TerminationType string

const (
	JustCause    TerminationType = "justa causa"
	WithoutCause TerminationType = "sem justa causa"
	Layoff       TerminationType = "demissão"
)

type EmploymentContract struct {
	AdmissionDate             time.Time       `json:"admission_date"`
	DismissalDate             time.Time       `json:"dismissal_date"`
	Salary                    float64         `json:"salary"`
	TerminationType           TerminationType `json:"termination_type"`
	AccumulatedVacationMonths int             `json:"accumulate_vacation_Months"` 

}

type ContractCalculationResults struct {
	DailySalary     string  `json:"daily_salary"`
	Years           int     `json:"years"`
	Months          int     `json:"months"`
	Days            int     `json:"days"`
	Thirteenth      float64 `json:"thirteenth"`
	AccruedHolidays float64 `json:"accrued_holidays"`
	Fgts            float64 `json:"fgts"`
	Total           float64 `json:"total"`
}

func (e EmploymentContract) Validations() error {

	var contractInfo EmploymentContract

	if contractInfo.AdmissionDate.IsZero() {
		return errors.New("data de admissão não pode ser vazio")
	}

	if contractInfo.DismissalDate.IsZero() {
		return errors.New("data de demissão não pode ser vazio")
	}

	if contractInfo.DismissalDate.Before(contractInfo.AdmissionDate) {
		return errors.New("a data de demissão não pode ser anterior à data de admissão")
	}

	if contractInfo.Salary <= 0 {
		return errors.New("salário deve ser maior que zero")
	}

	if contractInfo.TerminationType != JustCause && contractInfo.TerminationType != WithoutCause && contractInfo.TerminationType != Layoff {
		return errors.New("tipo de recisão inválido, expectativa: justa causa, sem justa causa ou demissão")
	}

	return nil
}
