build:
	go build -o bin/go-utils ./cmd

test:
	go test -v ./...

benchmark:
	go test -bench ./...

lint:
	golangci-lint run

clean:
	rm -rf ./bin
