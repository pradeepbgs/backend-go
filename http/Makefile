APP_NAME=go-http-rest-api
CMD_PATH=./cmd/server

.PHONY: run build clean

run:
	go run $(CMD_PATH)/main.go

build:
	go build -o bin/$(APP_NAME) $(CMD_PATH)

clean:
	rm -rf bin
	