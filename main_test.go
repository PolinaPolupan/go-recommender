package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecommendationsHandler(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodGet,
		"/recommendations?user_id=42&limit=2",
		nil,
	)

	rec := httptest.NewRecorder()

	recommendationsHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	if rec.Header().Get("Content-Type") != "application/json" {
		t.Fatalf("expected JSON content type, got %q",
			rec.Header().Get("Content-Type"))
	}
}

func TestRecommendationsHandlerMissingUserID(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodGet,
		"/recommendations",
		nil,
	)

	rec := httptest.NewRecorder()

	recommendationsHandler(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rec.Code)
	}
}

func TestRecommendationsHandlerInvalidLimit(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodGet,
		"/recommendations?user_id=42&limit=abc",
		nil,
	)

	rec := httptest.NewRecorder()

	recommendationsHandler(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rec.Code)
	}
}
