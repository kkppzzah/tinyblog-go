// Package service 各个服务。
package service

import (
	"context"

	"ppzzl.com/tinyblog-go/frontend/common"
	pb "ppzzl.com/tinyblog-go/frontend/genproto/user"
)

// UserService 用户服务。
type UserService struct {
	RPCServiceBase
}

// NewUserService 创建用户服务。。
func NewUserService(mustConnect bool) *UserService {
	svc := &UserService{
		RPCServiceBase: RPCServiceBase{
			name:           "user",
			serviceAddress: common.MustGetEnv(common.EnvUserServiceAddress, ""),
		},
	}
	svc.initialize(mustConnect)
	return svc
}

// Create 创建用户。
func (svc *UserService) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	return svc.getServiceClient().Create(ctx, req)
}

// Get 获取用户信息。
func (svc *UserService) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	return svc.getServiceClient().Get(ctx, req)
}

// Update 更新用户。
func (svc *UserService) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	return svc.getServiceClient().Update(ctx, req)
}

// Delete 删除用户。
func (svc *UserService) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.Empty, error) {
	return svc.getServiceClient().Delete(ctx, req)
}

func (svc *UserService) getServiceClient() pb.UserServiceClient {
	return pb.NewUserServiceClient(svc.GetConnection())
}
