.PHONY: build run test clean docker-up docker-down docker-restart docker-erase docker-log

# Variáveis
APP_NAME=farm-api
DOCKER_IMAGE=backend-farms-api:latest

# Comandos Go
build:
	go build -o $(APP_NAME) ./cmd/api

run:
	go run ./cmd/api

test:
	go test -v ./...

clean:
	rm -f $(APP_NAME)

# Comandos Docker
docker-build:
	docker build -t $(DOCKER_IMAGE) .

docker-up: docker-build
	docker-compose up -d

docker-down:
	docker-compose down

docker-restart: docker-down docker-up

docker-erase: docker-down
	docker rmi $(DOCKER_IMAGE)

docker-log:
	docker-compose logs -f

# Instalação de dependências
install:
	go mod tidy
