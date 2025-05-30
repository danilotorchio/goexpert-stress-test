# Go Stress Test

Um sistema CLI desenvolvido em Go para realizar testes de carga (stress test) em serviços web. A aplicação permite testar a performance e capacidade de resposta de endpoints HTTP através de requisições simultâneas.

## 🚀 Funcionalidades

- **Testes de Carga Configuráveis**: Defina o número total de requests e nível de concorrência
- **Relatórios Detalhados**: Análise completa dos resultados incluindo tempo de resposta, status codes e taxa de sucesso
- **Execução via Docker**: Aplicação totalmente containerizada para fácil deploy e execução
- **Arquitetura Robusta**: Implementação seguindo boas práticas do Go com separação de responsabilidades
- **Testes Automatizados**: Cobertura de testes para garantir qualidade do código

## 📁 Estrutura do Projeto

```
.
├── cmd/
│   └── stress-test/
│       └── main.go              # Entry point da aplicação
├── internal/
│   ├── app/
│   │   ├── config.go           # Configuração da aplicação
│   │   ├── load_tester.go      # Lógica principal do teste de carga
│   │   ├── load_tester_test.go # Testes unitários
│   │   └── reporter.go         # Geração de relatórios
│   └── models/
│       └── models.go           # Estruturas de dados
├── Dockerfile                  # Configuração do container
├── Makefile                   # Automação de tarefas
├── go.mod                     # Dependências do Go
├── go.sum                     # Checksums das dependências
└── README.md                  # Documentação
```

## 🛠️ Instalação e Execução

### Pré-requisitos

- Go 1.24.3 ou superior
- Docker (para execução containerizada)

### Executando com Docker (Recomendado)

1. **Build da imagem Docker:**
```bash
make docker-build
```

2. **Executar teste de carga:**
```bash
docker run --rm stress-test --url=http://google.com --requests=1000 --concurrency=10
```

### Executando Localmente

1. **Build da aplicação:**
```bash
make build
```

2. **Executar teste:**
```bash
./bin/stress-test --url=http://google.com --requests=1000 --concurrency=10
```

## 📋 Parâmetros

| Parâmetro | Descrição | Obrigatório | Exemplo |
|-----------|-----------|-------------|---------|
| `--url` | URL do serviço a ser testado | ✅ | `--url=http://google.com` |
| `--requests` | Número total de requests | ✅ | `--requests=1000` |
| `--concurrency` | Número de chamadas simultâneas | ✅ | `--concurrency=10` |

## 📊 Exemplo de Relatório

```
========================================
           RELATÓRIO DO TESTE
========================================
Tempo total gasto: 2.345s
Quantidade total de requests: 1000
Requests com status HTTP 200: 980

Distribuição de códigos de status HTTP:
  Status 200: 980 requests (98.00%)
  Status 404: 15 requests (1.50%)
  Status 500: 5 requests (0.50%)

Taxa de sucesso: 98.00%
========================================
```

## 🧪 Testes

Execute os testes unitários:

```bash
make test
```

## 🔧 Comandos Úteis

```bash
# Build da aplicação
make build

# Executar localmente
make run

# Build da imagem Docker
make docker-build

# Executar via Docker com parâmetros
make docker-run ARGS='--url=http://localhost:8080 --requests=100 --concurrency=5'

# Executar testes
make test

# Limpar arquivos gerados
make clean

# Formatar código
make fmt

# Atualizar dependências
make mod-tidy

# Ver exemplos de uso
make example
```

## 🏗️ Arquitetura

### Design Patterns Utilizados

- **Worker Pool Pattern**: Para gerenciar concorrência de forma eficiente
- **Command Pattern**: CLI estruturado com validação de parâmetros
- **Strategy Pattern**: Separação entre coleta de dados e geração de relatórios

### Componentes Principais

1. **LoadTester**: Core da aplicação responsável por coordenar os testes
2. **Worker Pool**: Gerencia a execução concorrente de requisições HTTP
3. **Reporter**: Formatação e exibição dos resultados
4. **Models**: Estruturas de dados compartilhadas

### Características de Performance

- **Goroutines**: Utilização de goroutines para execução paralela
- **Channels**: Comunicação segura entre workers
- **Context**: Controle de timeout e cancelamento
- **HTTP Client Reutilização**: Pool de conexões para melhor performance

## 📈 Exemplo de Uso Prático

### Teste Básico
```bash
docker run --rm stress-test \
  --url=http://httpbin.org/status/200 \
  --requests=100 \
  --concurrency=10
```

### Teste de Alta Concorrência
```bash
docker run --rm stress-test \
  --url=http://httpbin.org/delay/1 \
  --requests=1000 \
  --concurrency=100
```

### Teste de Endpoints com Diferentes Respostas
```bash
docker run --rm stress-test \
  --url=http://httpbin.org/status/200,404,500 \
  --requests=500 \
  --concurrency=25
```

## 🐳 Docker

### Build Multi-stage

O Dockerfile utiliza build multi-stage para otimizar o tamanho da imagem final:

- **Stage 1 (builder)**: Compila a aplicação Go
- **Stage 2 (runtime)**: Imagem Alpine mínima com apenas o binário

### Características de Segurança

- Execução com usuário não-root
- Imagem final baseada em Alpine Linux
- Certificados SSL incluídos para requisições HTTPS

## 📝 Licença

Este projeto está sob a licença MIT. Veja o arquivo `LICENSE` para mais detalhes.

## 👨‍💻 Autor

Desenvolvido por [Danilo Torchio](https://github.com/danilotorchio) como parte do projeto Go Expert.

---

**Nota**: Esta aplicação foi desenvolvida para fins educacionais e de demonstração. Para uso em produção, considere implementar features adicionais como logging estruturado, métricas mais detalhadas e configurações avançadas de rede. 