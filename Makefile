build:
	go build -o ./bin/Blockchain

run: build
	./bin/Blockchain

test:
	go test -v ./...