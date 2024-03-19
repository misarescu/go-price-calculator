build:
	@go build -o ./bin/ ./cmd/...

run-calculator:
	@bin/calculator

clean:
	@rm -rf bin