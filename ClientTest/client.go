package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"grpc_test/ClientTest/auth"
	"grpc_test/ClientTest/tutor"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	username        = "admin1"
	password        = "123456"
	refreshDuration = 30 * time.Second
)

func main() {

	serverAddress := "0.0.0.0:50050"
	enableTLS := true

	fmt.Printf("dial server serverAddress: %s,  TLS: %t  \n", serverAddress, enableTLS)

	transportOpt := grpc.WithInsecure()

	if enableTLS {
		tlsCredentials, err := loadTLSCredentials()
		if err != nil {
			log.Fatal("cannot load TLS credentials: ", err)
		}

		transportOpt = grpc.WithTransportCredentials(tlsCredentials)
	}

	conn1, err := grpc.Dial(serverAddress, transportOpt)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	authClient := auth.NewAuthClient(conn1, username, password)
	intercpetor, err := auth.NewAuthInterceptor(authClient, refreshDuration)
	if err != nil {
		log.Fatal("cannot creat auth intercpetor: ", err)
	}

	conn2, err := grpc.Dial(serverAddress,
		transportOpt,
		grpc.WithUnaryInterceptor(intercpetor.Unary()),
		grpc.WithStreamInterceptor(intercpetor.Stream()))

	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	tutorClient := tutor.NewGameControllerCLient(conn2)

	tutorClient.Unary() // 一元
	// tutorClient.TestServerStream("a") // server stream
	// tutorClient.TestClientStream([4]int32{5, 5, 6, 6})            // client stream
	// tutorClient.BidirectionalStream([]string{"z", "x", "c", "v"}) //bidirectional Stream
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed server's certificate
	pemServerCA, err := ioutil.ReadFile("cert/ca-cert.pem")
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	// Load client's certificate and private key
	clientCert, err := tls.LoadX509KeyPair("cert/client-cert.pem", "cert/client-key.pem")
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
	}

	return credentials.NewTLS(config), nil

}
