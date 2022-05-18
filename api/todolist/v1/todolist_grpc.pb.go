// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: todolist/v1/todolist.proto

package v1

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

// TodolistClient is the client API for Todolist service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TodolistClient interface {
	// Register a user
	RegisterUser(ctx context.Context, in *RegisterUserRequest, opts ...grpc.CallOption) (*RegisterUserReply, error)
	// Login a user
	LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserReply, error)
	// Add a todothings
	AddTodo(ctx context.Context, in *AddTodoRequest, opts ...grpc.CallOption) (*AddTodoReply, error)
	// Show all todothings
	ShowAllTodo(ctx context.Context, in *ShowAllTodoRequest, opts ...grpc.CallOption) (*ShowAllTodoReply, error)
	// Show KeyMatched todothings
	ShowKeyTodo(ctx context.Context, in *ShowKeyTodoRequest, opts ...grpc.CallOption) (*ShowKeyTodoReply, error)
	// Delete a todothings
	DeleteTodo(ctx context.Context, in *DeleteTodoRequest, opts ...grpc.CallOption) (*DeleteTodoReply, error)
	// Modify a todothings
	ModifyTodo(ctx context.Context, in *ModifyTodoRequest, opts ...grpc.CallOption) (*ModifyTodoReply, error)
}

type todolistClient struct {
	cc grpc.ClientConnInterface
}

func NewTodolistClient(cc grpc.ClientConnInterface) TodolistClient {
	return &todolistClient{cc}
}

func (c *todolistClient) RegisterUser(ctx context.Context, in *RegisterUserRequest, opts ...grpc.CallOption) (*RegisterUserReply, error) {
	out := new(RegisterUserReply)
	err := c.cc.Invoke(ctx, "/todolist.v1.Todolist/RegisterUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todolistClient) LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserReply, error) {
	out := new(LoginUserReply)
	err := c.cc.Invoke(ctx, "/todolist.v1.Todolist/LoginUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todolistClient) AddTodo(ctx context.Context, in *AddTodoRequest, opts ...grpc.CallOption) (*AddTodoReply, error) {
	out := new(AddTodoReply)
	err := c.cc.Invoke(ctx, "/todolist.v1.Todolist/AddTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todolistClient) ShowAllTodo(ctx context.Context, in *ShowAllTodoRequest, opts ...grpc.CallOption) (*ShowAllTodoReply, error) {
	out := new(ShowAllTodoReply)
	err := c.cc.Invoke(ctx, "/todolist.v1.Todolist/ShowAllTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todolistClient) ShowKeyTodo(ctx context.Context, in *ShowKeyTodoRequest, opts ...grpc.CallOption) (*ShowKeyTodoReply, error) {
	out := new(ShowKeyTodoReply)
	err := c.cc.Invoke(ctx, "/todolist.v1.Todolist/ShowKeyTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todolistClient) DeleteTodo(ctx context.Context, in *DeleteTodoRequest, opts ...grpc.CallOption) (*DeleteTodoReply, error) {
	out := new(DeleteTodoReply)
	err := c.cc.Invoke(ctx, "/todolist.v1.Todolist/DeleteTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todolistClient) ModifyTodo(ctx context.Context, in *ModifyTodoRequest, opts ...grpc.CallOption) (*ModifyTodoReply, error) {
	out := new(ModifyTodoReply)
	err := c.cc.Invoke(ctx, "/todolist.v1.Todolist/ModifyTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodolistServer is the server API for Todolist service.
// All implementations must embed UnimplementedTodolistServer
// for forward compatibility
type TodolistServer interface {
	// Register a user
	RegisterUser(context.Context, *RegisterUserRequest) (*RegisterUserReply, error)
	// Login a user
	LoginUser(context.Context, *LoginUserRequest) (*LoginUserReply, error)
	// Add a todothings
	AddTodo(context.Context, *AddTodoRequest) (*AddTodoReply, error)
	// Show all todothings
	ShowAllTodo(context.Context, *ShowAllTodoRequest) (*ShowAllTodoReply, error)
	// Show KeyMatched todothings
	ShowKeyTodo(context.Context, *ShowKeyTodoRequest) (*ShowKeyTodoReply, error)
	// Delete a todothings
	DeleteTodo(context.Context, *DeleteTodoRequest) (*DeleteTodoReply, error)
	// Modify a todothings
	ModifyTodo(context.Context, *ModifyTodoRequest) (*ModifyTodoReply, error)
	mustEmbedUnimplementedTodolistServer()
}

// UnimplementedTodolistServer must be embedded to have forward compatible implementations.
type UnimplementedTodolistServer struct {
}

func (UnimplementedTodolistServer) RegisterUser(context.Context, *RegisterUserRequest) (*RegisterUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUser not implemented")
}
func (UnimplementedTodolistServer) LoginUser(context.Context, *LoginUserRequest) (*LoginUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedTodolistServer) AddTodo(context.Context, *AddTodoRequest) (*AddTodoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTodo not implemented")
}
func (UnimplementedTodolistServer) ShowAllTodo(context.Context, *ShowAllTodoRequest) (*ShowAllTodoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowAllTodo not implemented")
}
func (UnimplementedTodolistServer) ShowKeyTodo(context.Context, *ShowKeyTodoRequest) (*ShowKeyTodoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowKeyTodo not implemented")
}
func (UnimplementedTodolistServer) DeleteTodo(context.Context, *DeleteTodoRequest) (*DeleteTodoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTodo not implemented")
}
func (UnimplementedTodolistServer) ModifyTodo(context.Context, *ModifyTodoRequest) (*ModifyTodoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyTodo not implemented")
}
func (UnimplementedTodolistServer) mustEmbedUnimplementedTodolistServer() {}

// UnsafeTodolistServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TodolistServer will
// result in compilation errors.
type UnsafeTodolistServer interface {
	mustEmbedUnimplementedTodolistServer()
}

func RegisterTodolistServer(s grpc.ServiceRegistrar, srv TodolistServer) {
	s.RegisterService(&Todolist_ServiceDesc, srv)
}

func _Todolist_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodolistServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todolist.v1.Todolist/RegisterUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodolistServer).RegisterUser(ctx, req.(*RegisterUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todolist_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodolistServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todolist.v1.Todolist/LoginUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodolistServer).LoginUser(ctx, req.(*LoginUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todolist_AddTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodolistServer).AddTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todolist.v1.Todolist/AddTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodolistServer).AddTodo(ctx, req.(*AddTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todolist_ShowAllTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowAllTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodolistServer).ShowAllTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todolist.v1.Todolist/ShowAllTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodolistServer).ShowAllTodo(ctx, req.(*ShowAllTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todolist_ShowKeyTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowKeyTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodolistServer).ShowKeyTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todolist.v1.Todolist/ShowKeyTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodolistServer).ShowKeyTodo(ctx, req.(*ShowKeyTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todolist_DeleteTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodolistServer).DeleteTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todolist.v1.Todolist/DeleteTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodolistServer).DeleteTodo(ctx, req.(*DeleteTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todolist_ModifyTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodolistServer).ModifyTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todolist.v1.Todolist/ModifyTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodolistServer).ModifyTodo(ctx, req.(*ModifyTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Todolist_ServiceDesc is the grpc.ServiceDesc for Todolist service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Todolist_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "todolist.v1.Todolist",
	HandlerType: (*TodolistServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterUser",
			Handler:    _Todolist_RegisterUser_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _Todolist_LoginUser_Handler,
		},
		{
			MethodName: "AddTodo",
			Handler:    _Todolist_AddTodo_Handler,
		},
		{
			MethodName: "ShowAllTodo",
			Handler:    _Todolist_ShowAllTodo_Handler,
		},
		{
			MethodName: "ShowKeyTodo",
			Handler:    _Todolist_ShowKeyTodo_Handler,
		},
		{
			MethodName: "DeleteTodo",
			Handler:    _Todolist_DeleteTodo_Handler,
		},
		{
			MethodName: "ModifyTodo",
			Handler:    _Todolist_ModifyTodo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todolist/v1/todolist.proto",
}