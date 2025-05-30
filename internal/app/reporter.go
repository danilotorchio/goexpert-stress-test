package app

import (
	"fmt"
	"sort"

	"github.com/danilotorchio/go-expert-stress-test/internal/models"
)

// PrintReport exibe o relatório do teste de carga
func PrintReport(result *models.TestResult) {
	fmt.Println("\n========================================")
	fmt.Println("           RELATÓRIO DO TESTE")
	fmt.Println("========================================")

	// Tempo total
	fmt.Printf("Tempo total gasto: %v\n", result.TotalDuration)

	// Quantidade total de requests
	fmt.Printf("Quantidade total de requests: %d\n", result.TotalRequests)

	// Requests com status 200
	fmt.Printf("Requests com status HTTP 200: %d\n", result.SuccessfulRequests)

	// Distribuição de status codes
	fmt.Println("\nDistribuição de códigos de status HTTP:")
	if len(result.StatusCodes) == 0 && len(result.Errors) > 0 {
		fmt.Println("  Nenhum código de status registrado (apenas erros)")
	} else {
		// Ordenar status codes para exibição consistente
		var codes []int
		for code := range result.StatusCodes {
			codes = append(codes, code)
		}
		sort.Ints(codes)

		for _, code := range codes {
			count := result.StatusCodes[code]
			percentage := float64(count) / float64(result.TotalRequests) * 100
			fmt.Printf("  Status %d: %d requests (%.2f%%)\n", code, count, percentage)
		}
	}

	// Erros
	if len(result.Errors) > 0 {
		fmt.Printf("\nErros encontrados: %d\n", len(result.Errors))
		// Mostrar apenas os primeiros 5 erros para não poluir a saída
		maxErrors := 5
		if len(result.Errors) < maxErrors {
			maxErrors = len(result.Errors)
		}

		for i := 0; i < maxErrors; i++ {
			fmt.Printf("  - %v\n", result.Errors[i])
		}

		if len(result.Errors) > 5 {
			fmt.Printf("  ... e mais %d erros\n", len(result.Errors)-5)
		}
	}

	// Taxa de sucesso
	if result.TotalRequests > 0 {
		successRate := float64(result.SuccessfulRequests) / float64(result.TotalRequests) * 100
		fmt.Printf("\nTaxa de sucesso: %.2f%%\n", successRate)
	}

	fmt.Println("========================================")
}
