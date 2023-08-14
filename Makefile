test:
	go test ./... -v
build:
	go build -o bin/loginApp cmd/loginApp/*.go
run:
	go run cmd/loginApp/main.go