// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: financial_manage/financial_manage.proto

package financial_manage

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/asim/go-micro/v3/api"
	client "github.com/asim/go-micro/v3/client"
	server "github.com/asim/go-micro/v3/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for FinancialService service

func NewFinancialServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for FinancialService service

type FinancialService interface {
	GetAssetClass(ctx context.Context, in *GetAssetClassRequest, opts ...client.CallOption) (*GetAssetClassResponse, error)
	AddAssetClass(ctx context.Context, in *AddAssetClassRequest, opts ...client.CallOption) (*AddAssetClassResponse, error)
	ModifyAssetClass(ctx context.Context, in *ModifyAssetClassRequest, opts ...client.CallOption) (*ModifyAssetClassResponse, error)
	DeleteAssetClass(ctx context.Context, in *DeleteAssetClassRequest, opts ...client.CallOption) (*DeleteAssetClassResponse, error)
	GetAssetObject(ctx context.Context, in *GetAssetObjectRequest, opts ...client.CallOption) (*GetAssetObjectResponse, error)
	AddAssetObject(ctx context.Context, in *AddAssetObjectRequest, opts ...client.CallOption) (*AddAssetObjectResponse, error)
	ModifyAssetObject(ctx context.Context, in *ModifyAssetObjectRequest, opts ...client.CallOption) (*ModifyAssetObjectResponse, error)
	DeleteAssetObject(ctx context.Context, in *DeleteAssetObjectRequest, opts ...client.CallOption) (*DeleteAssetObjectResponse, error)
	GetIncomeExpenditureClass(ctx context.Context, in *GetIncomeExpenditureClassRequest, opts ...client.CallOption) (*GetIncomeExpenditureClassResponse, error)
	AddIncomeExpenditureClass(ctx context.Context, in *AddIncomeExpenditureClassRequest, opts ...client.CallOption) (*AddIncomeExpenditureClassResponse, error)
	ModifyIncomeExpenditureClass(ctx context.Context, in *ModifyIncomeExpenditureClassRequest, opts ...client.CallOption) (*ModifyIncomeExpenditureClassResponse, error)
	DeleteIncomeExpenditureClass(ctx context.Context, in *DeleteIncomeExpenditureClassRequest, opts ...client.CallOption) (*DeleteIncomeExpenditureClassResponse, error)
	GetIncomeExpenditureFlow(ctx context.Context, in *GetIncomeExpenditureFlowRequest, opts ...client.CallOption) (*GetIncomeExpenditureFlowResponse, error)
	AddIncomeExpenditureFlow(ctx context.Context, in *AddIncomeExpenditureFlowRequest, opts ...client.CallOption) (*AddIncomeExpenditureFlowResponse, error)
	ModifyIncomeExpenditureFlow(ctx context.Context, in *ModifyIncomeExpenditureFlowRequest, opts ...client.CallOption) (*ModifyIncomeExpenditureFlowResponse, error)
	DeleteIncomeExpenditureFlow(ctx context.Context, in *DeleteIncomeExpenditureFlowRequest, opts ...client.CallOption) (*DeleteIncomeExpenditureFlowResponse, error)
	GetTotalAsset(ctx context.Context, in *GetTotalAssetRequest, opts ...client.CallOption) (*GetTotalAssetResponse, error)
}

type financialService struct {
	c    client.Client
	name string
}

func NewFinancialService(name string, c client.Client) FinancialService {
	return &financialService{
		c:    c,
		name: name,
	}
}

func (c *financialService) GetAssetClass(ctx context.Context, in *GetAssetClassRequest, opts ...client.CallOption) (*GetAssetClassResponse, error) {
	req := c.c.NewRequest(c.name, "FinancialService.GetAssetClass", in)
	out := new(GetAssetClassResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financialService) AddAssetClass(ctx context.Context, in *AddAssetClassRequest, opts ...client.CallOption) (*AddAssetClassResponse, error) {
	req := c.c.NewRequest(c.name, "FinancialService.AddAssetClass", in)
	out := new(AddAssetClassResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financialService) ModifyAssetClass(ctx context.Context, in *ModifyAssetClassRequest, opts ...client.CallOption) (*ModifyAssetClassResponse, error) {
	req := c.c.NewRequest(c.name, "FinancialService.ModifyAssetClass", in)
	out := new(ModifyAssetClassResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financialService) DeleteAssetClass(ctx context.Context, in *DeleteAssetClassRequest, opts ...client.CallOption) (*DeleteAssetClassResponse, error) {
	req := c.c.NewRequest(c.name, "FinancialService.DeleteAssetClass", in)
	out := new(DeleteAssetClassResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financialService) GetAssetObject(ctx context.Context, in *GetAssetObjectRequest, opts ...client.CallOption) (*GetAssetObjectResponse, error) {
	req := c.c.NewRequest(c.name, "FinancialService.GetAssetObject", in)
	out := new(GetAssetObjectResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financialService) AddAssetObject(ctx context.Context, in *AddAssetObjectRequest, opts ...client.CallOption) (*AddAssetObjectResponse, error) {
	req := c.c.NewRequest(c.name, "FinancialService.AddAssetObject", in)
	out := new(AddAssetObjectResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financialService) ModifyAssetObject(ctx context.Context, in *ModifyAssetObjectRequest, opts ...client.CallOption) (*ModifyAssetObjectResponse, error) {
	req := c.c.NewRequest(c.name, "FinancialService.ModifyAssetObject", in)
	out := new(ModifyAssetObjectResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financialService) DeleteAssetObject(ctx context.Context, in *DeleteAssetObjectRequest, opts ...client.CallOption) (*DeleteAssetObjectResponse, error) {
	req := c.c.NewRequest(c.name, "FinancialService.DeleteAssetObject", in)
	out := new(DeleteAssetObjectResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financialService) GetIncomeExpenditureClass(ctx context.Context, in *GetIncomeExpenditureClassRequest, opts ...client.CallOption) (*GetIncomeExpenditureClassResponse, error) {
	req := c.c.NewRequest(c.name, "FinancialService.GetIncomeExpenditureClass", in)
	out := new(GetIncomeExpenditureClassResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financialService) AddIncomeExpenditureClass(ctx context.Context, in *AddIncomeExpenditureClassRequest, opts ...client.CallOption) (*AddIncomeExpenditureClassResponse, error) {
	req := c.c.NewRequest(c.name, "FinancialService.AddIncomeExpenditureClass", in)
	out := new(AddIncomeExpenditureClassResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financialService) ModifyIncomeExpenditureClass(ctx context.Context, in *ModifyIncomeExpenditureClassRequest, opts ...client.CallOption) (*ModifyIncomeExpenditureClassResponse, error) {
	req := c.c.NewRequest(c.name, "FinancialService.ModifyIncomeExpenditureClass", in)
	out := new(ModifyIncomeExpenditureClassResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financialService) DeleteIncomeExpenditureClass(ctx context.Context, in *DeleteIncomeExpenditureClassRequest, opts ...client.CallOption) (*DeleteIncomeExpenditureClassResponse, error) {
	req := c.c.NewRequest(c.name, "FinancialService.DeleteIncomeExpenditureClass", in)
	out := new(DeleteIncomeExpenditureClassResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financialService) GetIncomeExpenditureFlow(ctx context.Context, in *GetIncomeExpenditureFlowRequest, opts ...client.CallOption) (*GetIncomeExpenditureFlowResponse, error) {
	req := c.c.NewRequest(c.name, "FinancialService.GetIncomeExpenditureFlow", in)
	out := new(GetIncomeExpenditureFlowResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financialService) AddIncomeExpenditureFlow(ctx context.Context, in *AddIncomeExpenditureFlowRequest, opts ...client.CallOption) (*AddIncomeExpenditureFlowResponse, error) {
	req := c.c.NewRequest(c.name, "FinancialService.AddIncomeExpenditureFlow", in)
	out := new(AddIncomeExpenditureFlowResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financialService) ModifyIncomeExpenditureFlow(ctx context.Context, in *ModifyIncomeExpenditureFlowRequest, opts ...client.CallOption) (*ModifyIncomeExpenditureFlowResponse, error) {
	req := c.c.NewRequest(c.name, "FinancialService.ModifyIncomeExpenditureFlow", in)
	out := new(ModifyIncomeExpenditureFlowResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financialService) DeleteIncomeExpenditureFlow(ctx context.Context, in *DeleteIncomeExpenditureFlowRequest, opts ...client.CallOption) (*DeleteIncomeExpenditureFlowResponse, error) {
	req := c.c.NewRequest(c.name, "FinancialService.DeleteIncomeExpenditureFlow", in)
	out := new(DeleteIncomeExpenditureFlowResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financialService) GetTotalAsset(ctx context.Context, in *GetTotalAssetRequest, opts ...client.CallOption) (*GetTotalAssetResponse, error) {
	req := c.c.NewRequest(c.name, "FinancialService.GetTotalAsset", in)
	out := new(GetTotalAssetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FinancialService service

type FinancialServiceHandler interface {
	GetAssetClass(context.Context, *GetAssetClassRequest, *GetAssetClassResponse) error
	AddAssetClass(context.Context, *AddAssetClassRequest, *AddAssetClassResponse) error
	ModifyAssetClass(context.Context, *ModifyAssetClassRequest, *ModifyAssetClassResponse) error
	DeleteAssetClass(context.Context, *DeleteAssetClassRequest, *DeleteAssetClassResponse) error
	GetAssetObject(context.Context, *GetAssetObjectRequest, *GetAssetObjectResponse) error
	AddAssetObject(context.Context, *AddAssetObjectRequest, *AddAssetObjectResponse) error
	ModifyAssetObject(context.Context, *ModifyAssetObjectRequest, *ModifyAssetObjectResponse) error
	DeleteAssetObject(context.Context, *DeleteAssetObjectRequest, *DeleteAssetObjectResponse) error
	GetIncomeExpenditureClass(context.Context, *GetIncomeExpenditureClassRequest, *GetIncomeExpenditureClassResponse) error
	AddIncomeExpenditureClass(context.Context, *AddIncomeExpenditureClassRequest, *AddIncomeExpenditureClassResponse) error
	ModifyIncomeExpenditureClass(context.Context, *ModifyIncomeExpenditureClassRequest, *ModifyIncomeExpenditureClassResponse) error
	DeleteIncomeExpenditureClass(context.Context, *DeleteIncomeExpenditureClassRequest, *DeleteIncomeExpenditureClassResponse) error
	GetIncomeExpenditureFlow(context.Context, *GetIncomeExpenditureFlowRequest, *GetIncomeExpenditureFlowResponse) error
	AddIncomeExpenditureFlow(context.Context, *AddIncomeExpenditureFlowRequest, *AddIncomeExpenditureFlowResponse) error
	ModifyIncomeExpenditureFlow(context.Context, *ModifyIncomeExpenditureFlowRequest, *ModifyIncomeExpenditureFlowResponse) error
	DeleteIncomeExpenditureFlow(context.Context, *DeleteIncomeExpenditureFlowRequest, *DeleteIncomeExpenditureFlowResponse) error
	GetTotalAsset(context.Context, *GetTotalAssetRequest, *GetTotalAssetResponse) error
}

func RegisterFinancialServiceHandler(s server.Server, hdlr FinancialServiceHandler, opts ...server.HandlerOption) error {
	type financialService interface {
		GetAssetClass(ctx context.Context, in *GetAssetClassRequest, out *GetAssetClassResponse) error
		AddAssetClass(ctx context.Context, in *AddAssetClassRequest, out *AddAssetClassResponse) error
		ModifyAssetClass(ctx context.Context, in *ModifyAssetClassRequest, out *ModifyAssetClassResponse) error
		DeleteAssetClass(ctx context.Context, in *DeleteAssetClassRequest, out *DeleteAssetClassResponse) error
		GetAssetObject(ctx context.Context, in *GetAssetObjectRequest, out *GetAssetObjectResponse) error
		AddAssetObject(ctx context.Context, in *AddAssetObjectRequest, out *AddAssetObjectResponse) error
		ModifyAssetObject(ctx context.Context, in *ModifyAssetObjectRequest, out *ModifyAssetObjectResponse) error
		DeleteAssetObject(ctx context.Context, in *DeleteAssetObjectRequest, out *DeleteAssetObjectResponse) error
		GetIncomeExpenditureClass(ctx context.Context, in *GetIncomeExpenditureClassRequest, out *GetIncomeExpenditureClassResponse) error
		AddIncomeExpenditureClass(ctx context.Context, in *AddIncomeExpenditureClassRequest, out *AddIncomeExpenditureClassResponse) error
		ModifyIncomeExpenditureClass(ctx context.Context, in *ModifyIncomeExpenditureClassRequest, out *ModifyIncomeExpenditureClassResponse) error
		DeleteIncomeExpenditureClass(ctx context.Context, in *DeleteIncomeExpenditureClassRequest, out *DeleteIncomeExpenditureClassResponse) error
		GetIncomeExpenditureFlow(ctx context.Context, in *GetIncomeExpenditureFlowRequest, out *GetIncomeExpenditureFlowResponse) error
		AddIncomeExpenditureFlow(ctx context.Context, in *AddIncomeExpenditureFlowRequest, out *AddIncomeExpenditureFlowResponse) error
		ModifyIncomeExpenditureFlow(ctx context.Context, in *ModifyIncomeExpenditureFlowRequest, out *ModifyIncomeExpenditureFlowResponse) error
		DeleteIncomeExpenditureFlow(ctx context.Context, in *DeleteIncomeExpenditureFlowRequest, out *DeleteIncomeExpenditureFlowResponse) error
		GetTotalAsset(ctx context.Context, in *GetTotalAssetRequest, out *GetTotalAssetResponse) error
	}
	type FinancialService struct {
		financialService
	}
	h := &financialServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&FinancialService{h}, opts...))
}

type financialServiceHandler struct {
	FinancialServiceHandler
}

func (h *financialServiceHandler) GetAssetClass(ctx context.Context, in *GetAssetClassRequest, out *GetAssetClassResponse) error {
	return h.FinancialServiceHandler.GetAssetClass(ctx, in, out)
}

func (h *financialServiceHandler) AddAssetClass(ctx context.Context, in *AddAssetClassRequest, out *AddAssetClassResponse) error {
	return h.FinancialServiceHandler.AddAssetClass(ctx, in, out)
}

func (h *financialServiceHandler) ModifyAssetClass(ctx context.Context, in *ModifyAssetClassRequest, out *ModifyAssetClassResponse) error {
	return h.FinancialServiceHandler.ModifyAssetClass(ctx, in, out)
}

func (h *financialServiceHandler) DeleteAssetClass(ctx context.Context, in *DeleteAssetClassRequest, out *DeleteAssetClassResponse) error {
	return h.FinancialServiceHandler.DeleteAssetClass(ctx, in, out)
}

func (h *financialServiceHandler) GetAssetObject(ctx context.Context, in *GetAssetObjectRequest, out *GetAssetObjectResponse) error {
	return h.FinancialServiceHandler.GetAssetObject(ctx, in, out)
}

func (h *financialServiceHandler) AddAssetObject(ctx context.Context, in *AddAssetObjectRequest, out *AddAssetObjectResponse) error {
	return h.FinancialServiceHandler.AddAssetObject(ctx, in, out)
}

func (h *financialServiceHandler) ModifyAssetObject(ctx context.Context, in *ModifyAssetObjectRequest, out *ModifyAssetObjectResponse) error {
	return h.FinancialServiceHandler.ModifyAssetObject(ctx, in, out)
}

func (h *financialServiceHandler) DeleteAssetObject(ctx context.Context, in *DeleteAssetObjectRequest, out *DeleteAssetObjectResponse) error {
	return h.FinancialServiceHandler.DeleteAssetObject(ctx, in, out)
}

func (h *financialServiceHandler) GetIncomeExpenditureClass(ctx context.Context, in *GetIncomeExpenditureClassRequest, out *GetIncomeExpenditureClassResponse) error {
	return h.FinancialServiceHandler.GetIncomeExpenditureClass(ctx, in, out)
}

func (h *financialServiceHandler) AddIncomeExpenditureClass(ctx context.Context, in *AddIncomeExpenditureClassRequest, out *AddIncomeExpenditureClassResponse) error {
	return h.FinancialServiceHandler.AddIncomeExpenditureClass(ctx, in, out)
}

func (h *financialServiceHandler) ModifyIncomeExpenditureClass(ctx context.Context, in *ModifyIncomeExpenditureClassRequest, out *ModifyIncomeExpenditureClassResponse) error {
	return h.FinancialServiceHandler.ModifyIncomeExpenditureClass(ctx, in, out)
}

func (h *financialServiceHandler) DeleteIncomeExpenditureClass(ctx context.Context, in *DeleteIncomeExpenditureClassRequest, out *DeleteIncomeExpenditureClassResponse) error {
	return h.FinancialServiceHandler.DeleteIncomeExpenditureClass(ctx, in, out)
}

func (h *financialServiceHandler) GetIncomeExpenditureFlow(ctx context.Context, in *GetIncomeExpenditureFlowRequest, out *GetIncomeExpenditureFlowResponse) error {
	return h.FinancialServiceHandler.GetIncomeExpenditureFlow(ctx, in, out)
}

func (h *financialServiceHandler) AddIncomeExpenditureFlow(ctx context.Context, in *AddIncomeExpenditureFlowRequest, out *AddIncomeExpenditureFlowResponse) error {
	return h.FinancialServiceHandler.AddIncomeExpenditureFlow(ctx, in, out)
}

func (h *financialServiceHandler) ModifyIncomeExpenditureFlow(ctx context.Context, in *ModifyIncomeExpenditureFlowRequest, out *ModifyIncomeExpenditureFlowResponse) error {
	return h.FinancialServiceHandler.ModifyIncomeExpenditureFlow(ctx, in, out)
}

func (h *financialServiceHandler) DeleteIncomeExpenditureFlow(ctx context.Context, in *DeleteIncomeExpenditureFlowRequest, out *DeleteIncomeExpenditureFlowResponse) error {
	return h.FinancialServiceHandler.DeleteIncomeExpenditureFlow(ctx, in, out)
}

func (h *financialServiceHandler) GetTotalAsset(ctx context.Context, in *GetTotalAssetRequest, out *GetTotalAssetResponse) error {
	return h.FinancialServiceHandler.GetTotalAsset(ctx, in, out)
}
