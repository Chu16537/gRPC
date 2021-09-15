package tutor

import (
	"context"
	"fmt"
	ClientSDK "grpc_test/proto"
	"io"
	"time"

	"google.golang.org/grpc"
)

// unary() // 一元
// testServerStream("a") // server stream
// testClientStream([4]int32{5, 5, 6, 6}) // client stream
// bidirectionalStream([]string{"z", "x", "c", "v"}) //bidirectional Stream

const TIME_OUT = 10 * time.Second

type GameControllerClient struct {
	service ClientSDK.GameControllerClient
}

func NewGameControllerCLient(cc *grpc.ClientConn) *GameControllerClient {
	service := ClientSDK.NewGameControllerClient(cc)
	return &GameControllerClient{service}
}

func (gameControllerClient *GameControllerClient) Unary() {

	ctx, cancel := context.WithTimeout(context.Background(), TIME_OUT)
	defer cancel()

	req := &ClientSDK.UnaryReq{
		UnaryReq: "5",
	}

	res, err := gameControllerClient.service.Unary(ctx, req)
	if err != nil {
		fmt.Println(" err :", err)
	}

	fmt.Println("response : ", res)
}

func (gameControllerClient *GameControllerClient) TestServerStream(data string) {

	ctx, cancel := context.WithTimeout(context.Background(), TIME_OUT)
	defer cancel()

	req := &ClientSDK.ServerStreamReq{
		ServerStreamReq: data,
	}

	stream, err := gameControllerClient.service.TestStreamServer(ctx, req)
	if err != nil {
		fmt.Println("testStream stream err:", err)
	}

	for {
		res, err := stream.Recv()

		// 表示 server 的 stream 資料傳完了
		if err == io.EOF {
			return
		}

		// 表示有錯誤發生
		if err != nil {
			fmt.Println("testStream stream.Recv err:", err)
			return
		}

		// 讀取 server stream 傳來的資料
		a := res.GetServerStreamRes()
		fmt.Println("a:", a)
	}
}

func (gameControllerClient *GameControllerClient) TestClientStream(data [4]int32) {

	ctx, cancel := context.WithTimeout(context.Background(), TIME_OUT)
	defer cancel()

	stream, err := gameControllerClient.service.TestStreamClient(ctx)
	if err != nil {
		fmt.Println("testClientStream client.TestStreamClient err:", err)
		return
	}

	for _, num := range data {

		req := &ClientSDK.ClientStreamReq{
			ClientStreamReq: num,
		}

		err := stream.Send(req)
		if err != nil {
			fmt.Println("testClientStream stream.Send err:", err)
			return
		}

		// time.Sleep(5 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		fmt.Println("testClientStream stream.CloseAndRecv err:", err)
		return
	}

	fmt.Println("testClientStream res", res)
}

func (gameControllerClient *GameControllerClient) BidirectionalStream(data []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), TIME_OUT)
	defer cancel()

	stream, err := gameControllerClient.service.BidirectionalStream(ctx)

	if err != nil {
		return fmt.Errorf("bidirectionalStream client.BidirectionalStream err: ", err)
	}

	// 開啟 channel 接收訊號
	waitRes := make(chan error)
	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				fmt.Println("no more data")
				waitRes <- nil
			}

			if err != nil {
				waitRes <- fmt.Errorf("bidirectionalStream stream recv err", err)
				return
			}

			fmt.Println("bidirectionalStream res", res)
		}
	}()

	// 發送訊號
	for _, num := range data {

		req := &ClientSDK.BidirectionalStreamReq{
			BidirectionalStreamReq: num,
		}

		err := stream.Send(req)

		if err != nil {
			return fmt.Errorf("bidirectionalStream stream.Send err:", err, stream.RecvMsg(nil))
		}

		fmt.Println("bidirectionalStream stream send", num)
	}

	err = stream.CloseSend()
	if err != nil {
		return fmt.Errorf("bidirectionalStream stream.Send err: ", err)
	}

	err = <-waitRes

	return err
}
