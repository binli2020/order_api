package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/binli2020/order_api/backend-challenge/internal/router"
)

func TestListProducts_Unimplemented(t *testing.T) {
	r := router.NewRouter()

	req := httptest.NewRequest(http.MethodGet, "/product", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotImplemented {
		t.Fatalf("expected 501 Not Implemented, got %d", w.Code)
	}
}
