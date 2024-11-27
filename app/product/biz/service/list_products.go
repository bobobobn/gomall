package service

import (
	"context"
	"gomall/app/product/biz/dal/mysql"
	"gomall/app/product/biz/model"

	product "gomall.local/rpc_gen/kitex_gen/product"
)

type ListProductsService struct {
	ctx context.Context
} // NewListProductsService new ListProductsService
func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx}
}

// Run create note info
func (s *ListProductsService) Run(req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	// Finish your business logic.
	categoris, err := model.GetProductsByCategoryName(mysql.DB, req.CategoryName)
	if err != nil {
		return nil, err
	}
	resp = &product.ListProductsResp{}
	for _, c := range categoris.Products {
		p, err := model.GetProductById(mysql.DB, int64(c.ID))
		if err != nil {
			return nil, err
		}
		resp.Products = append(resp.Products, &product.Product{
			Id:          uint32(p.ID),
			Name:        p.Name,
			Price:       p.Price,
			Description: p.Description,
			Pic:         p.Picture,
		})
	}
	return
}
