package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/binli2020/order_api/backend-challenge/internal/router"
)

func TestPlaceOrder_Unimplemented(t *testing.T) {
	r := router.NewRouter()

	body := []byte(`{"items":[{"productId":"1","quantity":1}]}`)
	req := httptest.NewRequest(http.MethodPost, "/order", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotImplemented {
		t.Fatalf("expected 501 Not Implemented, got %d", w.Code)
	}
}
