package usecase

import (
	"fmt"
	"recisao/internal/controllers"
	"recisao/internal/entity"
	"strconv"
	"time"
)

type UsecaseContract struct {
}

func (u UsecaseContract) CalculateContractDetails(contractInfo entity.EmploymentContract, month time.Month, year int) (entity.ContractCalculationResults, error) {

	contractInfo.Validations()

	u.dataFormat(contractInfo)

	dailySalary, err := u.calculateDailySalary(contractInfo, month, year)
	if err != nil {
		return entity.ContractCalculationResults{}, err
	}

	years, months, days, err := u.CalculateTimeOfWork(contractInfo)
	if err != nil {
		return entity.ContractCalculationResults{}, err
	}

	Thirteenth, err := u.calculateProportionalThirteenthSalary(contractInfo)
	if err != nil {
		return entity.ContractCalculationResults{}, err
	}

	accruedHolidays, err := u.calculateAccruedHolidays(contractInfo)
	if err != nil {
		return entity.ContractCalculationResults{}, err
	}

	fgts, err := u.CalculoFGTS(contractInfo)
	if err != nil {
		return entity.ContractCalculationResults{}, err
	}
	results := entity.ContractCalculationResults{
		DailySalary:     dailySalary,
		Years:           years,
		Months:          months,
		Days:            days,
		Thirteenth:      Thirteenth,
		AccruedHolidays: accruedHolidays,
		Fgts:            fgts,
		Total:           Thirteenth + accruedHolidays + fgts + contractInfo.Salary,
	}

	return results, nil
}

func (u UsecaseContract) dataFormat(contractInfo entity.EmploymentContract) string {
	data := contractInfo.AdmissionDate
	dataFormatada := data.Format("2006-01-02")
	return dataFormatada
} // não funciona

func (u UsecaseContract) calculateDailySalary(contractInfo entity.EmploymentContract, month time.Month, year int) (string, error) {
	daysInMonth := daysInMonth(month, year)

	salaryPerDay := contractInfo.Salary / float64(daysInMonth)
	formattedSalary := fmt.Sprintf("%.1f", salaryPerDay)
	return formattedSalary, nil
}

func daysInMonth(month time.Month, year int) int {
	return time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()
}

func (u UsecaseContract) CalculateTimeOfWork(contractInfo entity.EmploymentContract) (int, int, int, error) {

	diference := contractInfo.DismissalDate.Sub(contractInfo.AdmissionDate)
	years := int(diference.Hours() / 24 / 365)
	months := int(diference.Hours()/24/30) % 12
	days := int(diference.Hours()/24) % 30

	return years, months, days, nil

}

func (u UsecaseContract) calculateProportionalThirteenthSalary(contractInfo entity.EmploymentContract) (float64, error) {

	thirteenthSalary := (contractInfo.Salary / 12) * float64(contractInfo.DismissalDate.Month())

	formattedThirteenthSalary := fmt.Sprintf("%.1f", thirteenthSalary)

	thirteenthSalaryFloat, err := strconv.ParseFloat(formattedThirteenthSalary, 64)

	if err != nil {
		return 0, err
	}

	return thirteenthSalaryFloat, nil
}

func (u UsecaseContract) calculateAccruedHolidays(contractInfo entity.EmploymentContract) (float64, error) {

	if contractInfo.TerminationType == entity.JustCause && contractInfo.AccumulatedVacationMonths < 12 {

		return 0, nil

	} else if contractInfo.AccumulatedVacationMonths < 12 {

		proportionalVacation := (contractInfo.AccumulatedVacationMonths * 30) / 12                //7,5
		proportionalVacationValue := (contractInfo.Salary / 30) * float64(proportionalVacation ) //375
		abono := proportionalVacationValue / 3
		total := proportionalVacationValue + abono

		formattedTotal := fmt.Sprintf("%.1f", total)

		formattedFloat, err := strconv.ParseFloat(formattedTotal, 64)

		if err != nil {
			return 0, err
		}

		return formattedFloat, nil

	} else if contractInfo.AccumulatedVacationMonths >= 12 {

		proportionalVacationValue := contractInfo.Salary
		abono := proportionalVacationValue / 3
		totalFor12Months := proportionalVacationValue + abono

		monthsWorked := contractInfo.AccumulatedVacationMonths % 12
		proportionalVacation := (contractInfo.Salary / 30) * float64((monthsWorked * 30 / 12))
		abonoProportional := proportionalVacation / 3
		totalProportional:= proportionalVacation + abonoProportional

		TotalPayable := totalFor12Months + totalProportional

		formattedTotalPayable := fmt.Sprintf("%.1f", TotalPayable)

		formattedFloat, err := strconv.ParseFloat(formattedTotalPayable , 64)

		if err != nil {
			return 0, err
		}

		return formattedFloat, nil

	}
	return 0, fmt.Errorf("não foi possível calcular as férias")
}

func (u UsecaseContract) CalculoFGTS(contractInfo entity.EmploymentContract) (float64, error) {

	if contractInfo.TerminationType != entity.WithoutCause {
		tempoDeTrabalho, _, _, _ := u.CalculateTimeOfWork(contractInfo)

		fgts := contractInfo.Salary * 8 / 100

		totalDeFgts := fgts * float64(tempoDeTrabalho)

		formattedTotalFgts := fmt.Sprintf("%.1f", totalDeFgts)

		formattedFloat, err := strconv.ParseFloat(formattedTotalFgts, 64)

		if err != nil {
			return 0, err
		}

		return formattedFloat, nil

	} else if contractInfo.TerminationType == entity.WithoutCause {

		WorkingTime, _, _, _ := u.CalculateTimeOfWork(contractInfo)

		fgts := contractInfo.Salary * 8 / 100

		totalDeFgts := fgts * float64(WorkingTime)

		finePercentage40  := totalDeFgts * 40 / 100

		totalReceive:= totalDeFgts + finePercentage40 

		formatted := fmt.Sprintf("%.1f", totalReceive)

		formattedFloat, err := strconv.ParseFloat(formatted, 64)

		if err != nil {
			return 0, err
		}

		return formattedFloat, nil

	}

	return 0, fmt.Errorf("não foi possível calcular o fgts")
}

func NewUsecaseContract() controllers.Usecase {
	return UsecaseContract{}
}
