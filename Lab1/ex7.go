package main

import "fmt"

func main() {
	variable8 := int8(127)
	variable16 := int16(16383)

	fmt.Println("Приведення типів\n")

	fmt.Printf("variable8         = %-5d = %.16b\n", variable8, variable8)
	fmt.Printf("variable16        = %-5d = %.16b\n", variable16, variable16)
	fmt.Printf("uint16(variable8) = %-5d = %.16b\n", uint16(variable8), uint16(variable8))
	fmt.Printf("uint8(variable16) = %-5d = %.16b\n", uint8(variable16), uint8(variable16))

	// Завдання.
	// 1. Створіть 2 змінні різних типів. Виконайте арифметичні операції. Результат вивести в консоль
	var num1 int32 = 15
	var num2 float64 = 3.5
	
	fmt.Println("\nАрифметичні операції з різними типами:")
	fmt.Printf("num1 = %d (тип: %T)\n", num1, num1)
	fmt.Printf("num2 = %.2f (тип: %T)\n\n", num2, num2)
	
	// Приведення типів для арифметичних операцій
	fmt.Printf("float64(num1) + num2 = %.2f\n", float64(num1)+num2)
	fmt.Printf("float64(num1) - num2 = %.2f\n", float64(num1)-num2)
	fmt.Printf("float64(num1) * num2 = %.2f\n", float64(num1)*num2)
	fmt.Printf("float64(num1) / num2 = %.2f\n", float64(num1)/num2)
}

