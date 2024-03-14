FROM golang:1.22.1 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

RUN go build -o api-gateway .

EXPOSE 3000

FROM ubuntu:20.04
COPY --from=build /app/api-gateway /usr/local/bin

CMD [ "api-gateway" ]