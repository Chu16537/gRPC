package service

import (
	"context"
	"errors"
	"fmt"
	"io"
	"time"

	ClientSDK "grpc_test/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GameController struct{}

func NewGameController() *GameController {
	return &GameController{}
}

func (server *GameController) Unary(ctx context.Context, req *ClientSDK.UnaryReq) (*ClientSDK.UnaryRes, error) {

	fmt.Println("unary req", req)

	// timeout
	// time.Sleep(5 * time.Second)

	res := &ClientSDK.UnaryRes{
		UnaryRes: "unaryRes test",
	}

	if err := ctxError(ctx); err != nil {
		return nil, err
	}

	fmt.Println("unary res", res)

	return res, nil
}

func (server *GameController) TestStreamServer(req *ClientSDK.ServerStreamReq, stream ClientSDK.GameController_TestStreamServerServer) error {

	ctx := stream.Context()

	array := [5]string{"a", "s", "d", "f", "g"}

	for _, num := range array {

		if err := ctxError(ctx); err != nil {
			return err
		}

		req := &ClientSDK.ServerStreamRes{
			ServerStreamRes: num,
		}

		err := stream.Send(req)

		if err != nil {
			return err
		}

		fmt.Println("stream send:", num)

		time.Sleep(1 * time.Second)
	}

	return nil
}

func (server *GameController) TestStreamClient(stream ClientSDK.GameController_TestStreamClientServer) error {

	ctx := stream.Context()

	var a int32

	for {
		if err := ctxError(ctx); err != nil {
			return err
		}

		req, err := stream.Recv()

		// 表示 server 的 stream 資料傳完了
		if err == io.EOF {
			fmt.Println("no more data")
			break
		}

		// 表示有錯誤發生
		if err != nil {
			fmt.Println("TestStreamClient stream.Recv err:", err)
			return status.Errorf(codes.Unknown, " Recv err")
		}

		// 讀取 server stream 傳來的資料
		a += req.GetClientStreamReq()
		fmt.Println("a:", a)
		// time.Sleep(5 * time.Second)
	}

	res := &ClientSDK.ClientStreamRes{
		ClientStreamRes: a,
	}

	err := stream.SendAndClose(res)
	if err != nil {
		fmt.Println("req cannot stream SendAndClose: ", err)
		return status.Errorf(codes.Unknown, "req cannot stream SendAndClose")
	}

	return nil
}

func (server *GameController) BidirectionalStream(stream ClientSDK.GameController_BidirectionalStreamServer) error {

	for {
		err := ctxError(stream.Context())

		if err != nil {
			return err
		}

		req, err := stream.Recv()

		if err == io.EOF {
			fmt.Println("no more data")
			break
		}

		if err != nil {
			return status.Errorf(codes.Unknown, "req stream recv err")
		}

		num := req.GetBidirectionalStreamReq()
		fmt.Println("num", num)

		res := &ClientSDK.BidirectionalStreamRes{
			BidirectionalStreamRes: "q",
		}

		err = stream.Send(res)

		if err != nil {
			return status.Errorf(codes.Unknown, "res stream send err")
		}
	}

	return nil
}

// 判斷是否 超時 or client 取消
func ctxError(ctx context.Context) error {
	switch ctx.Err() {
	case context.Canceled:
		fmt.Println("client is Canceled")
		return errors.New("client is Canceled")
	case context.DeadlineExceeded:
		fmt.Println("deadline is exceeded")
		return errors.New("deadline is exceeded")
	default:
		return nil
	}
}
