package service

import (
	"context"
	"gomall/app/product/biz/dal/mysql"
	"gomall/app/product/biz/model"

	product "gomall.local/rpc_gen/kitex_gen/product"
)

type SearchProductsService struct {
	ctx context.Context
} // NewSearchProductsService new SearchProductsService
func NewSearchProductsService(ctx context.Context) *SearchProductsService {
	return &SearchProductsService{ctx: ctx}
}

// Run create note info
func (s *SearchProductsService) Run(req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	// Finish your business logic.
	p, err := model.SearchProduct(mysql.DB, req.Query)
	if err != nil {
		return nil, err
	}
	resp = &product.SearchProductsResp{}
	for _, v := range p {
		resp.Results = append(resp.Results, &product.Product{
			Id:          uint32(v.ID),
			Name:        v.Name,
			Price:       v.Price,
			Pic:         v.Picture,
			Description: v.Description,
		})
	}
	return
}
