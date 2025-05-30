package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/danilotorchio/go-expert-stress-test/internal/app"
)

func main() {
	var (
		url         = flag.String("url", "", "URL do serviço a ser testado")
		requests    = flag.Int("requests", 0, "Número total de requests")
		concurrency = flag.Int("concurrency", 1, "Número de chamadas simultâneas")
	)

	flag.Parse()

	// Validar parâmetros obrigatórios
	if *url == "" {
		fmt.Fprintf(os.Stderr, "Erro: O parâmetro --url é obrigatório\n")
		flag.Usage()
		os.Exit(1)
	}

	if *requests <= 0 {
		fmt.Fprintf(os.Stderr, "Erro: O parâmetro --requests deve ser maior que 0\n")
		flag.Usage()
		os.Exit(1)
	}

	if *concurrency <= 0 {
		fmt.Fprintf(os.Stderr, "Erro: O parâmetro --concurrency deve ser maior que 0\n")
		flag.Usage()
		os.Exit(1)
	}

	// Criar configuração do teste
	config := app.Config{
		URL:         *url,
		Requests:    *requests,
		Concurrency: *concurrency,
	}

	// Executar teste de carga
	tester := app.NewLoadTester(config)

	fmt.Printf("Iniciando teste de carga...\n")
	fmt.Printf("URL: %s\n", config.URL)
	fmt.Printf("Requests: %d\n", config.Requests)
	fmt.Printf("Concorrência: %d\n", config.Concurrency)
	fmt.Println("----------------------------------------")

	result, err := tester.Run()
	if err != nil {
		log.Fatalf("Erro durante execução do teste: %v", err)
	}

	// Exibir relatório
	app.PrintReport(result)
}
