package controller

import (
	"net/http"

	"github.com/binli2020/order_api/backend-challenge/internal/service"
)

type OrderController struct {
	orderService service.OrderService
}

func NewOrderController(orderService service.OrderService) *OrderController {
	return &OrderController{orderService: orderService}
}

// POST /order
func (oc *OrderController) PlaceOrder(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
