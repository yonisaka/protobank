// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: book/book.proto

package book

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

// BookServiceClient is the client API for BookService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookServiceClient interface {
	GetBook(ctx context.Context, in *BookByIDRequest, opts ...grpc.CallOption) (*BookResponse, error)
	GetBookList(ctx context.Context, in *BookListQuery, opts ...grpc.CallOption) (*BooksResponse, error)
	CreateBook(ctx context.Context, in *BookCreateRequest, opts ...grpc.CallOption) (*BookResponse, error)
	UpdateBook(ctx context.Context, in *BookUpdateRequest, opts ...grpc.CallOption) (*BookResponse, error)
	DeleteBook(ctx context.Context, in *BookByIDRequest, opts ...grpc.CallOption) (*BookDeleteResponse, error)
}

type bookServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBookServiceClient(cc grpc.ClientConnInterface) BookServiceClient {
	return &bookServiceClient{cc}
}

func (c *bookServiceClient) GetBook(ctx context.Context, in *BookByIDRequest, opts ...grpc.CallOption) (*BookResponse, error) {
	out := new(BookResponse)
	err := c.cc.Invoke(ctx, "/book.BookService/GetBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) GetBookList(ctx context.Context, in *BookListQuery, opts ...grpc.CallOption) (*BooksResponse, error) {
	out := new(BooksResponse)
	err := c.cc.Invoke(ctx, "/book.BookService/GetBookList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) CreateBook(ctx context.Context, in *BookCreateRequest, opts ...grpc.CallOption) (*BookResponse, error) {
	out := new(BookResponse)
	err := c.cc.Invoke(ctx, "/book.BookService/CreateBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) UpdateBook(ctx context.Context, in *BookUpdateRequest, opts ...grpc.CallOption) (*BookResponse, error) {
	out := new(BookResponse)
	err := c.cc.Invoke(ctx, "/book.BookService/UpdateBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) DeleteBook(ctx context.Context, in *BookByIDRequest, opts ...grpc.CallOption) (*BookDeleteResponse, error) {
	out := new(BookDeleteResponse)
	err := c.cc.Invoke(ctx, "/book.BookService/DeleteBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookServiceServer is the server API for BookService service.
// All implementations must embed UnimplementedBookServiceServer
// for forward compatibility
type BookServiceServer interface {
	GetBook(context.Context, *BookByIDRequest) (*BookResponse, error)
	GetBookList(context.Context, *BookListQuery) (*BooksResponse, error)
	CreateBook(context.Context, *BookCreateRequest) (*BookResponse, error)
	UpdateBook(context.Context, *BookUpdateRequest) (*BookResponse, error)
	DeleteBook(context.Context, *BookByIDRequest) (*BookDeleteResponse, error)
	mustEmbedUnimplementedBookServiceServer()
}

// UnimplementedBookServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBookServiceServer struct {
}

func (UnimplementedBookServiceServer) GetBook(context.Context, *BookByIDRequest) (*BookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBook not implemented")
}
func (UnimplementedBookServiceServer) GetBookList(context.Context, *BookListQuery) (*BooksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookList not implemented")
}
func (UnimplementedBookServiceServer) CreateBook(context.Context, *BookCreateRequest) (*BookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBook not implemented")
}
func (UnimplementedBookServiceServer) UpdateBook(context.Context, *BookUpdateRequest) (*BookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBook not implemented")
}
func (UnimplementedBookServiceServer) DeleteBook(context.Context, *BookByIDRequest) (*BookDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBook not implemented")
}
func (UnimplementedBookServiceServer) mustEmbedUnimplementedBookServiceServer() {}

// UnsafeBookServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookServiceServer will
// result in compilation errors.
type UnsafeBookServiceServer interface {
	mustEmbedUnimplementedBookServiceServer()
}

func RegisterBookServiceServer(s grpc.ServiceRegistrar, srv BookServiceServer) {
	s.RegisterService(&BookService_ServiceDesc, srv)
}

func _BookService_GetBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).GetBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/book.BookService/GetBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).GetBook(ctx, req.(*BookByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_GetBookList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookListQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).GetBookList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/book.BookService/GetBookList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).GetBookList(ctx, req.(*BookListQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_CreateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).CreateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/book.BookService/CreateBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).CreateBook(ctx, req.(*BookCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_UpdateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).UpdateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/book.BookService/UpdateBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).UpdateBook(ctx, req.(*BookUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_DeleteBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).DeleteBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/book.BookService/DeleteBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).DeleteBook(ctx, req.(*BookByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BookService_ServiceDesc is the grpc.ServiceDesc for BookService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "book.BookService",
	HandlerType: (*BookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBook",
			Handler:    _BookService_GetBook_Handler,
		},
		{
			MethodName: "GetBookList",
			Handler:    _BookService_GetBookList_Handler,
		},
		{
			MethodName: "CreateBook",
			Handler:    _BookService_CreateBook_Handler,
		},
		{
			MethodName: "UpdateBook",
			Handler:    _BookService_UpdateBook_Handler,
		},
		{
			MethodName: "DeleteBook",
			Handler:    _BookService_DeleteBook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "book/book.proto",
}
