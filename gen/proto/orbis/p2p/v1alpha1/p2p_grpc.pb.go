// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: orbis/p2p/v1alpha1/p2p.proto

package p2pv1alpha1

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

const (
	P2PService_Host_FullMethodName      = "/orbis.p2p.v1alpha1.P2PService/Host"
	P2PService_Peers_FullMethodName     = "/orbis.p2p.v1alpha1.P2PService/Peers"
	P2PService_Connect_FullMethodName   = "/orbis.p2p.v1alpha1.P2PService/Connect"
	P2PService_Send_FullMethodName      = "/orbis.p2p.v1alpha1.P2PService/Send"
	P2PService_Publish_FullMethodName   = "/orbis.p2p.v1alpha1.P2PService/Publish"
	P2PService_Subscribe_FullMethodName = "/orbis.p2p.v1alpha1.P2PService/Subscribe"
)

// P2PServiceClient is the client API for P2PService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type P2PServiceClient interface {
	// Host returns the information about the host node.
	Host(ctx context.Context, in *HostRequest, opts ...grpc.CallOption) (*HostResponse, error)
	// Peers lists information about connected peer nodes.
	Peers(ctx context.Context, in *PeersRequest, opts ...grpc.CallOption) (*PeersResponse, error)
	// Connect connects to a peer node.
	Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectResponse, error)
	// Send sends a message to a peer node.
	Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error)
	// Publish broadcasts a message to a topic.
	Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error)
	// Subscribe broadcasts a message to a topic.
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (P2PService_SubscribeClient, error)
}

type p2PServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewP2PServiceClient(cc grpc.ClientConnInterface) P2PServiceClient {
	return &p2PServiceClient{cc}
}

func (c *p2PServiceClient) Host(ctx context.Context, in *HostRequest, opts ...grpc.CallOption) (*HostResponse, error) {
	out := new(HostResponse)
	err := c.cc.Invoke(ctx, P2PService_Host_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *p2PServiceClient) Peers(ctx context.Context, in *PeersRequest, opts ...grpc.CallOption) (*PeersResponse, error) {
	out := new(PeersResponse)
	err := c.cc.Invoke(ctx, P2PService_Peers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *p2PServiceClient) Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectResponse, error) {
	out := new(ConnectResponse)
	err := c.cc.Invoke(ctx, P2PService_Connect_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *p2PServiceClient) Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error) {
	out := new(SendResponse)
	err := c.cc.Invoke(ctx, P2PService_Send_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *p2PServiceClient) Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error) {
	out := new(PublishResponse)
	err := c.cc.Invoke(ctx, P2PService_Publish_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *p2PServiceClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (P2PService_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &P2PService_ServiceDesc.Streams[0], P2PService_Subscribe_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &p2PServiceSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type P2PService_SubscribeClient interface {
	Recv() (*SubscribeResponse, error)
	grpc.ClientStream
}

type p2PServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *p2PServiceSubscribeClient) Recv() (*SubscribeResponse, error) {
	m := new(SubscribeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// P2PServiceServer is the server API for P2PService service.
// All implementations must embed UnimplementedP2PServiceServer
// for forward compatibility
type P2PServiceServer interface {
	// Host returns the information about the host node.
	Host(context.Context, *HostRequest) (*HostResponse, error)
	// Peers lists information about connected peer nodes.
	Peers(context.Context, *PeersRequest) (*PeersResponse, error)
	// Connect connects to a peer node.
	Connect(context.Context, *ConnectRequest) (*ConnectResponse, error)
	// Send sends a message to a peer node.
	Send(context.Context, *SendRequest) (*SendResponse, error)
	// Publish broadcasts a message to a topic.
	Publish(context.Context, *PublishRequest) (*PublishResponse, error)
	// Subscribe broadcasts a message to a topic.
	Subscribe(*SubscribeRequest, P2PService_SubscribeServer) error
	mustEmbedUnimplementedP2PServiceServer()
}

// UnimplementedP2PServiceServer must be embedded to have forward compatible implementations.
type UnimplementedP2PServiceServer struct {
}

func (UnimplementedP2PServiceServer) Host(context.Context, *HostRequest) (*HostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Host not implemented")
}
func (UnimplementedP2PServiceServer) Peers(context.Context, *PeersRequest) (*PeersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Peers not implemented")
}
func (UnimplementedP2PServiceServer) Connect(context.Context, *ConnectRequest) (*ConnectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedP2PServiceServer) Send(context.Context, *SendRequest) (*SendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedP2PServiceServer) Publish(context.Context, *PublishRequest) (*PublishResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (UnimplementedP2PServiceServer) Subscribe(*SubscribeRequest, P2PService_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedP2PServiceServer) mustEmbedUnimplementedP2PServiceServer() {}

// UnsafeP2PServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to P2PServiceServer will
// result in compilation errors.
type UnsafeP2PServiceServer interface {
	mustEmbedUnimplementedP2PServiceServer()
}

func RegisterP2PServiceServer(s grpc.ServiceRegistrar, srv P2PServiceServer) {
	s.RegisterService(&P2PService_ServiceDesc, srv)
}

func _P2PService_Host_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(P2PServiceServer).Host(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: P2PService_Host_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(P2PServiceServer).Host(ctx, req.(*HostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _P2PService_Peers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(P2PServiceServer).Peers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: P2PService_Peers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(P2PServiceServer).Peers(ctx, req.(*PeersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _P2PService_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(P2PServiceServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: P2PService_Connect_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(P2PServiceServer).Connect(ctx, req.(*ConnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _P2PService_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(P2PServiceServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: P2PService_Send_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(P2PServiceServer).Send(ctx, req.(*SendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _P2PService_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(P2PServiceServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: P2PService_Publish_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(P2PServiceServer).Publish(ctx, req.(*PublishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _P2PService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(P2PServiceServer).Subscribe(m, &p2PServiceSubscribeServer{stream})
}

type P2PService_SubscribeServer interface {
	Send(*SubscribeResponse) error
	grpc.ServerStream
}

type p2PServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *p2PServiceSubscribeServer) Send(m *SubscribeResponse) error {
	return x.ServerStream.SendMsg(m)
}

// P2PService_ServiceDesc is the grpc.ServiceDesc for P2PService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var P2PService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "orbis.p2p.v1alpha1.P2PService",
	HandlerType: (*P2PServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Host",
			Handler:    _P2PService_Host_Handler,
		},
		{
			MethodName: "Peers",
			Handler:    _P2PService_Peers_Handler,
		},
		{
			MethodName: "Connect",
			Handler:    _P2PService_Connect_Handler,
		},
		{
			MethodName: "Send",
			Handler:    _P2PService_Send_Handler,
		},
		{
			MethodName: "Publish",
			Handler:    _P2PService_Publish_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _P2PService_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "orbis/p2p/v1alpha1/p2p.proto",
}