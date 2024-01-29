export CGO_ENABLED=0
run:
	@go run main.go

build:
	@go build -ldflags="-s -w" -o ./bin/sz.exe cmd/sz/main.go

install:
	@go install -ldflags="-s -w" cmd/sz/main.go

build-tinygo:
	@tinygo build -o ./bin/sz.exe cmd/sz/main.go