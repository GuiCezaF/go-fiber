.PHONY: default run build watch clean
# Variables
APP_NAME=go-fiber

# Tasks
default: run

build:
	@go build -o $(APP_NAME) main.go
run: build
	@./$(APP_NAME)
watch:
	@air
clean:
	@rm -f $(APP_NAME)
	@rm -rf ./docs