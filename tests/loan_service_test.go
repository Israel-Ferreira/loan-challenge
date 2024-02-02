package tests

import (
	"github.com/Israel-Ferreira/loan-challenge/internal/models"
	"github.com/Israel-Ferreira/loan-challenge/internal/services"
	"testing"
)

func assertResult(t *testing.T, result any, expected any) {
	if result != expected {
		t.Error("Falha: Listas de Emprestimos permitidos com diferença de tamanho")
	}
}

func TestListCustomerLoansAllowed(t *testing.T) {
	t.Run("Deve retornar 1 Emprestimo para o cliente com Salário de R$7000, Idade=26 e Residente no estado de SP", func(t *testing.T) {
		customerInfo := models.PersonRequest{
			26,
			"",
			"Vuxaywua Zukiagou",
			7000,
			"SP",
		}

		expectedLoanListSize := 1

		result := services.ListCustomerLoansAllowed(customerInfo)

		assertResult(t, len(result.Loans), expectedLoanListSize)
	})

	t.Run("Deve retornar somente 2 Emprestimos, quando o cliente tiver renda de 3000 e ter 45 anos", func(t *testing.T) {
		expectedSize := 2

		customerInfo := models.PersonRequest{
			45,
			"",
			"Waldemar Aquino Rego",
			3000,
			"RJ",
		}

		result := services.ListCustomerLoansAllowed(customerInfo)

		assertResult(t, len(result.Loans), expectedSize)
	})

	t.Run("Deve retornar somente 1 Emprestimo, quando a renda for maior que 7000 e a idade for igual a 60 anos", func(t *testing.T) {
		expectedSize := 1

		customerInfo := models.PersonRequest{
			60,
			"",
			"Virgulino de Arruda Pinto",
			7000,
			"PE",
		}

		result := services.ListCustomerLoansAllowed(customerInfo)

		assertResult(t, len(result.Loans), expectedSize)
	})

	t.Run("Deve retornar 2 Emprestimos, quando o cliente morar em sp, tiver menos de 30 anos e ganhar 4000", func(t *testing.T) {
		expectedSize := 2

		customerInfo := models.PersonRequest{
			24,
			"",
			"Jacinto Aquino",
			4000,
			"SP",
		}

		result := services.ListCustomerLoansAllowed(customerInfo)

		assertResult(t, len(result.Loans), expectedSize)
	})

	t.Run("Deve retornar 2 Emprestimos, quando a renda do cliente for menor que 3000", func(t *testing.T) {
		expectedSize := 2

		customerInfo := models.PersonRequest{
			24,
			"",
			"Jacinto Aquino",
			2000,
			"SP",
		}

		result := services.ListCustomerLoansAllowed(customerInfo)

		assertResult(t, len(result.Loans), expectedSize)
	})

}
