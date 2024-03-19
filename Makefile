BUILD_DIR= bin

build:
	@go build -o ./$(BUILD_DIR)/ ./cmd/...

clean-bin:
	@rm -rf bin

clean-data:
	@rm data/result-*.json

clean: clean-bin clean-data
