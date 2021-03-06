package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math"
	"net"
	"time"

	calc_pb "../calculator_pb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) MaxService(stream calc_pb.CalculatorService_MaxServiceServer) error {

	max := int64(math.MinInt64)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			fmt.Printf("Error while receiving request: %v\n", err)
			return err
		}

		n := req.GetNum()
		if n > max {
			max = n
		}
		fmt.Printf("Received: %v, max is: %v\n", n, max)
		sendErr := stream.Send(&calc_pb.MaxResponse{
			Max: max,
		})
		if sendErr != nil {
			fmt.Printf("Error sending response: %v", sendErr)
			return sendErr
		}
	} // .for

} // .MaxService

func (*server) SumService(ctx context.Context, req *calc_pb.SumRequest) (*calc_pb.SumResponse, error) {
	fmt.Printf("Sum service func invoked with request: %v\n", req)

	a := req.GetSum().GetA()
	b := req.GetSum().GetB()

	result := a + b
	resp := &calc_pb.SumResponse{
		Result: result,
	}
	return resp, nil
} // .SumService

func (*server) FibonacciService(req *calc_pb.FibonacciRequest, stream calc_pb.CalculatorService_FibonacciServiceServer) error {
	fmt.Printf("Starting Fibonacci streaming service, with request: %v\n", req)

	fibo1 := 0
	fibo2 := 1

	n := req.GetNum()
	var i int32
	for i = 0; i <= n; i++ {
		fibo3 := fibo1 + fibo2

		resp := &calc_pb.FibonacciResponse{
			ResultNum: int32(fibo3),
		}
		stream.Send(resp)
		time.Sleep(500 * time.Millisecond)

		fibo1 = fibo2
		fibo2 = fibo3
	} // .for

	return nil
} // .FibonacciService

func (*server) MeanService(stream calc_pb.CalculatorService_MeanServiceServer) error {
	fmt.Println("Calculator MeanService has been invoked with streaming request")

	var mean, many int64
	mean = 0
	many = 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&calc_pb.MeanResponse{
				Mean: mean / many,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}
		many++
		mean += req.GetNumber()

	}
} // .MeanService

func main() {

	fmt.Printf("Starting Calculator RPC server\nWaiting client request/s...\n")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	calc_pb.RegisterCalculatorServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

} // .main
