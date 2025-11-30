package router

import (
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"

	"github.com/binli2020/order_api/backend-challenge/internal/controller"
	"github.com/binli2020/order_api/backend-challenge/internal/generated"
	"github.com/binli2020/order_api/backend-challenge/internal/service"
)

func NewRouter() chi.Router {
	r := chi.NewRouter()

	productService := service.NewProductService()
	promoService := service.NewPromoService()

	promoDir := os.Getenv("PROMO_DIR")
	promoFiles := []string{}
	if promoDir == "" {
		promoFiles = []string{
			"data/couponbase1.txt",
			"data/couponbase2.txt",
			"data/couponbase3.txt",
		}
	} else {
		promoFiles = []string{
			filepath.Join(promoDir, "coupon1.txt"),
			filepath.Join(promoDir, "coupon2.txt"),
			filepath.Join(promoDir, "coupon3.txt"),
		}
	}

	orderService := service.NewOrderService(
		productService,
		promoService,
		promoFiles,
	)

	api := controller.NewAPIController(&productService, &orderService)

	generated.HandlerFromMux(api, r)

	return r
}
