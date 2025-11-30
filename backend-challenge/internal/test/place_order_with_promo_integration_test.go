package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/binli2020/order_api/backend-challenge/internal/generated"
	"github.com/binli2020/order_api/backend-challenge/internal/router"
)

func createPromoFilesForTest(t *testing.T) string {
	t.Helper()

	dir := t.TempDir()

	os.WriteFile(filepath.Join(dir, "coupon1.txt"), []byte("ABC\nPROMO123\nZZZ"), 0644)
	os.WriteFile(filepath.Join(dir, "coupon2.txt"), []byte("NO_MATCH"), 0644)
	os.WriteFile(filepath.Join(dir, "coupon3.txt"), []byte("ANOTHER PROMO123"), 0644)

	return dir
}

func TestPlaceOrder_WithValidPromo_OK(t *testing.T) {
	dir := createPromoFilesForTest(t)
	os.Setenv("PROMO_DIR", dir)
	defer os.Unsetenv("PROMO_DIR")

	r := router.NewRouter()

	body := []byte(`{
        "couponCode": "PROMO123",
        "items": [{"productId": "1", "quantity": 2}]
    }`)

	req := httptest.NewRequest("POST", "/order", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d, body=%s", w.Code, w.Body.String())
	}

	var order generated.Order
	if err := json.Unmarshal(w.Body.Bytes(), &order); err != nil {
		t.Fatalf("cannot parse order: %v", err)
	}

	if order.Id == nil || *order.Id == "" {
		t.Fatalf("expected order ID")
	}
}

func TestPlaceOrder_InvalidPromo_400(t *testing.T) {
	dir := createPromoFilesForTest(t)
	os.Setenv("PROMO_DIR", dir)
	defer os.Unsetenv("PROMO_DIR")

	r := router.NewRouter()

	body := []byte(`{
        "couponCode": "NOT_REAL",
        "items": [{"productId": "1", "quantity": 1}]
    }`)

	req := httptest.NewRequest("POST", "/order", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}

func TestPlaceOrder_NoPromo_OK(t *testing.T) {
	dir := createPromoFilesForTest(t)
	os.Setenv("PROMO_DIR", dir)
	defer os.Unsetenv("PROMO_DIR")

	r := router.NewRouter()

	body := []byte(`{
        "items": [{"productId":"1","quantity":1}]
    }`)

	req := httptest.NewRequest("POST", "/order", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
}
