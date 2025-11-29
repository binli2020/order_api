package controller

import (
	"net/http"
)

type ProductController struct{}

func NewProductController() *ProductController {
	return &ProductController{}
}

// GET /product
func (pc *ProductController) ListProducts(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// GET /product/{productId}
func (pc *ProductController) GetProduct(w http.ResponseWriter, r *http.Request, productId int64) {
	w.WriteHeader(http.StatusNotImplemented)
}
