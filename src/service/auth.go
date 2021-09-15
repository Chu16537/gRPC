package service

import (
	"context"
	"fmt"
	ClientSDK "grpc_test/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type AuthServer struct {
	userStore  UserStore
	jwtManager *JWTManager
}

func NewAuthServer(userStore UserStore, jwtManager *JWTManager) *AuthServer {
	return &AuthServer{userStore, jwtManager}
}

func (server *AuthServer) Login(ctx context.Context, req *ClientSDK.LoginReq) (*ClientSDK.LoginRes, error) {
	user, err := server.userStore.Find(req.GetUsername())

	if err != nil {
		return nil, status.Errorf(codes.Internal, "connot find user: %v", err)
	}

	if user == nil || !user.IsCorrectPassword(req.GetPassword()) {
		return nil, status.Errorf(codes.NotFound, "incorrect username / password")
	}

	token, err := server.jwtManager.Generate(user)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "connot Generate access token")
	}

	res := &ClientSDK.LoginRes{AccessToken: token}

	return res, nil

}

//攔截器
//auth_interceptor
type AuthInterceptor struct {
	jwtManager      *JWTManager
	accessibleRoles map[string][]string
}

func getAccessibleRoles() map[string][]string {
	const AuthPath = "/ClientSDK.AuthService/"
	const GCPath = "/ClientSDK.GameController/"

	return map[string][]string{
		AuthPath + "login":          {"admin"},
		GCPath + "unary":            {"admin", "user"},
		GCPath + "testStreamServer": {"admin", "user"},
	}
}

func NewAuthInterceptor(jwtManager *JWTManager) *AuthInterceptor {
	return &AuthInterceptor{jwtManager, getAccessibleRoles()}
}

func (interceptor *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		fmt.Println(" unaryServerInterceptor ", info.FullMethod)

		err := interceptor.authorize(ctx, info.FullMethod)

		if err != nil {
			return nil, err
		}

		return handler(ctx, req)
	}
}

func (interceptor *AuthInterceptor) Stream() grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		fmt.Println(" streamServerInterceptor ", info.FullMethod)

		err := interceptor.authorize(stream.Context(), info.FullMethod)

		if err != nil {
			return err
		}

		return handler(srv, stream)
	}
}

func (interceptor *AuthInterceptor) authorize(ctx context.Context, method string) error {
	accessibleRoles, ok := interceptor.accessibleRoles[method]

	if !ok {
		return nil
	}

	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return status.Errorf(codes.Unauthenticated, "authorization token not provided")
	}

	accessToken := values[0]

	claims, err := interceptor.jwtManager.Verify(accessToken)

	if err != nil {
		return status.Errorf(codes.Unauthenticated, "access token is invalid : %v", err)
	}

	for _, role := range accessibleRoles {
		if role == claims.Role {
			fmt.Println("role", role)
			return nil
		}
	}

	return status.Errorf(codes.PermissionDenied, "no PermissionDenied to access this rpc ")
}
