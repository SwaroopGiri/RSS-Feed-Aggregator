build:
	@go build -o bin/GoAgg

run: build
	@./bin/GoAgg

test:
	@go test -v ./...