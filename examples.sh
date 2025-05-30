#!/bin/bash

echo "=========================================="
echo "    Exemplos de Uso do Go Stress Test"
echo "=========================================="

echo
echo "1. Teste básico com 50 requests e 5 conexões simultâneas:"
echo "   ./bin/stress-test --url=http://httpbin.org/status/200 --requests=50 --concurrency=5"
echo

echo "2. Teste de alta concorrência:"
echo "   ./bin/stress-test --url=http://httpbin.org/delay/1 --requests=100 --concurrency=20"
echo

echo "3. Teste com diferentes status codes (404):"
echo "   ./bin/stress-test --url=http://httpbin.org/status/404 --requests=30 --concurrency=3"
echo

echo "4. Teste com erro de conexão:"
echo "   ./bin/stress-test --url=http://site-inexistente.com --requests=10 --concurrency=2"
echo

echo "5. Via Docker (quando disponível):"
echo "   docker run --rm stress-test --url=http://google.com --requests=100 --concurrency=10"
echo

echo "=========================================="
echo "Para executar um exemplo, copie e cole o comando desejado"
echo "==========================================" 