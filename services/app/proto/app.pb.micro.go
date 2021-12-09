// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/app.proto

package app

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/micro/v3/service/api"
	client "github.com/micro/micro/v3/service/client"
	server "github.com/micro/micro/v3/service/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for App service

func NewAppEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for App service

type AppService interface {
	Reserve(ctx context.Context, in *ReserveRequest, opts ...client.CallOption) (*ReserveResponse, error)
	Vote(ctx context.Context, in *VoteRequest, opts ...client.CallOption) (*VoteResponse, error)
}

type appService struct {
	c    client.Client
	name string
}

func NewAppService(name string, c client.Client) AppService {
	return &appService{
		c:    c,
		name: name,
	}
}

func (c *appService) Reserve(ctx context.Context, in *ReserveRequest, opts ...client.CallOption) (*ReserveResponse, error) {
	req := c.c.NewRequest(c.name, "App.Reserve", in)
	out := new(ReserveResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appService) Vote(ctx context.Context, in *VoteRequest, opts ...client.CallOption) (*VoteResponse, error) {
	req := c.c.NewRequest(c.name, "App.Vote", in)
	out := new(VoteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for App service

type AppHandler interface {
	Reserve(context.Context, *ReserveRequest, *ReserveResponse) error
	Vote(context.Context, *VoteRequest, *VoteResponse) error
}

func RegisterAppHandler(s server.Server, hdlr AppHandler, opts ...server.HandlerOption) error {
	type app interface {
		Reserve(ctx context.Context, in *ReserveRequest, out *ReserveResponse) error
		Vote(ctx context.Context, in *VoteRequest, out *VoteResponse) error
	}
	type App struct {
		app
	}
	h := &appHandler{hdlr}
	return s.Handle(s.NewHandler(&App{h}, opts...))
}

type appHandler struct {
	AppHandler
}

func (h *appHandler) Reserve(ctx context.Context, in *ReserveRequest, out *ReserveResponse) error {
	return h.AppHandler.Reserve(ctx, in, out)
}

func (h *appHandler) Vote(ctx context.Context, in *VoteRequest, out *VoteResponse) error {
	return h.AppHandler.Vote(ctx, in, out)
}