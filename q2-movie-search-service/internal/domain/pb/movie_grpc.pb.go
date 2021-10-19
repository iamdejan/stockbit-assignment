// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// MovieSearchServiceClient is the client API for MovieSearchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MovieSearchServiceClient interface {
	Search(ctx context.Context, in *MoviePreviewListRequest, opts ...grpc.CallOption) (*MoviePreviewListResponse, error)
	Get(ctx context.Context, in *MovieRequest, opts ...grpc.CallOption) (*Movie, error)
}

type movieSearchServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMovieSearchServiceClient(cc grpc.ClientConnInterface) MovieSearchServiceClient {
	return &movieSearchServiceClient{cc}
}

func (c *movieSearchServiceClient) Search(ctx context.Context, in *MoviePreviewListRequest, opts ...grpc.CallOption) (*MoviePreviewListResponse, error) {
	out := new(MoviePreviewListResponse)
	err := c.cc.Invoke(ctx, "/MovieSearchService/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieSearchServiceClient) Get(ctx context.Context, in *MovieRequest, opts ...grpc.CallOption) (*Movie, error) {
	out := new(Movie)
	err := c.cc.Invoke(ctx, "/MovieSearchService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MovieSearchServiceServer is the server API for MovieSearchService service.
// All implementations must embed UnimplementedMovieSearchServiceServer
// for forward compatibility
type MovieSearchServiceServer interface {
	Search(context.Context, *MoviePreviewListRequest) (*MoviePreviewListResponse, error)
	Get(context.Context, *MovieRequest) (*Movie, error)
	mustEmbedUnimplementedMovieSearchServiceServer()
}

// UnimplementedMovieSearchServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMovieSearchServiceServer struct {
}

func (UnimplementedMovieSearchServiceServer) Search(context.Context, *MoviePreviewListRequest) (*MoviePreviewListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedMovieSearchServiceServer) Get(context.Context, *MovieRequest) (*Movie, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedMovieSearchServiceServer) mustEmbedUnimplementedMovieSearchServiceServer() {}

// UnsafeMovieSearchServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MovieSearchServiceServer will
// result in compilation errors.
type UnsafeMovieSearchServiceServer interface {
	mustEmbedUnimplementedMovieSearchServiceServer()
}

func RegisterMovieSearchServiceServer(s grpc.ServiceRegistrar, srv MovieSearchServiceServer) {
	s.RegisterService(&MovieSearchService_ServiceDesc, srv)
}

func _MovieSearchService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MoviePreviewListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieSearchServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MovieSearchService/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieSearchServiceServer).Search(ctx, req.(*MoviePreviewListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieSearchService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieSearchServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MovieSearchService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieSearchServiceServer).Get(ctx, req.(*MovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MovieSearchService_ServiceDesc is the grpc.ServiceDesc for MovieSearchService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MovieSearchService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "MovieSearchService",
	HandlerType: (*MovieSearchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _MovieSearchService_Search_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _MovieSearchService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/movie.proto",
}