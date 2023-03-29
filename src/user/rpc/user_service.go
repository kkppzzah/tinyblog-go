// Package rpc 处理gRPC服务请求。
package rpc

import (
	"context"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"ppzzl.com/tinyblog-go/user/common"
	pb "ppzzl.com/tinyblog-go/user/genproto/user"
	"ppzzl.com/tinyblog-go/user/interfaces"
	"ppzzl.com/tinyblog-go/user/model"
)

// UserService 实现用户服务。
type UserService struct {
	userRepository interfaces.UserRepository
}

// NewUserService 创建UserService实例。
func NewUserService(context interfaces.Context) *UserService {
	rs := &UserService{
		userRepository: context.GetUserRepository(),
	}
	return rs
}

// CreateUser 创建用户。
func CreateUser(name string, password string, userRepository interfaces.UserRepository) (*model.User, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := &model.User{
		Name:     name,
		Nickname: name,
		Password: string(passwordHash),
	}

	user, err = userRepository.Create(user)
	if err != nil {
		log.Printf("failed to creart user, user_name: %s, error: %v", user.Name, err)
		return nil, err
	}
	return user, nil
}

// Create 创建用户。
func (svc *UserService) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	user, err := CreateUser(req.Name, req.Password, svc.userRepository)
	if err != nil {
		return nil, err
	}

	rsp := &pb.CreateResponse{
		Id: user.ID,
	}
	return rsp, nil
}

// Update 更新用户。
func (svc *UserService) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	user := &model.User{
		ID:       req.Id,
		Name:     req.Name,
		Avatar:   req.Avatar,
		Nickname: req.Nickname,
		Bio:      req.Bio,
	}

	err := svc.userRepository.Update(user)
	if err != nil {
		switch v := err.(type) {
		case *common.Error:
			if v.Code == common.ErrorCodeNoFound {
				return nil, status.Error(codes.NotFound, fmt.Sprintf("user %d is not found", req.Id))
			}
			return nil, status.Error(codes.Internal, v.Error())
		default:
			return nil, status.Error(codes.Internal, v.Error())
		}
	}
	rsp := &pb.UpdateResponse{
		Name:     user.Name,
		Avatar:   user.Avatar,
		Nickname: user.Nickname,
		Bio:      user.Bio,
	}
	return rsp, nil
}

// Delete 删除用户。
func (svc *UserService) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.Empty, error) {
	err := svc.userRepository.Delete(req.Id)
	if err != nil {
		return &pb.Empty{}, status.Error(codes.Internal, err.Error())
	}
	return &pb.Empty{}, nil
}

// Get 获取用户信息。
func (svc *UserService) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	user, err := svc.userRepository.Get(req.Id)
	if err != nil {
		log.Printf("failed to get user, user_id: %d", req.Id)
		return nil, err
	}

	if user == nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("user %d is not found", req.Id))
	}

	rsp := &pb.GetResponse{
		Id:       user.ID,
		Name:     user.Name,
		Avatar:   user.Avatar,
		Nickname: user.Nickname,
		Bio:      user.Bio,
	}
	return rsp, nil
}
