package service

import (
	"context"
	"errors"
	"strconv"

	"github.com/binli2020/order_api/backend-challenge/internal/generated"
)

type OrderService interface {
	PlaceOrder(req generated.OrderReq) (*generated.Order, error)
}

type orderService struct {
	productService ProductService
	promoService   PromoService
	promoFiles     []string
	nextID         int64
}

func NewOrderService(
	ps ProductService,
	promoSvc PromoService,
	promoFiles []string,
) OrderService {
	return &orderService{
		productService: ps,
		promoService:   promoSvc,
		promoFiles:     promoFiles,
		nextID:         1,
	}
}

func (os *orderService) PlaceOrder(req generated.OrderReq) (*generated.Order, error) {
	if len(req.Items) == 0 {
		return nil, errors.New("order must contain at least one item")
	}

	// Validate promo code
	if req.CouponCode != nil && *req.CouponCode != "" {
		code := *req.CouponCode

		// must be between 8 and 10 characters
		if len(code) < 8 || len(code) > 10 {
			return nil, errors.New("promo code must be 8–10 characters long")
		}

		ctx := context.Background()

		matches, err := os.promoService.FindPromo(
			ctx,
			*req.CouponCode,
			os.promoFiles,
			2,
		)
		if err != nil {
			return nil, err
		}
		if len(matches) == 0 {
			return nil, errors.New("invalid promo code")
		}
	}

	// Build order items and associated products
	orderItems := make([]struct {
		ProductId *string `json:"productId,omitempty"`
		Quantity  *int    `json:"quantity,omitempty"`
	}, 0, len(req.Items))

	products := make([]generated.Product, 0, len(req.Items))

	for _, item := range req.Items {
		if item.Quantity <= 0 {
			return nil, errors.New("quantity must be positive")
		}

		idNum, err := strconv.ParseInt(item.ProductId, 10, 64)
		if err != nil {
			return nil, errors.New("invalid productId")
		}

		prod, ok := os.productService.GetProductByID(idNum)
		if !ok {
			return nil, errors.New("product not found: " + item.ProductId)
		}

		// stable pointers
		pid := item.ProductId
		qty := item.Quantity

		orderItems = append(orderItems, struct {
			ProductId *string `json:"productId,omitempty"`
			Quantity  *int    `json:"quantity,omitempty"`
		}{
			ProductId: &pid,
			Quantity:  &qty,
		})

		products = append(products, *prod)
	}

	// Generate order ID
	idStr := strconv.FormatInt(os.nextID, 10)
	os.nextID++

	order := &generated.Order{
		Id:       &idStr,
		Items:    &orderItems,
		Products: &products,
	}

	return order, nil
}
