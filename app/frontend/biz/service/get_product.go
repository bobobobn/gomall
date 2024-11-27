package service

import (
	"context"

	product "gomall/app/frontend/hertz_gen/frontend/product"
	"gomall/app/frontend/infra/rpc"

	rpcproduct "gomall.local/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/hertz/pkg/app"
)

type GetProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetProductService(Context context.Context, RequestContext *app.RequestContext) *GetProductService {
	return &GetProductService{RequestContext: RequestContext, Context: Context}
}

func (h *GetProductService) Run(req *product.ProductReq) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	r, err := rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductReq{Id: int64(req.Id)})
	if err != nil {
		return nil, err
	}
	resp = map[string]any{
		"item": r.Product,
	}
	return
}
