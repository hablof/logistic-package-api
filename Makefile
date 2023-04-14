# GO_VERSION_SHORT:=$(shell echo `go version` | sed -E 's/.* go(.*) .*/\1/g')
# ifneq ("1.16","$(shell printf "$(GO_VERSION_SHORT)\n1.16" | sort -V | head -1)")
# $(error NEED GO VERSION >= 1.16. Found: $(GO_VERSION_SHORT))
# endif

# export GO111MODULE=on

SERVICE_NAME=logistic-package-api
SERVICE_PATH=hablof/logistic-package-api

# PGV_VERSION:="v0.6.1"
# BUF_VERSION:="v0.56.0"

# OS_NAME=$(shell uname -s)
# OS_ARCH=$(shell uname -m)
# GO_BIN=$(shell go env GOPATH)/bin
BUF_EXE=$(GO_BIN)/buf$(shell go env GOEXE)

# ifeq ("NT", "$(findstring NT,$(OS_NAME))")
# OS_NAME=Windows
# endif

.PHONY: run
run:
	go run cmd/grpc-server/main.go

.PHONY: run-r
run-r:
	go run cmd/retranslator/main.go

.PHONY: lint
lint:
	golangci-lint run ./...


# .PHONY: test
# test:
# 	go test -v -race -timeout 30s -coverprofile cover.out ./...
# 	go tool cover -func cover.out | grep total | awk '{print $$3}'


# ----------------------------------------------------------------

# .PHONY: generate-install-buf
# generate-install-buf:
# 	@ command -v buf 2>&1 > /dev/null || (echo "Install buf" && \
#     		curl -sSL0 https://github.com/bufbuild/buf/releases/download/$(BUF_VERSION)/buf-$(OS_NAME)-$(OS_ARCH)$(shell go env GOEXE) --create-dirs -o "$(BUF_EXE)" && \
#     		chmod +x "$(BUF_EXE)")

.PHONY: generate
generate: generate-go generate-kafka-go

.PHONY: generate-go
generate-go:  .generate-go .generate-finalize-go

.generate-go:
	buf generate --template buf.gen.go.yaml

.generate-finalize-go:
	mv pkg/$(SERVICE_NAME)/github.com/$(SERVICE_PATH)/pkg/$(SERVICE_NAME)/* pkg/$(SERVICE_NAME)
	rm -rf pkg/$(SERVICE_NAME)/github.com/
	cd pkg/$(SERVICE_NAME) && ls go.mod || (go mod init github.com/$(SERVICE_PATH)/pkg/$(SERVICE_NAME) && go mod tidy)


.PHONY: generate-kafka-go
generate-kafka-go:  .generate-kafka-go .generate-finalize-kafka-go

.generate-kafka-go:
	buf generate --template buf.gen.kafka.yaml

.generate-finalize-kafka-go:
	mv pkg/kafka-proto/github.com/$(SERVICE_PATH)/pkg/kafka-proto/* pkg/kafka-proto
	rm -rf pkg/kafka-proto/github.com/
	cd pkg/kafka-proto && ls go.mod || (go mod init github.com/$(SERVICE_PATH)/pkg/kafka-proto && go mod tidy)


# ----------------------------------------------------------------

.PHONY: deps-go
deps-go:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.5.0
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.5.0
	go install github.com/envoyproxy/protoc-gen-validate@v0.9.1
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger@latest

.PHONY: build-server
build-server: .build-server

.build-server:
	go build \
		-tags='no_mysql no_sqlite3' \
		-ldflags=" \
			-X 'github.com/$(SERVICE_PATH)/internal/config.version=$(VERSION)' \
			-X 'github.com/$(SERVICE_PATH)/internal/config.commitHash=$(COMMIT_HASH)' \
		" \
		-o ./bin/grpc-server$(shell go env GOEXE) ./cmd/grpc-server/main.go


.PHONY: build-retranslator
build-retranslator: .build-retranslator

.build-retranslator:
	go build \
		-tags='no_mysql no_sqlite3' \
		-ldflags=" \
			-X 'github.com/$(SERVICE_PATH)/internal/config.version=$(VERSION)' \
			-X 'github.com/$(SERVICE_PATH)/internal/config.commitHash=$(COMMIT_HASH)' \
		" \
		-o ./bin/retranslator$(shell go env GOEXE) ./cmd/retranslator/main.go


# ----------------------------------------------------------------

.PHONY: vendor
vendor:
	go mod vendor

.PHONY: docker-s
docker-s:
	docker build -t hablof/logistic-package-api -f ./Dockerfile-server .

.PHONY: docker-run-s
docker-run-s:
	docker run -d \
		-v $(PWD)/config.yml:/root/config.yml \
		hablof/logistic-package-api


.PHONY: docker-r
docker-r:
	docker build -t hablof/logistic-package-api-retranslator -f ./Dockerfile-retranslator .

.PHONY: docker-run-r
docker-run-r:
	docker run -d \
		-v $(PWD)/config.yml:/root/config.yml \
		hablof/logistic-package-api-retranslator

.PHONY: up
up:
	docker-compose up -d

.PHONY: down
down:
	docker-compose down