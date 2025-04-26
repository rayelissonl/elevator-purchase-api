# Variáveis
APP_NAME=elevator-purchase-api
MAIN=cmd/server/main.go
SWAGGER_DIR=cmd/server/main.go

# Comandos
.PHONY: build run swag test migrate

# Builda a aplicação
build:
	go build -o bin/$(APP_NAME) $(MAIN)

# Roda a aplicação
run:
	go run $(MAIN)

# Gera documentação Swagger
swag:
	swag init -g $(SWAGGER_DIR)

# Roda os testes
test:
	go test ./...

# Executa as migrações do GORM
migrate:
	go run internal/infrastructure/migrations/main.go

# Formata o código (opcional)
fmt:
	go fmt ./...

# Verifica se o código tá certinho
vet:
	go vet ./...
