package handlers

import (
	"net/http/httptest"
	"testing"
)

func TestListProducts(t *testing.T) {
	rec := httptest.NewRecorder()
	ListProducts(rec, httptest.NewRequest("GET", "/api/products", nil))
	if rec.Code != 200 {
		t.Fatalf("status %d", rec.Code)
	}
}
