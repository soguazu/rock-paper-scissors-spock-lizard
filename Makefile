.DEFAULT_GOAL := dev

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	golint ./...
.PHONY:lint

vet: lint
	go vet ./...
	# shadow ./...
.PHONY:vet

dev: vet
	@reflex -r '.go' -s -- go run cmd/main.go

swagger:
	 ./start.sh
.PHONY:swagger


