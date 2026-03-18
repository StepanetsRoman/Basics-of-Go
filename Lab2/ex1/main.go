package main

import "fmt"

// Min знаходить мінімальне значення з трьох елементів
func Min(a, b, c float64) float64 {
	min := a
	if b < min {
		min = b
	}
	if c < min {
		min = c
	}
	return min
}

// Average обчислює середнє значення трьох елементів
func Average(a, b, c float64) float64 {
	return (a + b + c) / 3.0
}

// SolveLinearEquation розв'язує рівняння першого порядку ax + b = 0
// Повертає x = -b/a
func SolveLinearEquation(a, b float64) (float64, error) {
	if a == 0 {
		return 0, fmt.Errorf("ділення на нуль: коефіцієнт a не може бути рівним нулю")
	}
	return -b / a, nil
}

func main() {
	// Приклад використання функцій
	fmt.Println("Мінімальне значення з (5, 2, 8):", Min(5, 2, 8))
	fmt.Println("Середнє значення (5, 10, 15):", Average(5, 10, 15))

	x, err := SolveLinearEquation(2, -4)
	if err != nil {
		fmt.Println("Помилка:", err)
	} else {
		fmt.Println("Розв'язок рівняння 2x - 4 = 0: x =", x)
	}
}
