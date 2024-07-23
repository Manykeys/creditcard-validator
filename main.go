package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func luhnCheck(cardNumber string) bool {
	var sum int
	nDigits := len(cardNumber)
	parity := nDigits % 2

	for i, digit := range cardNumber {
		d, err := strconv.Atoi(string(digit))
		if err != nil {
			return false
		}

		if i%2 == parity {
			d *= 2
		}

		if d > 9 {
			d -= 9
		}

		sum += d
	}

	return sum%10 == 0
}

func isValidCardNumber(cardNumber string) bool {
	cardNumber = strings.ReplaceAll(cardNumber, " ", "")
	cardNumber = strings.ReplaceAll(cardNumber, "-", "")

	if len(cardNumber) < 13 || len(cardNumber) > 19 {
		return false
	}

	for _, ch := range cardNumber {
		if ch < '0' || ch > '9' {
			return false
		}
	}

	return luhnCheck(cardNumber)
}

func validateCardHandler(w http.ResponseWriter, r *http.Request) {
	cardNumber := r.URL.Query().Get("cardNumber")
	if cardNumber == "" {
		http.Error(w, "cardNumber parameter is required", http.StatusBadRequest)
		return
	}

	valid := isValidCardNumber(cardNumber)
	response := map[string]bool{"valid": valid}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/validateCard", validateCardHandler)
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
