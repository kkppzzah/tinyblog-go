GOPATH:=$(shell go env GOPATH)
BUILDDIR=build

ifeq ($(OS),Windows_NT)
	EXE_SUFFIX:=.exe
endif

.PHONY: init
init:
	@go get -u google.golang.org/protobuf/proto
	@go install github.com/golang/protobuf/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

.PHONY: proto
proto:
	@protoc --proto_path=. --go_out=. --go-grpc_out=require_unimplemented_servers=false:. proto/hello.proto

.PHONY: update
update:
	@go get -u

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: build
build:
	@go build -o $(BUILDDIR)/grpc_server$(EXE_SUFFIX) server/server.go
	@go build -o $(BUILDDIR)/grpc_client$(EXE_SUFFIX) client/client.go

.PHONY: test
test:
	@go test -v ./... -cover