// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.2.2

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type TodolistHTTPServer interface {
	Add(context.Context, *AddTodoRequest) (*AddTodoReply, error)
	Delete(context.Context, *DeleteTodoRequest) (*DeleteTodoReply, error)
	ShowAll(context.Context, *ShowAllTodoRequest) (*TodoListReply, error)
	ShowKey(context.Context, *ShowKeyTodoRequest) (*TodoListReply, error)
	Update(context.Context, *UpdateTodoRequest) (*UpdateTodoReply, error)
}

func RegisterTodolistHTTPServer(s *http.Server, srv TodolistHTTPServer) {
	r := s.Route("/")
	r.POST("/api/todo", _Todolist_Add0_HTTP_Handler(srv))
	r.GET("/api/todo/all", _Todolist_ShowAll0_HTTP_Handler(srv))
	r.GET("/api/todo/key", _Todolist_ShowKey0_HTTP_Handler(srv))
	r.DELETE("/api/todo/{id}", _Todolist_Delete0_HTTP_Handler(srv))
	r.PUT("/api/todo/{id}", _Todolist_Update0_HTTP_Handler(srv))
}

func _Todolist_Add0_HTTP_Handler(srv TodolistHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AddTodoRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/todolist.v1.Todolist/Add")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Add(ctx, req.(*AddTodoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AddTodoReply)
		return ctx.Result(200, reply)
	}
}

func _Todolist_ShowAll0_HTTP_Handler(srv TodolistHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ShowAllTodoRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/todolist.v1.Todolist/ShowAll")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ShowAll(ctx, req.(*ShowAllTodoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*TodoListReply)
		return ctx.Result(200, reply)
	}
}

func _Todolist_ShowKey0_HTTP_Handler(srv TodolistHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ShowKeyTodoRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/todolist.v1.Todolist/ShowKey")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ShowKey(ctx, req.(*ShowKeyTodoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*TodoListReply)
		return ctx.Result(200, reply)
	}
}

func _Todolist_Delete0_HTTP_Handler(srv TodolistHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteTodoRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/todolist.v1.Todolist/Delete")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Delete(ctx, req.(*DeleteTodoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteTodoReply)
		return ctx.Result(200, reply)
	}
}

func _Todolist_Update0_HTTP_Handler(srv TodolistHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateTodoRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/todolist.v1.Todolist/Update")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Update(ctx, req.(*UpdateTodoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateTodoReply)
		return ctx.Result(200, reply)
	}
}

type TodolistHTTPClient interface {
	Add(ctx context.Context, req *AddTodoRequest, opts ...http.CallOption) (rsp *AddTodoReply, err error)
	Delete(ctx context.Context, req *DeleteTodoRequest, opts ...http.CallOption) (rsp *DeleteTodoReply, err error)
	ShowAll(ctx context.Context, req *ShowAllTodoRequest, opts ...http.CallOption) (rsp *TodoListReply, err error)
	ShowKey(ctx context.Context, req *ShowKeyTodoRequest, opts ...http.CallOption) (rsp *TodoListReply, err error)
	Update(ctx context.Context, req *UpdateTodoRequest, opts ...http.CallOption) (rsp *UpdateTodoReply, err error)
}

type TodolistHTTPClientImpl struct {
	cc *http.Client
}

func NewTodolistHTTPClient(client *http.Client) TodolistHTTPClient {
	return &TodolistHTTPClientImpl{client}
}

func (c *TodolistHTTPClientImpl) Add(ctx context.Context, in *AddTodoRequest, opts ...http.CallOption) (*AddTodoReply, error) {
	var out AddTodoReply
	pattern := "/api/todo"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/todolist.v1.Todolist/Add"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TodolistHTTPClientImpl) Delete(ctx context.Context, in *DeleteTodoRequest, opts ...http.CallOption) (*DeleteTodoReply, error) {
	var out DeleteTodoReply
	pattern := "/api/todo/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/todolist.v1.Todolist/Delete"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TodolistHTTPClientImpl) ShowAll(ctx context.Context, in *ShowAllTodoRequest, opts ...http.CallOption) (*TodoListReply, error) {
	var out TodoListReply
	pattern := "/api/todo/all"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/todolist.v1.Todolist/ShowAll"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TodolistHTTPClientImpl) ShowKey(ctx context.Context, in *ShowKeyTodoRequest, opts ...http.CallOption) (*TodoListReply, error) {
	var out TodoListReply
	pattern := "/api/todo/key"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/todolist.v1.Todolist/ShowKey"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TodolistHTTPClientImpl) Update(ctx context.Context, in *UpdateTodoRequest, opts ...http.CallOption) (*UpdateTodoReply, error) {
	var out UpdateTodoReply
	pattern := "/api/todo/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/todolist.v1.Todolist/Update"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
