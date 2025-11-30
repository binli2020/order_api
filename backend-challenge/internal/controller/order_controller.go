package controller

import (
	"encoding/json"
	"net/http"

	"github.com/binli2020/order_api/backend-challenge/internal/generated"
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
	var req generated.PlaceOrderJSONRequestBody // alias for generated.OrderReq

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteJSONError(w, http.StatusBadRequest, ErrorTypeValidation, "invalid request body")
		return
	}

	order, err := oc.orderService.PlaceOrder(req)
	if err != nil {
		WriteJSONError(w, http.StatusBadRequest, ErrorTypeValidation, err.Error())
		return
	}

	WriteJSON(w, http.StatusOK, order)
}
