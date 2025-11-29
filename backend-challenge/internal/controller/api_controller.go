package controller

import (
	"net/http"

	"github.com/binli2020/order_api/backend-challenge/internal/generated"
	"github.com/binli2020/order_api/backend-challenge/internal/service"
)

type APIController struct {
	ProductCtrl *ProductController
	OrderCtrl   *OrderController
}

var _ generated.ServerInterface = (*APIController)(nil)

func NewAPIController(productService *service.ProductService, orderService *service.OrderService) *APIController {
	return &APIController{
		ProductCtrl: NewProductController(*productService),
		OrderCtrl:   NewOrderController(*orderService),
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
