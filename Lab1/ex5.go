package main

import "fmt"

func main() {
	fmt.Println("Синоніми цілих типів\n")

	fmt.Println("byte    - int8")
	fmt.Println("rune    - int32")
	fmt.Println("int     - int32 або int64, в залежності від платформи")
	fmt.Println("uint    - uint32 або uint64, в залежності від платформи")

	// Завдання.
	// 1. Визначити розрядність платформи
	fmt.Printf("\nРозрядність int: %d біт\n", 32<<(^uint(0)>>63))
	fmt.Printf("Розрядність uint: %d біт\n", 32<<(^uint(0)>>63))
}

