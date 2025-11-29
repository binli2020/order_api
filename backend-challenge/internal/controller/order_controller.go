package controller

import (
	"net/http"
)

type OrderController struct{}

func NewOrderController() *OrderController {
	return &OrderController{}
}

// POST /order
func (oc *OrderController) PlaceOrder(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
