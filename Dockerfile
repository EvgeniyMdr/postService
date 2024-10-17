# Используем официальный образ Go для сборки
FROM golang:1.22-alpine AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /build

# Копируем файлы модулей и устанавливаем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем все исходные файлы проекта в рабочую директорию
COPY . .

# Собираем бинарный файл
RUN go build -o postservice ./cmd/postservice/main.go

# Используем минимальный образ для запуска
FROM alpine:3.18

# Устанавливаем рабочую директорию для финального образа
WORKDIR /app

# Копируем скомпилированный бинарник из builder
COPY --from=builder /build/postservice .

# Копируем SQL файлы в финальный образ
COPY --from=builder /build/db/sql /app/db/sql

# Экспонируем порты для HTTP сервера
EXPOSE 8080

# Запускаем приложение
CMD ["./postservice"]
