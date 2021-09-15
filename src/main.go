package main

import (
	"flag"
	"fmt"
	grcpManager "grpc_test/src/grpc"
	"log"
	"net"
)

func main() {
	lis, err := CreatListen()
	if err != nil {
		return
	}

	grcpManager.Creat(lis)
}

func CreatListen() (net.Listener, error) {
	// address := "0.0.0.0:50051"

	address := flag.String("address", "0.0.0.0:50051", "the server address")
	flag.Parse()

	fmt.Println("listener address: ", *address)

	lis, err := net.Listen("tcp", *address)
	if err != nil {
		log.Fatalf("failed to listen: %v \n", err)
		return nil, err
	}
	fmt.Println("lis is ok")

	return lis, nil
}
