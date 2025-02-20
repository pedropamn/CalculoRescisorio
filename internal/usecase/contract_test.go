package usecase

import (
	"recisao/internal/entity"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// ajustar lógica para que ao enves de fazer comparações diretas passe a verificar corretamente as lógicas

func Test_Validations(t *testing.T) {
	assert.New(t)

	admissionDate := time.Date(2024,time.January,1,0,0,0,0,time.UTC)
	dismissionDate := time.Date(2025,time.January,1,0,0,0,0,time.UTC)
	salary := 10.0
	terminationType := entity.WithoutCause

	info := entity.EmploymentContract{}
	//uc := UsecaseContract{}
	uc := info.Validations()
	t.Run("missing admission date", func(t *testing.T) {
		assert.ErrorContains(t, uc,
			"data de admissão não pode ser vazio")
	})

	t.Run("missing dismissal date", func(t *testing.T) {
		info.AdmissionDate = admissionDate
		assert.ErrorContains(t, uc, "data de demissão não pode ser vazio")
	})

	t.Run("dismissal date before admission date", func(t *testing.T) {
		info.AdmissionDate = admissionDate
		info.DismissalDate = dismissionDate.AddDate(-1, 0, 0)
		assert.ErrorContains(t, uc, "a data de demissão não pode ser anterior á data de admissão")
	})

	t.Run("missing salary", func(t *testing.T) {
		info.AdmissionDate = admissionDate
		info.DismissalDate = dismissionDate
		assert.ErrorContains(t, uc, "salario deve ser maior que zero")
	})

	t.Run("not just cause or without cause", func(t *testing.T) {
		info.AdmissionDate = admissionDate
		info.DismissalDate = dismissionDate
		info.Salary = salary
		assert.ErrorContains(t, uc, "tipo de recisão inválido")
	})

	t.Run("success", func(t *testing.T) {
		info.AdmissionDate = admissionDate
		info.DismissalDate = dismissionDate
		info.Salary = salary
		info.TerminationType = terminationType
		assert.Nil(t, info.Validations())
	})

}
