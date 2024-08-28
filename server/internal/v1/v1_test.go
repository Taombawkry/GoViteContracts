package v1

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthzHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/healthz", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthzHandler)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := map[string]string{"status": "ok", "route": "v1"}
	var actual map[string]string
	body, _ := io.ReadAll(rr.Body)
	err = json.Unmarshal(body, &actual)
	if err != nil {
		t.Fatal("Could not unmarshal response:", err)
	}

	if actual["status"] != expected["status"] || actual["route"] != expected["route"] {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}
