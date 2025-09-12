FROM golang:1.21-alpine

WORKDIR /app

# Копируем только go.mod сначала
COPY go.mod ./

# Запускаем приложение
CMD ["go run main.go"]