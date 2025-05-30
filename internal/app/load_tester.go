package app

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/danilotorchio/go-expert-stress-test/internal/models"
)

// LoadTester executa testes de carga
type LoadTester struct {
	config models.Config
	client *http.Client
}

// NewLoadTester cria uma nova instância do load tester
func NewLoadTester(config models.Config) *LoadTester {
	return &LoadTester{
		config: config,
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// Run executa o teste de carga
func (lt *LoadTester) Run() (*models.TestResult, error) {
	start := time.Now()

	// Canal para distribuir trabalho
	jobs := make(chan int, lt.config.Requests)
	results := make(chan models.RequestResult, lt.config.Requests)

	// Inicializar resultado
	result := &models.TestResult{
		TotalRequests: lt.config.Requests,
		StatusCodes:   make(map[int]int),
		Errors:        make([]error, 0),
	}

	// Criar workers
	var wg sync.WaitGroup
	for i := 0; i < lt.config.Concurrency; i++ {
		wg.Add(1)
		go lt.worker(&wg, jobs, results)
	}

	// Enviar trabalhos
	go func() {
		defer close(jobs)
		for i := 0; i < lt.config.Requests; i++ {
			jobs <- i
		}
	}()

	// Aguardar conclusão dos workers
	go func() {
		wg.Wait()
		close(results)
	}()

	// Coletar resultados
	for reqResult := range results {
		if reqResult.Error != nil {
			result.Errors = append(result.Errors, reqResult.Error)
		} else {
			result.StatusCodes[reqResult.StatusCode]++
			if reqResult.StatusCode == http.StatusOK {
				result.SuccessfulRequests++
			}
		}
	}

	result.TotalDuration = time.Since(start)
	return result, nil
}

// worker executa requisições HTTP
func (lt *LoadTester) worker(wg *sync.WaitGroup, jobs <-chan int, results chan<- models.RequestResult) {
	defer wg.Done()

	for range jobs {
		reqResult := lt.makeRequest()
		results <- reqResult
	}
}

// makeRequest faz uma requisição HTTP
func (lt *LoadTester) makeRequest() models.RequestResult {
	start := time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", lt.config.URL, nil)
	if err != nil {
		return models.RequestResult{
			Error:    err,
			Duration: time.Since(start),
		}
	}

	resp, err := lt.client.Do(req)
	if err != nil {
		return models.RequestResult{
			Error:    err,
			Duration: time.Since(start),
		}
	}
	defer resp.Body.Close()

	return models.RequestResult{
		StatusCode: resp.StatusCode,
		Duration:   time.Since(start),
	}
}
