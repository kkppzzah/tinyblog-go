// Package service 各个服务。
package service

import (
	"context"

	"ppzzl.com/tinyblog-go/frontend/common"
	pb "ppzzl.com/tinyblog-go/frontend/genproto/auth"
)

// AuthService 鉴权服务。
type AuthService struct {
	RPCServiceBase
}

// NewAuthService 创建鉴权服务。
func NewAuthService(mustConnect bool) *AuthService {
	svc := &AuthService{
		RPCServiceBase: RPCServiceBase{
			name:           "auth",
			serviceAddress: common.MustGetEnv(common.EnvAuthServiceAddress, ""),
		},
	}
	svc.initialize(mustConnect)
	return svc
}

// SignUp 用户注册。
func (svc *AuthService) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	return svc.getAuthServiceClient().SignUp(ctx, req)
}

// SignIn 用户登录。
func (svc *AuthService) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {
	return svc.getAuthServiceClient().SignIn(ctx, req)
}

// SignOut 用户取消登录。
func (svc *AuthService) SignOut(ctx context.Context, req *pb.SignOutRequest) (*pb.Empty, error) {
	return svc.getAuthServiceClient().SignOut(ctx, req)
}

// GetAuthUserInfo 获取认证用户信息。
func (svc *AuthService) GetAuthUserInfo(ctx context.Context, req *pb.GetAuthUserInfoRequest) (*pb.GetAuthUserInfoResponse, error) {
	return svc.getAuthServiceClient().GetAuthUserInfo(ctx, req)
}

func (svc *AuthService) getAuthServiceClient() pb.AuthServiceClient {
	return pb.NewAuthServiceClient(svc.GetConnection())
}
