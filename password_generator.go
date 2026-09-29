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
	minLength = 8
)

// NextRandom возвращает следующее псевдослучайное число
// по линейному конгруэнтному генератору.
func NextRandom(number int) int {
	return (16807 * number) % 2147483647
}

// GeneratePassword генерирует пароль заданной длины по числу-ключу seed.
// Набор символов определяется флагами useUppercase, useDigits, useSpecial.
func GeneratePassword(length int, seed int, useUppercase, useDigits, useSpecial bool) string {
	alphabet := lowercase
	if useUppercase {
		alphabet += uppercase
	}
	if useDigits {
		alphabet += digits
	}
	if useSpecial {
		alphabet += special
	}

	var sb strings.Builder
	current := seed

	for i := 0; i < length; i++ {
		current = NextRandom(current)
		sb.WriteByte(alphabet[current%len(alphabet)])
	}

	return sb.String()
}

// has сообщает, содержит ли password хотя бы один символ из chars.
func has(password, chars string) bool {
	for _, c := range password {
		if strings.ContainsRune(chars, c) {
			return true
		}
	}
	return false
}

// passwordScore возвращает число выполненных критериев надёжности (0–5).
func passwordScore(password string) int {
	score := 0
	if has(password, lowercase) {
		score++
	}
	if has(password, uppercase) {
		score++
	}
	if has(password, digits) {
		score++
	}
	if has(password, special) {
		score++
	}
	if len(password) >= minLength {
		score++
	}
	return score
}

// CheckPassword оценивает надёжность пароля по пяти критериям
// и возвращает текстовый вердикт с оценкой.
func CheckPassword(password string) string {
	score := passwordScore(password)

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
