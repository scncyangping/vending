package example

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"vending/proto/example/another"
	"vending/proto/example/greeter"

	"google.golang.org/grpc"
)

type Server struct {
	greeter.UnimplementedGreeterServer
}

// 实现服务接口 简单的短连接
func (s *Server) SayHello(ctx context.Context, in *greeter.HelloRequest) (*greeter.HelloReply, error) {
	// 调用 service, 链接数据库, 处理缓存等...
	return &greeter.HelloReply{Message: "Hello " + in.Name}, nil
}

// 实现服务接口
func (s *Server) LotsOfReplies(in *greeter.HelloRequest, stream greeter.Greeter_LotsOfRepliesServer) error {
	// 服务端 长连接发送信息, 发送给客户端大量数据
	for i := 0; i < 10; i++ {
		stream.Send(&greeter.HelloReply{Message: fmt.Sprintf("Hello %s %d", in.Name, i)})
	}
	return nil
}

// 实现服务接口
func (s *Server) LotsOfGreetings(stream greeter.Greeter_LotsOfGreetingsServer) error {
	// 客户端 长连接 发送连续大量数据
	var total int32
	var name string
	for {
		greeting, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&greeter.HelloReply{
				Message: fmt.Sprintf("Hello %s, total %d", name, total),
			})
		}
		if err != nil {
			return err
		}
		name = greeting.Name
		total += greeting.Index
	}
}

// 实现服务接口
func (s *Server) BidiHello(stream greeter.Greeter_BidiHelloServer) error {
	// 客户端服务端分别建立长连接, 发送响应的数据
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		message := strings.Replace(in.Name, "吗", "", -1)
		message = strings.Replace(message, "?", "!", -1)
		err = stream.Send(&greeter.HelloReply{Message: message})
		if err != nil {
			return err
		}
	}
}

type AServer struct{}

// 实现服务接口
func (as *AServer) AnotherHello(ctx context.Context, request *another.AnotherRequest) (*another.AnotherReplay, error) {
	return &another.AnotherReplay{Message: request.Name}, nil
}

func RunServer(ctx context.Context, grpcServer *grpc.Server, port string, httpServer *http.Server) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	greeter.RegisterGreeterServer(grpcServer, &Server{})

	// graceful shutdowns
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	go func() {
		for range sigChan {
			fmt.Println("sig received, shutting down PBnJ")
			grpcServer.GracefulStop()
			<-ctx.Done()
		}
	}()

	go func() {
		<-ctx.Done()
		fmt.Println("ctx cancelled, shutting down PBnJ")
		grpcServer.GracefulStop()
	}()

	fmt.Println("starting gRPC server")
	return grpcServer.Serve(listen)
}
