package app

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/danilotorchio/go-expert-stress-test/internal/models"
)

func TestLoadTester_Run(t *testing.T) {
	// Criar servidor de teste
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}))
	defer server.Close()

	// Configurar teste
	config := models.Config{
		URL:         server.URL,
		Requests:    10,
		Concurrency: 2,
	}

	// Executar teste
	tester := NewLoadTester(config)
	result, err := tester.Run()

	// Verificar resultados
	if err != nil {
		t.Fatalf("Erro inesperado: %v", err)
	}

	if result.TotalRequests != 10 {
		t.Errorf("Esperado 10 requests, obtido %d", result.TotalRequests)
	}

	if result.SuccessfulRequests != 10 {
		t.Errorf("Esperado 10 requests bem-sucedidos, obtido %d", result.SuccessfulRequests)
	}

	if result.StatusCodes[200] != 10 {
		t.Errorf("Esperado 10 status 200, obtido %d", result.StatusCodes[200])
	}

	if result.TotalDuration <= 0 {
		t.Error("Duração total deve ser maior que zero")
	}
}

func TestLoadTester_Run_WithErrors(t *testing.T) {
	// Configurar teste com URL inválida
	config := models.Config{
		URL:         "http://invalid-url-that-does-not-exist.com",
		Requests:    5,
		Concurrency: 2,
	}

	// Executar teste
	tester := NewLoadTester(config)
	result, err := tester.Run()

	// Verificar resultados
	if err != nil {
		t.Fatalf("Erro inesperado: %v", err)
	}

	if result.TotalRequests != 5 {
		t.Errorf("Esperado 5 requests, obtido %d", result.TotalRequests)
	}

	if len(result.Errors) == 0 {
		t.Error("Esperado erros, mas nenhum foi encontrado")
	}

	if result.SuccessfulRequests > 0 {
		t.Errorf("Esperado 0 requests bem-sucedidos, obtido %d", result.SuccessfulRequests)
	}
}

func TestLoadTester_Run_MixedResponses(t *testing.T) {
	requestCount := 0

	// Servidor que alterna entre sucesso e erro
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		if requestCount%2 == 0 {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		w.Write([]byte("Response"))
	}))
	defer server.Close()

	// Configurar teste
	config := models.Config{
		URL:         server.URL,
		Requests:    10,
		Concurrency: 1, // Usar 1 para garantir ordem previsível
	}

	// Executar teste
	tester := NewLoadTester(config)
	result, err := tester.Run()

	// Verificar resultados
	if err != nil {
		t.Fatalf("Erro inesperado: %v", err)
	}

	if result.TotalRequests != 10 {
		t.Errorf("Esperado 10 requests, obtido %d", result.TotalRequests)
	}

	// Deve ter tanto status 200 quanto 404
	if result.StatusCodes[200] == 0 {
		t.Error("Esperado pelo menos um status 200")
	}

	if result.StatusCodes[404] == 0 {
		t.Error("Esperado pelo menos um status 404")
	}

	// Total de status codes deve ser igual ao total de requests
	totalStatusCodes := 0
	for _, count := range result.StatusCodes {
		totalStatusCodes += count
	}

	if totalStatusCodes != result.TotalRequests {
		t.Errorf("Total de status codes (%d) não coincide com total de requests (%d)",
			totalStatusCodes, result.TotalRequests)
	}
}

func TestNewLoadTester(t *testing.T) {
	config := models.Config{
		URL:         "http://example.com",
		Requests:    100,
		Concurrency: 10,
	}

	tester := NewLoadTester(config)

	if tester == nil {
		t.Fatal("LoadTester não deve ser nil")
	}

	if tester.config.URL != config.URL {
		t.Errorf("Esperado URL %s, obtido %s", config.URL, tester.config.URL)
	}

	if tester.config.Requests != config.Requests {
		t.Errorf("Esperado %d requests, obtido %d", config.Requests, tester.config.Requests)
	}

	if tester.config.Concurrency != config.Concurrency {
		t.Errorf("Esperado concorrência %d, obtido %d", config.Concurrency, tester.config.Concurrency)
	}

	if tester.client == nil {
		t.Error("HTTP client não deve ser nil")
	}

	if tester.client.Timeout != 30*time.Second {
		t.Errorf("Esperado timeout de 30s, obtido %v", tester.client.Timeout)
	}
}
