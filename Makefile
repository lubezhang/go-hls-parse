run:
	@go run main.go

build: clean
	@go build -o bin/hls-parse ./main.go
	
clean:
	@rm -rf bin/

test:
	clear
	@go test ./...
	@#go test -v ./tests/...
