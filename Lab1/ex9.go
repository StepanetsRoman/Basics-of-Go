package main

import "fmt"

func main() {
	var first, second bool
	var third bool = true
	fourth := !third
	var fifth = true

	fmt.Println("first  = ", first)       // false
	fmt.Println("second = ", second)      // false
	fmt.Println("third  = ", third)       // true
	fmt.Println("fourth = ", fourth)      // false
	fmt.Println("fifth  = ", fifth, "\n") // true

	fmt.Println("!true  = ", !true)        // false
	fmt.Println("!false = ", !false, "\n") // true

	fmt.Println("true && true   = ", true && true)         // true
	fmt.Println("true && false  = ", true && false)        // false
	fmt.Println("false && false = ", false && false, "\n") // false

	fmt.Println("true || true   = ", true || true)         // true
	fmt.Println("true || false  = ", true || false)        // true
	fmt.Println("false || false = ", false || false, "\n") // false

	fmt.Println("2 < 3  = ", 2 < 3)        // true
	fmt.Println("2 > 3  = ", 2 > 3)        // false
	fmt.Println("3 < 3  = ", 3 < 3)        // false
	fmt.Println("3 <= 3 = ", 3 <= 3)       // true
	fmt.Println("3 > 3  = ", 3 > 3)        // false
	fmt.Println("3 >= 3 = ", 3 >= 3)       // true
	fmt.Println("2 == 3 = ", 2 == 3)       // false
	fmt.Println("3 == 3 = ", 3 == 3)       // true
	fmt.Println("2 != 3 = ", 2 != 3)       // true
	fmt.Println("3 != 3 = ", 3 != 3, "\n") // false

	// Завдання.
	// 1. Пояснити результати операцій
	fmt.Println("\nПояснення результатів операцій:")
	fmt.Println("! - логічне НІ: інвертує значення (true -> false, false -> true)")
	fmt.Println("&& - логічне І: повертає true тільки якщо обидва операнди true")
	fmt.Println("|| - логічне АБО: повертає true якщо хоча б один операнд true")
	fmt.Println("<, >, <=, >= - операції порівняння: порівнюють числові значення")
	fmt.Println("== - рівність: перевіряє чи рівні значення")
	fmt.Println("!= - нерівність: перевіряє чи не рівні значення")
}

