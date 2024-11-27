package service

import (
	"context"

	category "gomall/app/frontend/hertz_gen/frontend/category"
	"gomall/app/frontend/infra/rpc"

	"github.com/cloudwego/hertz/pkg/app"
	"gomall.local/rpc_gen/kitex_gen/product"
)

type CategoryService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCategoryService(Context context.Context, RequestContext *app.RequestContext) *CategoryService {
	return &CategoryService{RequestContext: RequestContext, Context: Context}
}

func (h *CategoryService) Run(req *category.CategoryReq) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	list_products_req := product.ListProductsReq{
		CategoryName: req.Category,
	}
	r, err := rpc.ProductClient.ListProducts(context.Background(), &list_products_req)
	if err != nil {
		return nil, err
	}
	resp = map[string]any{
		"title": "Category",
		"items": r.Products,
	}

	return
}
