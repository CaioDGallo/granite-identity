FROM golang:1.22.6-alpine

WORKDIR /app
# Install tern
RUN go install github.com/jackc/tern@latest

COPY go.mod go.sum ./

COPY . .

RUN go mod tidy && \
    go build -o granite ./cmd/granite-identity && \
    chmod +x granite

EXPOSE 8080

CMD [ "./granite" ]

