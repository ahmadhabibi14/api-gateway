FROM golang:1.22.1 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

RUN go build -o service-user .

EXPOSE 3001

FROM ubuntu:20.04
COPY --from=build /app/service-user /usr/local/bin

CMD [ "service-user" ]