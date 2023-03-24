# USER SERVICE

## Diagram
![alt diagram](https://github.com/yonisaka/protobank/blob/main/diagram.png?raw=true)

## Proto compiler

As we're using Go, the compiler should be the same language
```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

## Installation
Clone repository
```shell
git clone https://github.com/yonisaka/loan-service.git
```
Install

```shell
cd loan-service && go mod download 
```
## Running Application
Run gRPC server
```shell
go run main.go grpc:start
```
Run HTTP Server
```shell
go run main.go
```

## gRPC Server
Any related things to gRPC server could be found in `grpc` directory, including `handler`, `server` and `interceptors`

## gRPC Client
Any related things to gRPC client could be found in `grpc/client` directory, including `connection`, `method call`

## REST API
Any related things to REST API could be found in `rest` directory

## gRPC Interceptor
All interceptors are in `grpc/interceptor` directory