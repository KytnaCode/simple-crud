all:
	go run ./cmd/app

dev:
	air

build:
	go build -o ./out/ ./cmd/...
