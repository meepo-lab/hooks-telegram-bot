SHELL=/bin/bash
BINARY_NAME=hooks-telegram-bot.bin

prerequisites:
	go mod tidy

clean:
	go clean
	rm -f "$(BINARY_NAME)" __debug_bin

build: prerequisites
	go build -o ./build/$(BINARY_NAME) cmd/hooks-telegram-bot/main.go

run: build
	bash -c './build/$(BINARY_NAME)'
