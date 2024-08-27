FROM golang:1.22.6-alpine

WORKDIR /app

RUN go install github.com/jackc/tern@latest && \
    go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest && \
    go install github.com/air-verse/air@latest && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

COPY go.mod go.sum ./

COPY . .

RUN go mod tidy && \
    go build -o bin/granite ./cmd/granite-identity && \
    chmod +x bin/granite

EXPOSE 8080 50051

CMD ["air", "-c", ".air.toml"]
