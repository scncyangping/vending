// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package greeter

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreeterClient interface {
	// 接口声明, 类似http 的一来一回短消息
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	// server stream    client端短连接 server端长连接
	LotsOfReplies(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (Greeter_LotsOfRepliesClient, error)
	// client stream    server端短连接 client端长连接
	LotsOfGreetings(ctx context.Context, opts ...grpc.CallOption) (Greeter_LotsOfGreetingsClient, error)
	// Bidirectional streaming      client,server全部长连接
	BidiHello(ctx context.Context, opts ...grpc.CallOption) (Greeter_BidiHelloClient, error)
}

type greeterClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterClient(cc grpc.ClientConnInterface) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/Greeter/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) LotsOfReplies(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (Greeter_LotsOfRepliesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Greeter_ServiceDesc.Streams[0], "/Greeter/LotsOfReplies", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterLotsOfRepliesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Greeter_LotsOfRepliesClient interface {
	Recv() (*HelloReply, error)
	grpc.ClientStream
}

type greeterLotsOfRepliesClient struct {
	grpc.ClientStream
}

func (x *greeterLotsOfRepliesClient) Recv() (*HelloReply, error) {
	m := new(HelloReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greeterClient) LotsOfGreetings(ctx context.Context, opts ...grpc.CallOption) (Greeter_LotsOfGreetingsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Greeter_ServiceDesc.Streams[1], "/Greeter/LotsOfGreetings", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterLotsOfGreetingsClient{stream}
	return x, nil
}

type Greeter_LotsOfGreetingsClient interface {
	Send(*HelloRequest1) error
	CloseAndRecv() (*HelloReply, error)
	grpc.ClientStream
}

type greeterLotsOfGreetingsClient struct {
	grpc.ClientStream
}

func (x *greeterLotsOfGreetingsClient) Send(m *HelloRequest1) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greeterLotsOfGreetingsClient) CloseAndRecv() (*HelloReply, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(HelloReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greeterClient) BidiHello(ctx context.Context, opts ...grpc.CallOption) (Greeter_BidiHelloClient, error) {
	stream, err := c.cc.NewStream(ctx, &Greeter_ServiceDesc.Streams[2], "/Greeter/BidiHello", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterBidiHelloClient{stream}
	return x, nil
}

type Greeter_BidiHelloClient interface {
	Send(*HelloRequest1) error
	Recv() (*HelloReply, error)
	grpc.ClientStream
}

type greeterBidiHelloClient struct {
	grpc.ClientStream
}

func (x *greeterBidiHelloClient) Send(m *HelloRequest1) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greeterBidiHelloClient) Recv() (*HelloReply, error) {
	m := new(HelloReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreeterServer is the server API for Greeter service.
// All implementations must embed UnimplementedGreeterServer
// for forward compatibility
type GreeterServer interface {
	// 接口声明, 类似http 的一来一回短消息
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	// server stream    client端短连接 server端长连接
	LotsOfReplies(*HelloRequest, Greeter_LotsOfRepliesServer) error
	// client stream    server端短连接 client端长连接
	LotsOfGreetings(Greeter_LotsOfGreetingsServer) error
	// Bidirectional streaming      client,server全部长连接
	BidiHello(Greeter_BidiHelloServer) error
	mustEmbedUnimplementedGreeterServer()
}

// UnimplementedGreeterServer must be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (UnimplementedGreeterServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedGreeterServer) LotsOfReplies(*HelloRequest, Greeter_LotsOfRepliesServer) error {
	return status.Errorf(codes.Unimplemented, "method LotsOfReplies not implemented")
}
func (UnimplementedGreeterServer) LotsOfGreetings(Greeter_LotsOfGreetingsServer) error {
	return status.Errorf(codes.Unimplemented, "method LotsOfGreetings not implemented")
}
func (UnimplementedGreeterServer) BidiHello(Greeter_BidiHelloServer) error {
	return status.Errorf(codes.Unimplemented, "method BidiHello not implemented")
}
func (UnimplementedGreeterServer) mustEmbedUnimplementedGreeterServer() {}

// UnsafeGreeterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreeterServer will
// result in compilation errors.
type UnsafeGreeterServer interface {
	mustEmbedUnimplementedGreeterServer()
}

func RegisterGreeterServer(s grpc.ServiceRegistrar, srv GreeterServer) {
	s.RegisterService(&Greeter_ServiceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_LotsOfReplies_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HelloRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreeterServer).LotsOfReplies(m, &greeterLotsOfRepliesServer{stream})
}

type Greeter_LotsOfRepliesServer interface {
	Send(*HelloReply) error
	grpc.ServerStream
}

type greeterLotsOfRepliesServer struct {
	grpc.ServerStream
}

func (x *greeterLotsOfRepliesServer) Send(m *HelloReply) error {
	return x.ServerStream.SendMsg(m)
}

func _Greeter_LotsOfGreetings_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreeterServer).LotsOfGreetings(&greeterLotsOfGreetingsServer{stream})
}

type Greeter_LotsOfGreetingsServer interface {
	SendAndClose(*HelloReply) error
	Recv() (*HelloRequest1, error)
	grpc.ServerStream
}

type greeterLotsOfGreetingsServer struct {
	grpc.ServerStream
}

func (x *greeterLotsOfGreetingsServer) SendAndClose(m *HelloReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greeterLotsOfGreetingsServer) Recv() (*HelloRequest1, error) {
	m := new(HelloRequest1)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Greeter_BidiHello_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreeterServer).BidiHello(&greeterBidiHelloServer{stream})
}

type Greeter_BidiHelloServer interface {
	Send(*HelloReply) error
	Recv() (*HelloRequest1, error)
	grpc.ServerStream
}

type greeterBidiHelloServer struct {
	grpc.ServerStream
}

func (x *greeterBidiHelloServer) Send(m *HelloReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greeterBidiHelloServer) Recv() (*HelloRequest1, error) {
	m := new(HelloRequest1)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Greeter_ServiceDesc is the grpc.ServiceDesc for Greeter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Greeter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "LotsOfReplies",
			Handler:       _Greeter_LotsOfReplies_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "LotsOfGreetings",
			Handler:       _Greeter_LotsOfGreetings_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "BidiHello",
			Handler:       _Greeter_BidiHello_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "helloworld.proto",
}
