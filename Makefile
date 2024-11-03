.PHONY: build
build: 
	go mod tidy
	go build -o ./cmd/bin/space-trouble ./cmd/main.go

.PHONY: clean
clean:	
	rm -f ./cmd/bin/space-trouble

.PHONY: test
test:
	go test ./...
