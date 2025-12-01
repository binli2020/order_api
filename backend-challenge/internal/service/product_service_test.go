package service

import (
	"testing"
)

func TestProductService_GetAllProducts(t *testing.T) {
	ps := NewProductService()

	products := ps.GetAllProducts()
	if len(products) != 3 {
		t.Fatalf("expected 3 products, got %d", len(products))
	}
}

func TestProductService_GetProductByID_Found(t *testing.T) {
	ps := NewProductService()

	prod, ok := ps.GetProductByID(1)
	if !ok {
		t.Fatalf("expected product with ID=1 to be found")
	}

	if prod.Id == nil || *prod.Id != "1" {
		t.Fatalf("expected product ID '1', got %#v", prod.Id)
	}
	if prod.Name == nil || *prod.Name == "" {
		t.Fatalf("expected non-nil product name")
	}
}

func TestProductService_GetProductByID_NotFound(t *testing.T) {
	ps := NewProductService()

	_, ok := ps.GetProductByID(999)
	if ok {
		t.Fatalf("expected product not found")
	}
}
