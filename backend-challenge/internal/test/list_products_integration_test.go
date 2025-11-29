package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/binli2020/order_api/backend-challenge/internal/generated"
	"github.com/binli2020/order_api/backend-challenge/internal/router"
)

func TestListProducts_OK(t *testing.T) {
	r := router.NewRouter()

	req := httptest.NewRequest(http.MethodGet, "/product", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	var products []generated.Product
	if err := json.Unmarshal(w.Body.Bytes(), &products); err != nil {
		t.Fatalf("failed to parse response: %v", err)
	}

	if len(products) != 3 {
		t.Fatalf("expected 3 products, got %d", len(products))
	}

	if *products[0].Id != "1" || *products[0].Name != "Dumplings" || *products[0].Category != "Food" || *products[0].Price != 16.0 {
		t.Fatalf("expected first product ID to be 1, got %s", *products[0].Id)
	}
}
