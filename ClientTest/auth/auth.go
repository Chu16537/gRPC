package auth

import (
	"context"
	"fmt"
	ClientSDK "grpc_test/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const TIME_OUT = 10 * time.Second

type AuthClient struct {
	service  ClientSDK.AuthServiceClient
	username string
	password string
}

func NewAuthClient(cc *grpc.ClientConn, username string, password string) *AuthClient {
	service := ClientSDK.NewAuthServiceClient(cc)
	return &AuthClient{service, username, password}
}

func (client *AuthClient) Login() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &ClientSDK.LoginReq{
		Username: client.username,
		Password: client.password,
	}

	res, err := client.service.Login(ctx, req)
	if err != nil {
		fmt.Println("Login err :", err)
	}

	return res.GetAccessToken(), nil

}

//AuthInterceptor
type AuthInterceptor struct {
	authClient  *AuthClient
	authMethods map[string]bool
	accessToken string
}

func getAuthMethods() map[string]bool {
	const AuthPath = "/ClientSDK.AuthService/"
	const GCPath = "/ClientSDK.GameController/"

	return map[string]bool{
		AuthPath + "login":          true,
		GCPath + "unary":            true,
		GCPath + "testStreamServer": true,
	}
}

func NewAuthInterceptor(authClient *AuthClient, refreshDuration time.Duration) (*AuthInterceptor, error) {

	interceptor := &AuthInterceptor{
		authClient:  authClient,
		authMethods: getAuthMethods(),
	}

	err := interceptor.scheduleRefreshToken(refreshDuration)

	if err != nil {
		return nil, err
	}

	return interceptor, nil
}

func (interceptor *AuthInterceptor) Unary() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		fmt.Printf("UnaryClientInterceptor :%s", method)

		if interceptor.authMethods[method] {
			return invoker(interceptor.attachToken(ctx), method, req, reply, cc, opts...)
		}

		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func (interceptor *AuthInterceptor) Stream() grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		fmt.Printf("StreamClientInterceptor :%s", method)

		if interceptor.authMethods[method] {
			return streamer(interceptor.attachToken(ctx), desc, cc, method, opts...)
		}

		return streamer(ctx, desc, cc, method, opts...)
	}
}

func (interceptor *AuthInterceptor) attachToken(ctx context.Context) context.Context {
	return metadata.AppendToOutgoingContext(ctx, "authorization", interceptor.accessToken)
}

func (interceptor *AuthInterceptor) scheduleRefreshToken(refreshDuration time.Duration) error {
	err := interceptor.refreshToken()

	if err != nil {
		return err
	}

	// 自動更新token
	go func() {
		wait := refreshDuration
		for {
			time.Sleep(wait)
			err := interceptor.refreshToken()
			if err != nil {
				wait = time.Second
			} else {
				wait = refreshDuration
			}
		}
	}()

	return nil
}

func (interceptor *AuthInterceptor) refreshToken() error {
	accessToken, err := interceptor.authClient.Login()
	if err != nil {
		return err
	}

	interceptor.accessToken = accessToken
	fmt.Printf("token refreshToken : %v \n", accessToken)
	return nil
}
