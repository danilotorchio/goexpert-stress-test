.PHONY: build run clean docker-build docker-run test

# Nome da aplicação
APP_NAME=stress-test
DOCKER_IMAGE=stress-test

# Build da aplicação
build:
	go build -o bin/$(APP_NAME) ./cmd/stress-test

# Executar aplicação localmente
run: build
	./bin/$(APP_NAME)

# Limpar arquivos gerados
clean:
	rm -rf bin/
	docker rmi $(DOCKER_IMAGE) 2>/dev/null || true

# Build da imagem Docker
docker-build:
	docker build -t $(DOCKER_IMAGE) .

# Executar via Docker
docker-run: docker-build
	docker run --rm $(DOCKER_IMAGE) $(ARGS)

# Executar testes
test:
	go test -v ./...

# Formatar código
fmt:
	go fmt ./...

# Verificar módulos
mod-tidy:
	go mod tidy

# Exemplo de uso
example:
	@echo "Exemplos de uso:"
	@echo "  make docker-run ARGS='--url=http://google.com --requests=100 --concurrency=10'"
	@echo "  make run ARGS='--url=http://localhost:8080 --requests=1000 --concurrency=50'" 