// Package rpc 处理gRPC服务请求。
package rpc

import (
	"context"
	"fmt"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"ppzzl.com/tinyblog-go/user/common"
	pb "ppzzl.com/tinyblog-go/user/genproto/auth"
	"ppzzl.com/tinyblog-go/user/interfaces"
	"ppzzl.com/tinyblog-go/user/model"
)

// AuthService 实现鉴权服务。
type AuthService struct {
	userRepository     interfaces.UserRepository
	sessionRepository  interfaces.SessionRepository
	jwtSecret          []byte
	userEventPublisher interfaces.UserEventPublisher
}

// NewAuthService 创建AuthService实例。
func NewAuthService(context interfaces.Context) *AuthService {
	rs := &AuthService{
		userRepository:     context.GetUserRepository(),
		sessionRepository:  context.GetSessionRepository(),
		jwtSecret:          []byte(common.MustLoadSecretAsString(common.EnvJWTSecret, common.EnvJWTSecretSecretFile)),
		userEventPublisher: context.GetUserEventPublisher(),
	}
	return rs
}

// SignUp 创建用户。
func (svc *AuthService) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	user, err := CreateUser(req.Name, req.Password, svc.userRepository)
	if err != nil {
		return nil, err
	}
	authInfo, err := svc.generateAuthInfo(user, req.AuthMethod)
	if err != nil {
		return nil, err
	}
	rsp := &pb.SignUpResponse{
		Id:       user.ID,
		AuthInfo: authInfo,
	}
	svc.userEventPublisher.Publish(&interfaces.UserEvent{
		EventType: interfaces.UserEventTypeCreate,
		UserInfo: &interfaces.UserInfo{
			ID:       user.ID,
			Name:     user.Name,
			Nickname: user.Nickname,
			Avatar:   user.Avatar,
		},
	})
	return rsp, nil
}

// SignIn 登录。
func (svc *AuthService) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {
	user, err := svc.userRepository.GetByName(req.Name)
	if err != nil {
		switch v := err.(type) {
		case *common.Error:
			if v.Code == common.ErrorCodeNoFound {
				return nil, status.Error(codes.NotFound, fmt.Sprintf("user %s is not found", req.Name))
			}
			return nil, status.Error(codes.Internal, v.Error())
		default:
			return nil, status.Error(codes.Internal, v.Error())
		}
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "")
	}
	authInfo, err := svc.generateAuthInfo(user, req.AuthMethod)
	if err != nil {
		return nil, err
	}
	rsp := &pb.SignInResponse{
		AuthInfo: authInfo,
	}
	return rsp, nil
}

// GetAuthUserInfo 获取认证用户信息。
func (svc *AuthService) GetAuthUserInfo(ctx context.Context, req *pb.GetAuthUserInfoRequest) (*pb.GetAuthUserInfoResponse, error) {
	userID, err := svc.getUserID(req.AuthInfo)
	if err != nil {
		return nil, err
	}
	user, err := svc.userRepository.Get(userID)
	if err != nil {
		return nil, err
	}
	rsp := &pb.GetAuthUserInfoResponse{
		Id:       userID,
		Name:     user.Name,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
		Bio:      user.Bio,
	}
	return rsp, nil
}

// SignOut 退出登录。
func (svc *AuthService) SignOut(ctx context.Context, req *pb.SignOutRequest) (*pb.Empty, error) {
	err := svc.sessionRepository.Delete(req.AuthInfo.Session.Session)
	return &pb.Empty{}, err
}

func (svc *AuthService) getUserID(authInfo *pb.AuthInfo) (int64, error) {
	if authInfo.Session != nil {
		return svc.sessionRepository.Get(authInfo.Session.Session)
	} else if authInfo.Jwt != nil {
		token, err := jwt.ParseWithClaims(authInfo.Jwt.Jwt, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			return svc.jwtSecret, nil
		})

		if err != nil || !token.Valid {
			return 0, err
		}

		if claims, ok := token.Claims.(*jwt.StandardClaims); ok {
			// TODO 此处需要校验token各个claim是否合法。
			userID, _ := strconv.ParseInt(claims.Id, 10, 64)
			return userID, nil
		}
	}
	return 0, nil
}

func (svc *AuthService) generateAuthInfo(user *model.User, authMethod pb.AuthMethod) (*pb.AuthInfo, error) {
	switch authMethod {
	case pb.AuthMethod_AUTH_METHOD_SESSION:
		sessionID, err := svc.generateSession(user)
		if err != nil {
			return nil, err
		}
		return &pb.AuthInfo{Session: &pb.SessionInfo{Session: sessionID}}, nil
	case pb.AuthMethod_AUTH_METHOD_JWT:
		jwtToken, err := svc.generateJWT(user)
		if err != nil {
			return nil, err
		}
		return &pb.AuthInfo{Jwt: &pb.JwtInfo{Jwt: jwtToken}}, nil
	default:
		return nil, common.NewError(common.ErrorCodeAuthMethodNotImplemented, nil)
	}
}

func (svc *AuthService) generateSession(user *model.User) (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	sessionID := fmt.Sprintf("%s%d", id.String(), (53678+user.ID)%10000)
	err = svc.sessionRepository.Create(sessionID, user.ID)
	if err != nil {
		return "", err
	}
	return sessionID, nil
}

func (svc *AuthService) generateJWT(user *model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Id: strconv.Itoa(int(user.ID)),
	})
	jwtToken, err := token.SignedString(svc.jwtSecret)
	return jwtToken, err
}
