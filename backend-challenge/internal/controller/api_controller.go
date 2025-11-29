package controller

import (
	"net/http"

	"github.com/binli2020/order_api/backend-challenge/internal/generated"
)

type APIController struct {
	ProductCtrl *ProductController
	OrderCtrl   *OrderController
}

var _ generated.ServerInterface = (*APIController)(nil)

func NewAPIController(productController *ProductController, orderController *OrderController) *APIController {
	return &APIController{
		ProductCtrl: productController,
		OrderCtrl:   orderController,
	}
}

func (a *APIController) ListProducts(w http.ResponseWriter, r *http.Request) {
	a.ProductCtrl.ListProducts(w, r)
}

func (a *APIController) GetProduct(w http.ResponseWriter, r *http.Request, productId int64) {
	a.ProductCtrl.GetProduct(w, r, productId)
}

func (a *APIController) PlaceOrder(w http.ResponseWriter, r *http.Request) {
	a.OrderCtrl.PlaceOrder(w, r)
}
