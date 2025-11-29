package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/binli2020/order_api/backend-challenge/internal/generated"
	"github.com/binli2020/order_api/backend-challenge/internal/router"
)

func TestGetProduct_OK(t *testing.T) {
	r := router.NewRouter()

	req := httptest.NewRequest(http.MethodGet, "/product/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	var product generated.Product
	if err := json.Unmarshal(w.Body.Bytes(), &product); err != nil {
		t.Fatalf("failed to parse product: %v", err)
	}

	if *product.Id != "1" {
		t.Fatalf("expected product ID 1, got %s", *product.Id)
	}
}

func TestGetProduct_NotFound(t *testing.T) {
	r := router.NewRouter()

	req := httptest.NewRequest(http.MethodGet, "/product/99", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("expected 404 for missing product, got %d", w.Code)
	}
}
