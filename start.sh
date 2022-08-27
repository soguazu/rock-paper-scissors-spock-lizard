#!/bin/sh
awhile=3
go install github.com/swaggo/swag/cmd/swag@latest
go mod download golang.org/x/lint
go install golang.org/x/lint/golint@latest
go install github.com/cespare/reflex@latest
go mod tidy
rm -rf docs && swag init -g cmd/main.go
sleep $awhile && open http://localhost:8085/swagger/index.html &
make