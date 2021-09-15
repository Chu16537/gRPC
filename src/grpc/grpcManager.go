package grpc

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	ClientSDK "grpc_test/proto"
	"grpc_test/src/service"
	"io/ioutil"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

type Manager interface {
	Creat()
}

const (
	serverCertFile   = "cert/server-cert.pem"
	serverKeyFile    = "cert/server-key.pem"
	clientCACertFile = "cert/ca-cert.pem"
)

func Creat(lis net.Listener) {
	enableTLS := true

	fmt.Println("grpc server creat start TLS: ", enableTLS)

	userStore := service.NewInMemoryUserStore()

	err := service.SeedUsers(userStore)
	if err != nil {
		fmt.Println("connot seed users")
	}

	jwtManager := service.NewJWTManager()
	authServer := service.NewAuthServer(userStore, jwtManager)

	gamecontroller := service.NewGameController()

	interceptor := service.NewAuthInterceptor(jwtManager)

	serverOptions := []grpc.ServerOption{
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	}

	if enableTLS {
		tlsCredentalis, err := loadTLSCredentials()
		if err != nil {
			log.Fatal("connot load TLS credentails:  ", err)
		}
		serverOptions = append(serverOptions, grpc.Creds(tlsCredentalis))
	}

	grpcServer := grpc.NewServer(serverOptions...)

	ClientSDK.RegisterAuthServiceServer(grpcServer, authServer)
	ClientSDK.RegisterGameControllerServer(grpcServer, gamecontroller)

	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v \n", err)
	}

	fmt.Println("grpc server creat end")

}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed client's certificate
	pemClientCA, err := ioutil.ReadFile(clientCACertFile)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemClientCA) {
		return nil, fmt.Errorf("failed to add client CA's certificate")
	}

	// Load server's certificate and private key
	serverCert, err := tls.LoadX509KeyPair(serverCertFile, serverKeyFile)
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	}

	return credentials.NewTLS(config), nil
}
