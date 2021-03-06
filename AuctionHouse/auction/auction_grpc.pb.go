// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package auction

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

// AuctionHouseClient is the client API for AuctionHouse service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuctionHouseClient interface {
	Join(ctx context.Context, in *JoinMessage, opts ...grpc.CallOption) (AuctionHouse_JoinClient, error)
	Bid(ctx context.Context, in *BidMessage, opts ...grpc.CallOption) (*BidResponse, error)
	Result(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BidMessage, error)
}

type auctionHouseClient struct {
	cc grpc.ClientConnInterface
}

func NewAuctionHouseClient(cc grpc.ClientConnInterface) AuctionHouseClient {
	return &auctionHouseClient{cc}
}

func (c *auctionHouseClient) Join(ctx context.Context, in *JoinMessage, opts ...grpc.CallOption) (AuctionHouse_JoinClient, error) {
	stream, err := c.cc.NewStream(ctx, &AuctionHouse_ServiceDesc.Streams[0], "/chat.AuctionHouse/Join", opts...)
	if err != nil {
		return nil, err
	}
	x := &auctionHouseJoinClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AuctionHouse_JoinClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type auctionHouseJoinClient struct {
	grpc.ClientStream
}

func (x *auctionHouseJoinClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *auctionHouseClient) Bid(ctx context.Context, in *BidMessage, opts ...grpc.CallOption) (*BidResponse, error) {
	out := new(BidResponse)
	err := c.cc.Invoke(ctx, "/chat.AuctionHouse/Bid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionHouseClient) Result(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BidMessage, error) {
	out := new(BidMessage)
	err := c.cc.Invoke(ctx, "/chat.AuctionHouse/Result", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuctionHouseServer is the server API for AuctionHouse service.
// All implementations must embed UnimplementedAuctionHouseServer
// for forward compatibility
type AuctionHouseServer interface {
	Join(*JoinMessage, AuctionHouse_JoinServer) error
	Bid(context.Context, *BidMessage) (*BidResponse, error)
	Result(context.Context, *Empty) (*BidMessage, error)
	mustEmbedUnimplementedAuctionHouseServer()
}

// UnimplementedAuctionHouseServer must be embedded to have forward compatible implementations.
type UnimplementedAuctionHouseServer struct {
}

func (UnimplementedAuctionHouseServer) Join(*JoinMessage, AuctionHouse_JoinServer) error {
	return status.Errorf(codes.Unimplemented, "method Join not implemented")
}
func (UnimplementedAuctionHouseServer) Bid(context.Context, *BidMessage) (*BidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Bid not implemented")
}
func (UnimplementedAuctionHouseServer) Result(context.Context, *Empty) (*BidMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Result not implemented")
}
func (UnimplementedAuctionHouseServer) mustEmbedUnimplementedAuctionHouseServer() {}

// UnsafeAuctionHouseServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuctionHouseServer will
// result in compilation errors.
type UnsafeAuctionHouseServer interface {
	mustEmbedUnimplementedAuctionHouseServer()
}

func RegisterAuctionHouseServer(s grpc.ServiceRegistrar, srv AuctionHouseServer) {
	s.RegisterService(&AuctionHouse_ServiceDesc, srv)
}

func _AuctionHouse_Join_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(JoinMessage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AuctionHouseServer).Join(m, &auctionHouseJoinServer{stream})
}

type AuctionHouse_JoinServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type auctionHouseJoinServer struct {
	grpc.ServerStream
}

func (x *auctionHouseJoinServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func _AuctionHouse_Bid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BidMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionHouseServer).Bid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.AuctionHouse/Bid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionHouseServer).Bid(ctx, req.(*BidMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuctionHouse_Result_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionHouseServer).Result(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.AuctionHouse/Result",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionHouseServer).Result(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// AuctionHouse_ServiceDesc is the grpc.ServiceDesc for AuctionHouse service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuctionHouse_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chat.AuctionHouse",
	HandlerType: (*AuctionHouseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Bid",
			Handler:    _AuctionHouse_Bid_Handler,
		},
		{
			MethodName: "Result",
			Handler:    _AuctionHouse_Result_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Join",
			Handler:       _AuctionHouse_Join_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "AuctionHouse/auction/auction.proto",
}
