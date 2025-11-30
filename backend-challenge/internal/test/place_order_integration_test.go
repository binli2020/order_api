package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/binli2020/order_api/backend-challenge/internal/generated"
	"github.com/binli2020/order_api/backend-challenge/internal/router"
)

func TestPlaceOrder_OK(t *testing.T) {
	r := router.NewRouter()

	body := []byte(`{"items":[{"productId":"1","quantity":2}]}`)
	req := httptest.NewRequest(http.MethodPost, "/order", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	var order generated.Order
	if err := json.Unmarshal(w.Body.Bytes(), &order); err != nil {
		t.Fatalf("failed to parse order response: %v", err)
	}

	if order.Id == nil || *order.Id == "" {
		t.Fatalf("expected non-empty order ID")
	}

	itemsCount := 0
	if order.Items != nil {
		itemsCount = len(*order.Items)
	}
	if itemsCount != 1 {
		t.Fatalf("expected 1 order item, got %d", itemsCount)
	}

	productsCount := 0
	if order.Products != nil {
		productsCount = len(*order.Products)
	}
	if productsCount != 1 {
		t.Fatalf("expected 1 product in order, got %d", productsCount)
	}

	if (*order.Products)[0].Id == nil || *(*order.Products)[0].Id != "1" {
		t.Fatalf("expected product ID 1 in order, got %+v", order.Products)
	}
}
