FROM golang:1.18-alpine as builder

LABEL maintainer="Grey <soguazu@gmail.com>"

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go install golang.org/x/lint/golint@latest
RUN go install github.com/cespare/reflex@latest

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN rm -rf docs && swag init -g cmd/main.go

RUN CGO_ENABLED=0 GOOS=linux go build -a -o /game cmd/main.go

EXPOSE 8085

CMD ["/game"]

