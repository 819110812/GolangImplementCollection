build:
	go build cmd/main.go

run:
	go mod tidy
	go run cmd/main.go