package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"
	"vending/proto/example/greeter"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

const (
	address     = "localhost:50011"
	defaultName = "world"
)

func main() {
	// 创建 rpc 连接配置
	opts := []grpc.DialOption{
		// 配置长连接 心跳参数
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			//Time:                10 * time.Second,
			//Timeout:             5 * time.Second,
			PermitWithoutStream: true,
		}),
	}
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// 创建连接
	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		log.Fatalf("did not connect: %#v \n", err)
	}

	defer conn.Close()
	// 生成 AnotherService client
	//ac := another.NewAnotherServiceClient(conn)
	// 生成 context
	// ctx1, cancel1 := context.WithTimeout(context.Background(), time.Second)
	// defer cancel1()
	// // 调用 服务的方法
	// ar, err := ac.AnotherHello(ctx1, &another.AnotherRequest{Name: "another..."})
	// if err != nil {
	// 	log.Fatalf("could not greet: %#v \n", err)
	// }
	// // 打印返回
	// log.Printf("Greeting: %s \n", ar.Message)

	// 生成 NewGreeterClient client
	c := greeter.NewGreeterClient(conn)

	name := defaultName

	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	// 生成 context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// normal
	log.Printf("=========== normal =========== \n")
	r, err := c.SayHello(ctx, &greeter.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %#v \n", err)
	}
	log.Printf("Greeting: %s \n", r.Message)

	// server side stream
	log.Printf("=========== server side stream =========== \n")
	stream, err := c.LotsOfReplies(ctx, &greeter.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %#v \n", err)
	}
	for {
		replay, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%#v.LotsOfReplies() = _, %#v \n", c, err)
		}
		log.Printf("Greeting: %s \n", replay.Message)
	}
	stream.CloseSend()

	// client side stream
	log.Printf("=========== client side stream =========== \n")
	streamClientSide, err := c.LotsOfGreetings(ctx)
	if err != nil {
		log.Fatalf("could not greet: %#v \n", err)
	}
	for i := 0; i < 10; i++ {
		if err := streamClientSide.Send(&greeter.HelloRequest1{
			Name:  name,
			Index: int32(i),
		}); err != nil {
			log.Fatalf("send err: %#v \n", err)
		}
	}
	reply, err := streamClientSide.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", c, err, nil)
	}
	log.Printf("Greeting: %s \n", reply.Message)

	// Bidirectional stream
	start := time.Now().Unix()
	log.Printf("=========== bidirectional stream =========== \n")
	ctxTimeout, cancelTimeout := context.WithTimeout(context.Background(), 60*time.Minute)
	defer cancelTimeout()
	streamBidi, err := c.BidiHello(ctxTimeout)
	if err != nil {
		log.Fatalf("%v.BidiHello() got error %v, want %v", c, err, nil)
	}
	watic := make(chan struct{})
	// 读取
	go func() {
		for {
			in, err := streamBidi.Recv()
			if err == io.EOF {
				close(watic)
				return
			}
			if err != nil {
				end := time.Now().Unix()
				fmt.Printf("time: %d \n", end-start)
				log.Fatalf("Failed to receive a note: %#v", err.Error())
			}

			fmt.Println("AI: %#s \n", in.Message)
		}
	}()
	// 发送
	for {
		request := &greeter.HelloRequest1{}
		fmt.Scanln(&request.Name)
		if request.Name == "quit" {
			break
		}
		streamBidi.Trailer()
		if err := streamBidi.Send(request); err != nil {
			log.Fatalf("Failed to send a req: %#v", err)
		}
	}

	streamBidi.CloseSend()
	<-watic
}
