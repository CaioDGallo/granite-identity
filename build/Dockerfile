FROM golang:1.22.6-alpine

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN go mod tidy && \
    RUN go install github.com/jackc/tern@latest && \
    go build -o granite ./cmd/granite-identity && \
    chmod +x granite

EXPOSE 8080

CMD [ "./granite" ]

