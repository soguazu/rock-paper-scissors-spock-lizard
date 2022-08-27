FROM golang:latest

LABEL maintainer="Grey <soguazu@gmail.com>"

WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go mod download golang.org/x/lint
RUN go install golang.org/x/lint/golint@latest

RUN rm -rf docs && swag init -g cmd/main.go
RUN go mod download

COPY . .

RUN go build -o server cmd/main.go

CMD ["./server"]


