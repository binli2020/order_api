package service

import (
	"strconv"

	"github.com/binli2020/order_api/backend-challenge/internal/generated"
	"github.com/binli2020/order_api/backend-challenge/internal/utils"
)

type ProductService interface {
	GetAllProducts() []generated.Product
	GetProductByID(id int64) (*generated.Product, bool)
}

type productService struct {
	products []generated.Product
}

func NewProductService() ProductService {
	return &productService{
		products: []generated.Product{
			{
				Id:       utils.StrPtr("1"),
				Name:     utils.StrPtr("Dumplings"),
				Category: utils.StrPtr("Food"),
				Price:    utils.Float32Ptr(16.0),
			},
			{
				Id:       utils.StrPtr("2"),
				Name:     utils.StrPtr("Lasagna"),
				Category: utils.StrPtr("Food"),
				Price:    utils.Float32Ptr(18.0),
			},
			{
				Id:       utils.StrPtr("3"),
				Name:     utils.StrPtr("Chicken Burger"),
				Category: utils.StrPtr("Food"),
				Price:    utils.Float32Ptr(12.5),
			},
		},
	}
}

func (ps *productService) GetAllProducts() []generated.Product {
	return ps.products
}

func (ps *productService) GetProductByID(id int64) (*generated.Product, bool) {
	idStr := strconv.FormatInt(id, 10)
	for _, p := range ps.products {
		// generated.Product.Id is *string, convert id to string
		if p.Id != nil && *p.Id == idStr {
			return &p, true
		}
	}

	return nil, false
}
