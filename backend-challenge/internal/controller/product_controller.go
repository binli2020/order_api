package controller

import (
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

	WriteJSON(w, http.StatusOK, products)
}

// GET /product/{productId}
func (pc *ProductController) GetProduct(w http.ResponseWriter, r *http.Request, productId int64) {
	product, ok := pc.productService.GetProductByID(productId)
	if !ok {
		WriteJSONError(w, http.StatusNotFound, ErrorTypeNotFound, "product not found")

		return
	}

	WriteJSON(w, http.StatusOK, product)
}
