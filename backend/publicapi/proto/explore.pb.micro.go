// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/explore.proto

package publicapi

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

// Api Endpoints for Explore service

func NewExploreEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Explore service

type ExploreService interface {
	Index(ctx context.Context, in *IndexRequest, opts ...client.CallOption) (*IndexResponse, error)
	Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchResponse, error)
	API(ctx context.Context, in *APIRequest, opts ...client.CallOption) (*APIResponse, error)
	ListCategories(ctx context.Context, in *ListCategoriesRequest, opts ...client.CallOption) (*ListCategoriesResponse, error)
	Pricing(ctx context.Context, in *PricingRequest, opts ...client.CallOption) (*PricingResponse, error)
}

type exploreService struct {
	c    client.Client
	name string
}

func NewExploreService(name string, c client.Client) ExploreService {
	return &exploreService{
		c:    c,
		name: name,
	}
}

func (c *exploreService) Index(ctx context.Context, in *IndexRequest, opts ...client.CallOption) (*IndexResponse, error) {
	req := c.c.NewRequest(c.name, "Explore.Index", in)
	out := new(IndexResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exploreService) Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchResponse, error) {
	req := c.c.NewRequest(c.name, "Explore.Search", in)
	out := new(SearchResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exploreService) API(ctx context.Context, in *APIRequest, opts ...client.CallOption) (*APIResponse, error) {
	req := c.c.NewRequest(c.name, "Explore.API", in)
	out := new(APIResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exploreService) ListCategories(ctx context.Context, in *ListCategoriesRequest, opts ...client.CallOption) (*ListCategoriesResponse, error) {
	req := c.c.NewRequest(c.name, "Explore.ListCategories", in)
	out := new(ListCategoriesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exploreService) Pricing(ctx context.Context, in *PricingRequest, opts ...client.CallOption) (*PricingResponse, error) {
	req := c.c.NewRequest(c.name, "Explore.Pricing", in)
	out := new(PricingResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Explore service

type ExploreHandler interface {
	Index(context.Context, *IndexRequest, *IndexResponse) error
	Search(context.Context, *SearchRequest, *SearchResponse) error
	API(context.Context, *APIRequest, *APIResponse) error
	ListCategories(context.Context, *ListCategoriesRequest, *ListCategoriesResponse) error
	Pricing(context.Context, *PricingRequest, *PricingResponse) error
}

func RegisterExploreHandler(s server.Server, hdlr ExploreHandler, opts ...server.HandlerOption) error {
	type explore interface {
		Index(ctx context.Context, in *IndexRequest, out *IndexResponse) error
		Search(ctx context.Context, in *SearchRequest, out *SearchResponse) error
		API(ctx context.Context, in *APIRequest, out *APIResponse) error
		ListCategories(ctx context.Context, in *ListCategoriesRequest, out *ListCategoriesResponse) error
		Pricing(ctx context.Context, in *PricingRequest, out *PricingResponse) error
	}
	type Explore struct {
		explore
	}
	h := &exploreHandler{hdlr}
	return s.Handle(s.NewHandler(&Explore{h}, opts...))
}

type exploreHandler struct {
	ExploreHandler
}

func (h *exploreHandler) Index(ctx context.Context, in *IndexRequest, out *IndexResponse) error {
	return h.ExploreHandler.Index(ctx, in, out)
}

func (h *exploreHandler) Search(ctx context.Context, in *SearchRequest, out *SearchResponse) error {
	return h.ExploreHandler.Search(ctx, in, out)
}

func (h *exploreHandler) API(ctx context.Context, in *APIRequest, out *APIResponse) error {
	return h.ExploreHandler.API(ctx, in, out)
}

func (h *exploreHandler) ListCategories(ctx context.Context, in *ListCategoriesRequest, out *ListCategoriesResponse) error {
	return h.ExploreHandler.ListCategories(ctx, in, out)
}

func (h *exploreHandler) Pricing(ctx context.Context, in *PricingRequest, out *PricingResponse) error {
	return h.ExploreHandler.Pricing(ctx, in, out)
}