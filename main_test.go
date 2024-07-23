package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestValidateCardHandler(t *testing.T) {
	tests := []struct {
		cardNumber string
		expected   bool
	}{
		{"4539148803436467", true},  // Valid Visa
		{"6011111111111117", true},  // Valid Discover
		{"1234567890123456", false}, // Invalid
	}

	for _, test := range tests {
		req, err := http.NewRequest("GET", "/validateCard?cardNumber="+test.cardNumber, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(validateCardHandler)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		body, err := ioutil.ReadAll(rr.Body)
		if err != nil {
			t.Fatal(err)
		}

		var response map[string]bool
		if err := json.Unmarshal(body, &response); err != nil {
			t.Fatal(err)
		}

		if response["valid"] != test.expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				response["valid"], test.expected)
		}
	}
}
