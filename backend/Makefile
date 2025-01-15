# Создать и запустить контейнеры со сборкой 
DEFAULT_CONFIG_PATH = ./configs/default_config.yaml

up:
	docker compose up -d

down:
	docker compose down

restart:
	docker compose restart

restart-go:
	docker compose restart golang

dev:
	go run ./cmd/esoft/main.go \
		--config_path=$(DEFAULT_CONFIG_PATH)

# Применить миграции к базе данных
migrate: start
	docker run --rm --network gateway_database -v ./:/code -w /code golang:latest go run ./cmd/migrator/migrator.go 

# Сгенерировать описание API по комментариям в коде
swagger:
	docker run --rm -v ./:/code ghcr.io/swaggo/swag:latest init -g ./cmd/main/main.go -o ./api
