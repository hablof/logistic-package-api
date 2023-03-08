.PHONY: build
build:
	go build cmd/omp-demo-api/main.go

.PHONY: test
test:
	go test -v -race ./internal/app/retranslator

.PHONY: run
run:
	go run ./cmd/logistic-package-api/main.go