package models

import (
	"time"
)

// RequestResult representa o resultado de uma requisição HTTP
type RequestResult struct {
	StatusCode int
	Duration   time.Duration
	Error      error
}

// TestResult representa o resultado completo do teste de carga
type TestResult struct {
	TotalRequests      int
	TotalDuration      time.Duration
	SuccessfulRequests int
	StatusCodes        map[int]int
	Errors             []error
}

// Config representa a configuração do teste de carga
type Config struct {
	URL         string
	Requests    int
	Concurrency int
}
