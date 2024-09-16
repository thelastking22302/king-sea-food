# Sử dụng image Go làm base image
FROM golang:alpine AS builder 

WORKDIR /build
COPY . .
RUN go mod download 

RUN go build -o kingseafood ./cmd/dev/main.go

FROM  scratch

COPY  --from=builder /build/kingseafood /
ENTRYPOINT [ "/kingseafood" ]

