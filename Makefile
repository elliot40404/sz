export CGO_ENABLED=0

sz:
	go build -ldflags="-s -w" -o ./bin/sz.exe ./cmd/sz/

sz-tiny:
	tinygo build -o ./bin/sz.exe ./cmd/sz/

sz-install:
	go install ./cmd/sz/
	