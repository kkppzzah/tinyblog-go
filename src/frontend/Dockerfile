# Build
FROM golang:1.20.1-alpine AS builder
WORKDIR /src

# 
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk --update --no-cache add ca-certificates make protoc

#
ENV HTTP_PROXY "http://192.168.1.7:10801"
ENV HTTPS_PROXY "http://192.168.1.7:10801"

COPY Makefile go.mod go.sum ./
RUN make init && go mod download
COPY . .
RUN make tidy

# Skaffold passes in debug-oriented compiler flags
ARG SKAFFOLD_GO_GCFLAGS
RUN go build -gcflags="${SKAFFOLD_GO_GCFLAGS}" -o /go/bin/frontend .

# Deployment
FROM alpine:3.17.2
WORKDIR /src

COPY --from=builder /go/bin/frontend /tinyblog-go/frontend
COPY ./templates ./templates
COPY ./static ./static

# Definition of this variable is used by 'skaffold debug' to identify a golang binary.
# Default behavior - a failure prints a stack trace for the current goroutine.
# See https://golang.org/pkg/runtime/
ENV GOTRACEBACK=single

# Set TimeZone
RUN rm -rf /etc/localtime && ln -s /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

# Set TimeZone
RUN rm -rf /etc/localtime && ln -s /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

ENTRYPOINT ["/tinyblog-go/frontend"]

