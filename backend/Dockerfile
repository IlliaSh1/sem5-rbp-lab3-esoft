# Используем официальный образ Go версии
FROM golang:1.23.4-alpine

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

# Загружаем зависимости
# RUN go mod tidy

# Копируем файлы в контейнер
COPY . .

# Собираем приложение
# RUN go build -o ./cmd/main/main.out ./cmd/main/main.go
# RUN go build -o ./cmd/fixture/user_data.out ./cmd/fixture/user_data.go
# RUN go build -o ./cmd/migrator/migrator.out ./cmd/migrator/migrator.go

# RUN ./cmd/migrator/migrator.out
# RUN ./cmd/fixture/user_data.out

# Указываем команду, которая будет выполняться при запуске контейнера
CMD ["go", "run", "./cmd/main/main.go"]
