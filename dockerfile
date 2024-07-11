# Sử dụng image Go làm base image
FROM golang:1.21-alpine

WORKDIR /app
COPY . .
RUN go mod download

# Thiết lập biến môi trường
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 

# Build ứng dụng Go trong thư mục cmd/dev
WORKDIR /app/cmd/dev
RUN go build -o main .

# Quay lại thư mục gốc của ứng dụng
WORKDIR /app

# Command để chạy ứng dụng Go từ thư mục cmd/dev
CMD ["/app/cmd/dev/main"]

