package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	pb "grpctest/proto"
	"log"
	"net"
	"net/http"
	"sync"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (res *pb.HelloReply, err error) {
	log.Printf("recive:%v", req.Name)
	return &pb.HelloReply{Message: "responce"}, nil
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		registerGateWay(&wg)
	}()
	go func() {
		registerGRPC(&wg)
	}()
	wg.Wait()
}

func registerGateWay(wg *sync.WaitGroup) {
	conn, err := grpc.DialContext(
		context.Background(),
		"127.0.0.1:50051",
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("err :%v", err)
	}
	mux := runtime.NewServeMux()

	gwServer := &http.Server{
		Handler: mux,
		Addr:    ":8090",
	}
	err = pb.RegisterGreeterHandler(context.Background(), mux, conn)
	if err != nil {
		log.Fatalf("RegisterGreeterHandler err :%v", err)
	}
	err = gwServer.ListenAndServe()
	if err != nil {
		log.Fatalf("listenAndServe:%v", err)
	}
	wg.Done()
}

func registerGRPC(wg *sync.WaitGroup) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("fail to listen:%v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v:", lis.Addr())
	s.Serve(lis)
	wg.Done()
}
