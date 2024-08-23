FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go mod tidy

RUN go build -o myapp ./cmd/myapp

EXPOSE 8080

CMD ["./myapp"]