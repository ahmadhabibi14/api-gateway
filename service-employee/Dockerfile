FROM golang:1.22.1 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

RUN go build -o service-employee .

EXPOSE 3002

FROM ubuntu:20.04
COPY --from=build /app/service-employee /usr/local/bin

CMD [ "service-employee" ]