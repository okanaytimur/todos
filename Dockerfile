# Stage 1: Build
FROM golang:1.25 AS builder
WORKDIR /app
COPY . .
RUN go mod init todoapp && go mod tidy && go build -o todo-server main.go

# Stage 2: Runtime
FROM debian:bookworm-slim
WORKDIR /app
COPY --from=builder /app/todo-server .
COPY ./data ./data
EXPOSE 8080
CMD ["./todo-server"]
