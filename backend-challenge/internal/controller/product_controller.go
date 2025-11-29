package controller

import (
	"encoding/json"
	"net/http"

	"github.com/binli2020/order_api/backend-challenge/internal/service"
)

type ProductController struct {
	productService service.ProductService
}

func NewProductController(ps service.ProductService) *ProductController {
	return &ProductController{productService: ps}
}

// GET /product
func (pc *ProductController) ListProducts(w http.ResponseWriter, r *http.Request) {
	products := pc.productService.GetAllProducts()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

// GET /product/{productId}
func (pc *ProductController) GetProduct(w http.ResponseWriter, r *http.Request, productId int64) {
	product, ok := pc.productService.GetProductByID(productId)
	if !ok {
		http.Error(w, "product not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}
