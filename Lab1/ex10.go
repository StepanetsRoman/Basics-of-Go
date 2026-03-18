package main

import "fmt"

func main() {
	var chartype int8 = 'R'

	fmt.Printf("Code '%c' - %d\n", chartype, chartype)

	// Завдання.
	// 1. Вивести українську літеру 'Ї'
	var ukrainianLetter rune = 'Ї'
	fmt.Printf("Code '%c' - %d\n", ukrainianLetter, ukrainianLetter)
	
	// 2. Пояснити призначення типу "rune"
	fmt.Println("\nПояснення типу 'rune':")
	fmt.Println("Тип 'rune' є синонімом типу int32 і використовується для представлення")
	fmt.Println("кодових точок Unicode. Він дозволяє коректно працювати з символами")
	fmt.Println("різних мов, включаючи українські літери, які потребують більше 1 байта.")
	fmt.Println("rune забезпечує коректну роботу з багатобайтовими символами UTF-8.")
}

