package router

import (
	"github.com/go-chi/chi/v5"

	"github.com/binli2020/order_api/backend-challenge/internal/controller"
	"github.com/binli2020/order_api/backend-challenge/internal/generated"
)

func NewRouter() chi.Router {
	r := chi.NewRouter()

	productController := controller.NewProductController()
	orderController := controller.NewOrderController()

	api := controller.NewAPIController(productController, orderController)

	generated.HandlerFromMux(api, r)

	return r
}
