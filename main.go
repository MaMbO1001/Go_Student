package main

import (
	"fmt"
)

func main() {

	// Раздел 5.
	// Задание 1.
	for q := 1; q < 11; q++ {
		fmt.Println(q)
	}

	// Задание 2.
	for w := 0; w < 21; w += 2 {
		fmt.Println(w)
	}

	// Задание 3.
	for e := 0; e < 21; e++ {
		if e%2 == 0 {
			continue
		}
		fmt.Println(e)
	}

	// Задание 4
	z := 0
	for r := 0; r < 100; r++ {
		z += r
		if z > 50 {
			break
		}
		fmt.Println(z)
	}

	// Задание 5
	num := []int{1, 2, 3, 4, 5}
	for _, numb := range num {
		fmt.Println(numb)
	}

	// Задание 6
	vvv := map[string]int{
		"Ff": 3,
		"Fa": 61,
		"Fd": 39,
		"Fv": 15,
	}
	for x, c := range vvv {
		fmt.Println(x, c)
	}

	// Задание 7
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for t, y := range slice {
		fmt.Println(t, y)
	}
	o := "qwertyuiop"
	for j := 0; j < len(o); j++ {
		fmt.Printf("%c", o[j])
	}
}
