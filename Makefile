.PHONY: build
build: 
	go mod tidy
	go build -o ./cmd/bin/space-trouble ./cmd/main.go

.PHONY: clean
clean:	
	rm -f ./cmd/bin/space-trouble

.PHONY: generate
test:
	go generate ./...

.PHONY: test
test:
	go test ./...
