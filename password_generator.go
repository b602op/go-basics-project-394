package main

import (
	"fmt"
	"strings"
)

const (
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits    = "0123456789"
	special   = "!@#$%^&*"
)

const allChars = lowercase + uppercase + digits + special

func NextRandom(number int) int {
	return (16807 * number) % 2147483647
}

func GeneratePassword(length int, seed int, useUppercase, useDigits, useSpecial bool) string {
	if length <= 0 {
		return ""
	}
	currentAllChars := lowercase

	if useUppercase {
		currentAllChars += uppercase
	}

	if useDigits {
		currentAllChars += digits
	}

	if useSpecial {
		currentAllChars += special
	}

	b := make([]byte, length)

	for i := range b {
		b[i] = currentAllChars[NextRandom(seed+i)%len(currentAllChars)]
	}

	return string(b)
}

func CheckPassword(password string) string {
	hasMinLength := len(password) >= 8
	hasLower := false
	hasUpper := false
	hasDigit := false
	hasSpecial := false

	for i := 0; i < len(password); i++ {
		c := password[i]

		switch {
		case strings.IndexByte(lowercase, c) >= 0:
			hasLower = true
		case strings.IndexByte(uppercase, c) >= 0:
			hasUpper = true
		case strings.IndexByte(digits, c) >= 0:
			hasDigit = true
		case strings.IndexByte(special, c) >= 0:
			hasSpecial = true
		}
	}

	score := 0
	if hasLower {
		score++
	}
	if hasUpper {
		score++
	}
	if hasDigit {
		score++
	}
	if hasSpecial {
		score++
	}
	if hasMinLength {
		score++
	}

	var verdict string

	switch score {
	case 5:
		verdict = "Очень надёжный"
	case 4:
		verdict = "Надёжный"
	case 3:
		verdict = "Средний"
	default:
		verdict = "Слабый"
	}

	return fmt.Sprintf("%s пароль (оценка %d из 5)", verdict, score)
}
