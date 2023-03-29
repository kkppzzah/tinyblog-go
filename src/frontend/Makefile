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
	@go install github.com/go-micro/generator/cmd/protoc-gen-micro@latest

.PHONY: proto
proto:
	@protoc --proto_path=$(PROTODIR) --go_out=. --go-grpc_out=require_unimplemented_servers=false:./ $(PROTODIR)/user.proto
	@protoc --proto_path=$(PROTODIR) --go_out=. --go-grpc_out=require_unimplemented_servers=false:./ $(PROTODIR)/article.proto
	@protoc --proto_path=$(PROTODIR) --go_out=. --go-grpc_out=require_unimplemented_servers=false:./ $(PROTODIR)/recommend.proto
	@protoc --proto_path=$(PROTODIR) --go_out=. --go-grpc_out=require_unimplemented_servers=false:./ $(PROTODIR)/storage.proto
	@protoc --proto_path=$(PROTODIR) --go_out=. --go-grpc_out=require_unimplemented_servers=false:./ $(PROTODIR)/search.proto
	@protoc --proto_path=$(PROTODIR) --go_out=. --go-grpc_out=require_unimplemented_servers=false:./ $(PROTODIR)/auth.proto

.PHONY: update
update:
	@go get -u

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: build
build:
	@go build -o frontend$(EXE_SUFFIX) *.go

.PHONY: test
test:
	@go test -v ./... -cover

.PHONY: gencss
gencss:
	@sass -I ./static/src/css --no-source-map ./webfrontend/css/home.sass:./static/css/home.css
	@sass -I ./static/src/css --no-source-map ./webfrontend/css/account_signin.sass:./static/css/account_signin.css
	@sass -I ./static/src/css --no-source-map ./webfrontend/css/account_signup.sass:./static/css/account_signup.css
	@sass -I ./static/src/css --no-source-map ./webfrontend/css/user_article_post.sass:./static/css/user_article_post.css
	@sass -I ./static/src/css --no-source-map ./webfrontend/css/article.sass:./static/css/article.css
	@sass -I ./static/src/css --no-source-map ./webfrontend/css/article_edit.sass:./static/css/article_edit.css
	@sass -I ./static/src/css --no-source-map ./webfrontend/css/my_article_list.sass:./static/css/my_article_list.css
	@sass -I ./static/src/css --no-source-map ./webfrontend/css/search.sass:./static/css/search.css

.PHONY: docker
docker:
	@docker build -t frontend:latest .