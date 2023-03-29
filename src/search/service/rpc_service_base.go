// Package service 各个服务。
package service

import (
	"log"

	"google.golang.org/grpc"
)

// RPCServiceBase 是各个RPC服务的公共功能。
type RPCServiceBase struct {
	name           string
	serviceAddress string
	connection     *grpc.ClientConn
}

// Initialize 初始化服务，建立到RPC服务的连接。
func (svc *RPCServiceBase) initialize(mustConnect bool) {
	connection, err := grpc.Dial(svc.serviceAddress, grpc.WithInsecure())
	if err != nil {
		if mustConnect {
			log.Fatalf("failed connect service %s@'%s', %v", svc.name, svc.serviceAddress, err)
		}
		log.Printf("failed connect service '%s', %v", svc.name, err)
	}
	svc.connection = connection
}

// GetName 返回服务名。
func (svc *RPCServiceBase) GetName() string {
	return svc.name
}

// GetServiceAddress 返回服务地址。
func (svc *RPCServiceBase) GetServiceAddress() string {
	return svc.serviceAddress
}

// GetConnection 返回到服务的连接。
func (svc *RPCServiceBase) GetConnection() *grpc.ClientConn {
	return svc.connection
}
