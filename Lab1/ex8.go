package main

//Импорт кількох пакетів
import (
	"fmt"
	"math"
)

func main() {
	var defaultFloat float32
	var defaultDouble float64 = 5.5

	fmt.Println("defaultfloat       = ", defaultFloat)
	fmt.Printf("defaultDouble (%T) = %f\n\n", defaultDouble, defaultDouble)

	fmt.Println("MAX float32        = ", math.MaxFloat32)
	fmt.Println("MIN float32        = ", math.SmallestNonzeroFloat32, "\n")

	fmt.Println("MAX float64        = ", math.MaxFloat64)
	fmt.Println("MIN float64        = ", math.SmallestNonzeroFloat64, "\n")

	// Завдання.
	// 1. Створіть змінні різних типів, використовуючи короткий запис та ініціалізацію за замовчуванням. Результат вивести в консоль
	var defaultInt int
	var defaultFloat2 float32
	var defaultBool bool
	var defaultString string

	shortInt := 42
	shortFloat := 3.14
	shortBool := true
	shortString := "Go"

	fmt.Println("\nЗмінні з ініціалізацією за замовчуванням:")
	fmt.Printf("defaultInt = %d (тип: %T)\n", defaultInt, defaultInt)
	fmt.Printf("defaultFloat2 = %f (тип: %T)\n", defaultFloat2, defaultFloat2)
	fmt.Printf("defaultBool = %v (тип: %T)\n", defaultBool, defaultBool)
	fmt.Printf("defaultString = \"%s\" (тип: %T)\n\n", defaultString, defaultString)

	fmt.Println("Змінні з коротким записом:")
	fmt.Printf("shortInt = %d (тип: %T)\n", shortInt, shortInt)
	fmt.Printf("shortFloat = %.2f (тип: %T)\n", shortFloat, shortFloat)
	fmt.Printf("shortBool = %v (тип: %T)\n", shortBool, shortBool)
	fmt.Printf("shortString = \"%s\" (тип: %T)\n", shortString, shortString)
}
