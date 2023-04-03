GOPATH:=$(shell go env GOPATH)
PROTODIR:=../../pb
BUILDDIR=build

ifeq ($(OS),Windows_NT)
	EXE_SUFFIX:=.exe
endif

.PHONY: init
init:
	@go get -u google.golang.org/protobuf/proto
	@go install github.com/golang/protobuf/protoc-gen-go@latest

.PHONY: proto
proto:
	@protoc --proto_path=$(PROTODIR) --go_out=. --go-grpc_out=require_unimplemented_servers=false:./ $(PROTODIR)/user.proto
	@protoc --proto_path=$(PROTODIR) --go_out=. --go-grpc_out=require_unimplemented_servers=false:./ $(PROTODIR)/article.proto
	@protoc --proto_path=$(PROTODIR) --go_out=. --go-grpc_out=require_unimplemented_servers=false:./ $(PROTODIR)/recommend.proto
	@protoc --proto_path=$(PROTODIR) --go_out=. --go-grpc_out=require_unimplemented_servers=false:./ $(PROTODIR)/storage.proto
	@protoc --proto_path=$(PROTODIR) --go_out=. --go-grpc_out=require_unimplemented_servers=false:./ $(PROTODIR)/search.proto
	@protoc --proto_path=$(PROTODIR) --go_out=. --go-grpc_out=require_unimplemented_servers=false:./ $(PROTODIR)/health.proto

.PHONY: update
update:
	@go get -u

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: build
build:
	@go build -o recommend$(EXE_SUFFIX) *.go

.PHONY: test
test:
	@go test -v ./... -cover

.PHONY: docker
docker:
	@docker build -t tinyblog-go.recommend:latest .