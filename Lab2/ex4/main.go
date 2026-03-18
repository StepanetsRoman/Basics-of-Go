package main

import (
	"fmt"
	"ex4/calc"
)

func main() {
	// Тестування функції Min
	fmt.Println("Мінімальне значення з (5, 2, 8):", calc.Min(5, 2, 8))
	fmt.Println("Мінімальне значення з (-3, -1, -5):", calc.Min(-3, -1, -5))
	fmt.Println("Мінімальне значення з (10, 10, 10):", calc.Min(10, 10, 10))

	// Тестування функції Average
	fmt.Println("Середнє значення (5, 10, 15):", calc.Average(5, 10, 15))
	fmt.Println("Середнє значення (1, 2, 3):", calc.Average(1, 2, 3))
	fmt.Println("Середнє значення (-5, 0, 5):", calc.Average(-5, 0, 5))

	// Тестування функції SolveLinearEquation
	x1, err1 := calc.SolveLinearEquation(2, -4)
	if err1 != nil {
		fmt.Println("Помилка:", err1)
	} else {
		fmt.Println("Розв'язок рівняння 2x - 4 = 0: x =", x1)
	}

	x2, err2 := calc.SolveLinearEquation(3, 6)
	if err2 != nil {
		fmt.Println("Помилка:", err2)
	} else {
		fmt.Println("Розв'язок рівняння 3x + 6 = 0: x =", x2)
	}

	x3, err3 := calc.SolveLinearEquation(0, 5)
	if err3 != nil {
		fmt.Println("Помилка:", err3)
	} else {
		fmt.Println("Розв'язок рівняння 0x + 5 = 0: x =", x3)
	}
}
