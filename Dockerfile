# Build stage
FROM golang:1.24.3-alpine AS builder

# Instalar certificados SSL e git
RUN apk add --no-cache ca-certificates git

# Configurar diretório de trabalho
WORKDIR /app

# Copiar go mod e sum files
COPY go.mod go.sum ./

# Baixar dependências
RUN go mod download

# Copiar código fonte
COPY . .

# Build da aplicação
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o stress-test ./cmd/stress-test

# Runtime stage
FROM alpine:latest

# Instalar certificados SSL
RUN apk --no-cache add ca-certificates

# Criar usuário não-root
RUN adduser -D -s /bin/sh appuser

# Configurar diretório de trabalho
WORKDIR /app

# Copiar binário do stage anterior
COPY --from=builder /app/stress-test .

# Alterar ownership para o usuário não-root
RUN chown appuser:appuser /app/stress-test

# Mudar para usuário não-root
USER appuser

# Comando padrão
ENTRYPOINT ["./stress-test"] 