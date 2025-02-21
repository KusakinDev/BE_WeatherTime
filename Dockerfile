FROM golang:1.23-alpine

WORKDIR /app

# Установить зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копирование исходного кода в контейнер
COPY . .

# Экспонирование порта
EXPOSE 8080

# Команда для запуска в режиме разработки  docker run -v ${PWD}:/app -w /app -p 8080:8080 --rm bewtdev go run api.go
CMD ["go", "run", "api.go"]
