package service

import (
	"context"
	"testing"

	"github.com/binli2020/order_api/backend-challenge/internal/generated"
	"github.com/binli2020/order_api/backend-challenge/internal/utils"
)

// --- Mock ProductService ---
type mockProductService struct {
	product *generated.Product
	exists  bool
}

func (m mockProductService) GetAllProducts() []generated.Product {
	return []generated.Product{}
}
func (m mockProductService) GetProductByID(id int64) (*generated.Product, bool) {
	return m.product, m.exists
}

// --- Mock PromoService ---
type mockPromoService struct {
	matches int
	err     error
}

func (m mockPromoService) FindPromo(ctx context.Context, code string, files []string, maxMatches int) ([]PromoLocation, error) {
	if m.err != nil {
		return nil, m.err
	}
	list := []PromoLocation{}
	for i := 0; i < m.matches; i++ {
		list = append(list, PromoLocation{File: "f", Line: i})
	}
	return list, nil
}

func TestOrderService_EmptyItems(t *testing.T) {
	os := NewOrderService(mockProductService{}, mockPromoService{}, nil)

	req := generated.OrderReq{}
	_, err := os.PlaceOrder(req)
	if err == nil {
		t.Fatalf("expected error for empty items")
	}
}

func TestOrderService_InvalidQuantity(t *testing.T) {
	os := NewOrderService(mockProductService{}, mockPromoService{}, nil)

	req := generated.OrderReq{
		Items: []struct {
			ProductId string `json:"productId"`
			Quantity  int    `json:"quantity"`
		}{
			{ProductId: "1", Quantity: -1},
		},
	}

	_, err := os.PlaceOrder(req)
	if err == nil || err.Error() != "quantity must be positive" {
		t.Fatalf("expected quantity positive error, got %v", err)
	}
}

func TestOrderService_InvalidProductID_NotFound(t *testing.T) {
	os := NewOrderService(
		mockProductService{exists: false},
		mockPromoService{},
		nil,
	)

	req := generated.OrderReq{
		Items: []struct {
			ProductId string `json:"productId"`
			Quantity  int    `json:"quantity"`
		}{
			{ProductId: "99", Quantity: 1},
		},
	}

	_, err := os.PlaceOrder(req)
	if err == nil || err.Error() == "" {
		t.Fatalf("expected 'product not found' error")
	}
}

func TestOrderService_InvalidPromoCode_Length(t *testing.T) {
	os := NewOrderService(mockProductService{exists: true}, mockPromoService{}, nil)

	code := "ABC" // <8 chars
	req := generated.OrderReq{
		CouponCode: &code,
		Items: []struct {
			ProductId string `json:"productId"`
			Quantity  int    `json:"quantity"`
		}{
			{ProductId: "1", Quantity: 1},
		},
	}

	_, err := os.PlaceOrder(req)
	if err == nil {
		t.Fatalf("expected promo length error")
	}
}

func TestOrderService_InvalidPromoCode_NotFound(t *testing.T) {
	os := NewOrderService(
		mockProductService{exists: true},
		mockPromoService{matches: 0}, // no matches
		nil,
	)

	code := "VALID123"
	req := generated.OrderReq{
		CouponCode: &code,
		Items: []struct {
			ProductId string `json:"productId"`
			Quantity  int    `json:"quantity"`
		}{
			{ProductId: "1", Quantity: 1},
		},
	}

	_, err := os.PlaceOrder(req)
	if err == nil || err.Error() != "invalid promo code" {
		t.Fatalf("expected invalid promo code error")
	}
}

func TestOrderService_Success(t *testing.T) {
	os := NewOrderService(
		mockProductService{
			product: &generated.Product{
				Id:       utils.StrPtr("1"),
				Name:     utils.StrPtr("Test"),
				Price:    utils.Float32Ptr(10),
				Category: utils.StrPtr("Food"),
			},
			exists: true,
		},
		mockPromoService{matches: 2}, // valid promo (found >=2 files)
		nil,
	)

	code := "DISCOUNT!"
	req := generated.OrderReq{
		CouponCode: &code,
		Items: []struct {
			ProductId string `json:"productId"`
			Quantity  int    `json:"quantity"`
		}{
			{ProductId: "1", Quantity: 2},
		},
	}

	order, err := os.PlaceOrder(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if order.Id == nil {
		t.Fatalf("expected order id to be set")
	}
	if len(*order.Items) != 1 {
		t.Fatalf("expected 1 item")
	}
}
