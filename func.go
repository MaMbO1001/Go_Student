package main

import (
	"fmt"
)

func main() {
	// Раздел 6. ФУНКЦИИ.
	// Задание 1
	hello()

	// Задание 2.
	result := add(14, 22)
	fmt.Println(result)

	// Задание 3.
	numb := addNumb(94, 22)
	fmt.Println(numb)

	// Задание 4.
	Rere, rerere := reReturn(34, 21)
	fmt.Println(Rere)
	fmt.Println(rerere)

	// Задание 5.
	ad(1)
	ad(1, 2, 3, 4, 11)
	ad(78, 111, 73)

	// Задание 6.
	one()
}

// Раздел 6. ФУНКЦИИ.
// Задание 1.
func hello() {
	fmt.Println("Приветствую вас!")
}

// Задание 2.
func add(x int, y int) int {
	return x + y
}

// Задание 3.
func addNumb(x int, y int) int {
	return x * y
}

// Задание 4.
func reReturn(x int, y int) (int, int) {
	return x + y, x - y
}

// Задание 5.
func ad(numb ...int) {
	for _, n := range numb {
		fmt.Printf("Number: %d\n", n)
	}
}

// Задание 6.
func one() {
	fmt.Println("It's one")
	two()
}
func two() {
	fmt.Println("It's two")
}

// Задание 7.
// Go не допускает значения по умолчанию в прототипах функций или при перегрузке
//функций.

//Задание 8.
//Функции в Go тоже являются значениями. Их можно передавать так же как и
//остальные значения. Функции могут быть использованы как аргументы других функций
// и как возвращаемые значения.

// Но про 8 задание я бы послушал.Ибо не до конца понял
