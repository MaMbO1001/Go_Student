package main

import (
	"fmt"
)

func main() {
	// Задание 1.
	var zero []int = make([]int, 3)
	zero[0] = 1
	zero[1] = 2
	zero[2] = 3
	fmt.Println(zero)
	// Задание 2.
	var qwer []int = make([]int, 5)
	qwer[0] = 1
	qwer[1] = 2
	qwer[2] = 3
	qwer[3] = 4
	qwer[4] = 5
	qwer[4] = 8
	fmt.Println(qwer)
	// Задание 3.
	var nnn []int
	nnn = append(nnn, 1, 2, 3, 4, 6)
	fmt.Println(nnn)
	// Задание 4.
	var x []int
	x = append(x, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	var c = x[1:3]
	fmt.Println(c)
	// Задание 5.
	var b []int = make([]int, 5)
	b[0] = 1
	b[1] = 2
	b[2] = 3
	b[3] = 4
	b[4] = 5
	var v []int = make([]int, len(b))
	copy(v, b)
	fmt.Println(v)
}
