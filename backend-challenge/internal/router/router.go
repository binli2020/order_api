package router

import (
	"github.com/go-chi/chi/v5"

	"github.com/binli2020/order_api/backend-challenge/internal/controller"
	"github.com/binli2020/order_api/backend-challenge/internal/generated"
	"github.com/binli2020/order_api/backend-challenge/internal/service"
)

func NewRouter() chi.Router {
	r := chi.NewRouter()

	productService := service.NewProductService()
	orderService := service.NewOrderService()

	api := controller.NewAPIController(&productService, &orderService)

	generated.HandlerFromMux(api, r)

	return r
}
