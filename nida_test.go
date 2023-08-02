package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCardGet(t *testing.T) {
	testCases := []struct {
		cardNumber       string
		expectedResponse string
	}{
		{"1234567890123456", "Response for card 1234567890123456"},
		{"9876543210987654", "Response for card 9876543210987654"},
	}

	for _, tc := range testCases {
		t.Run(tc.cardNumber, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodPost {
					t.Errorf("Expected POST request, got %s", r.Method)
					return
				}
				if r.URL.String() != fmt.Sprintf("/%s", tc.cardNumber) {
					t.Errorf("Expected URL: /%s, got %s", tc.cardNumber, r.URL.String())
					return
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				_, err := fmt.Fprintf(w, "Response for card %s", tc.cardNumber)
				if err != nil {
					return
				}
			}))
			defer server.Close()

			BaseUrl = server.URL

			card := &Card{Number: tc.cardNumber}
			response := card.Get()

			if response != tc.expectedResponse {
				t.Errorf("Expected response: %s, got %v", tc.expectedResponse, response)
			}
		})
	}
}
