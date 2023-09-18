package main

import (
	"fmt"
)
	const Pi = 3.14159
	const q = 14
	const e = 45.99
	const r = "Book"
	const w = true


func main () {

	// Раздел 3. Константы. Задание 1.
    fmt.Println(Pi)

    // Задание 2.
	const Beta = 100
	fmt.Println(Beta)

	// Задание 3.
    const Pi = 3.11111
	fmt.Println(Pi, Pi)

	// Задание 4.
	fmt.Println(Pi + e)
	fmt.Println(e + Pi)
	fmt.Println(e * Pi)
	fmt.Println(e / Pi)

	// Задание 5.
	// Константы не могут быть изменены.

	// Задание 6.
	fmt.Println("Мне", q, "лет")
	fmt.Println("Размер ноги:", e)
	fmt.Println("Беру на учёбу:", r)
	fmt.Println(w)
}