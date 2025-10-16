package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloWorldHandler(t *testing.T) {
	tests := []struct {
		name       string
		query      string
		wantStatus int
		wantKey    string
		wantValue  string
	}{
		{"A-M uppercase", "?name=Alice", http.StatusOK, "message", "Hello Alice"},
		{"A-M lowercase", "?name=bob", http.StatusOK, "message", "Hello bob"},
		{"N-Z uppercase", "?name=Zane", http.StatusBadRequest, "error", "Invalid Input"},
		{"N-Z lowercase", "?name=nina", http.StatusBadRequest, "error", "Invalid Input"},
		{"Missing name", "", http.StatusBadRequest, "error", "Invalid Input"},
		{"Empty name", "?name=", http.StatusBadRequest, "error", "Invalid Input"},
	}

	for _, tt := range tests {
		req := httptest.NewRequest(http.MethodGet, "/hello-world"+tt.query, nil)
		rec := httptest.NewRecorder()

		helloWorldHandler(rec, req)

		if rec.Code != tt.wantStatus {
			t.Errorf("%s: expected status %d, got %d", tt.name, tt.wantStatus, rec.Code)
		}

		var resp map[string]string
		json.Unmarshal(rec.Body.Bytes(), &resp)
		if val, ok := resp[tt.wantKey]; !ok || val != tt.wantValue {
			t.Errorf("%s: expected %s=%s, got %v", tt.name, tt.wantKey, tt.wantValue, resp)
		}
	}
}
