export CGO_ENABLED=0
run:
	@go run main.go

build:
	@go build -ldflags="-s -w" -o ./bin/sz.exe main.go