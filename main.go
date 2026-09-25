package main

import "fmt"

func main() {
	fmt.Println("== Генерация паролей ==")
	fmt.Println("буквы и цифры:   ", GeneratePassword(12, 123, true, true, false))
	fmt.Println("со спецсимволами:", GeneratePassword(16, 7, true, true, true))

	fmt.Println()
	fmt.Println("== Проверка надёжности ==")
	fmt.Println("abc        ->", CheckPassword("abc"))
	fmt.Println("abcdef1234 ->", CheckPassword("abcdef1234"))
	fmt.Println("Abcdef123! ->", CheckPassword("Abcdef123!"))
}
