FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o service-user .

CMD ["./service-user"]

EXPOSE 3001