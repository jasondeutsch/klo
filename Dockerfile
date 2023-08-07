FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY . /app

#RUN apk update && apk add alpine-sdk

#RUN go mod download

RUN go build -o /flightplanner ./...

ENTRYPOINT /flightplanner
