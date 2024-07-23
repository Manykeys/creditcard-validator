package main

import (
	"fmt"
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

func main() {
	cardNumbers := []string{
		"4539 1488 0343 6467",
		"6011 1111 1111 1117",
		"1234 5678 9012 3456",
	}

	for _, cardNumber := range cardNumbers {
		if isValidCardNumber(cardNumber) {
			fmt.Printf("The card number %s is valid.\n", cardNumber)
		} else {
			fmt.Printf("The card number %s is invalid.\n", cardNumber)
		}
	}
}
