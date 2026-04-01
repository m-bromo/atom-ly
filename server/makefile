.PHONY: run build test docker-up docker-down start

run:
	@echo "🚀 Rodando a aplicação localmente..."
	go run main.go

build:
	@echo "🔨 Construindo o binário localmente..."
	go build -o atom-ly

test:
	@echo "🧪 Rodando testes..."
	go test ./...

docker-up:
	@echo "🐳 Construindo a imagem Docker..."
	docker-compose up -d

docker-down:
	@echo "🛑 Parando o container..."
	docker-compose down -v

start: docker-up run
	@echo "🚀 Aplicação inicializada"