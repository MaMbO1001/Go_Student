package main

import (
	"fmt"
	"strconv"
)

func main() {

	//Раздел 2. Задание 1.
	var a int = 10
	var b float64 = 77.1
	var c string = "Name"
	var isA bool = 4 == 0
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(isA)

	// Задание 2.
	var d int = 35
	fmt.Println(a + d)
	fmt.Println(a - d)
	fmt.Println(a * d)
	fmt.Println(d / a)

	// Задание 3.
	x := true
	y := false
	fmt.Println(x && y)
	fmt.Println(x || y)

	// Задание 4.
	var g string = "10"
	var h, err = strconv.Atoi(g)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(h)
	}

	// Задание 5 понятно.
	// задание 6 понятно.

	// Задание 7.
	fmt.Println(!y)
	fmt.Println(!x)

	// Задание 8.
	var o float64 = 17.9
	var p int = int(o)
	fmt.Println(p)
	var l float64 = float64(p)
	fmt.Println(l)

	// Задание 9.
	e := 10
	fmt.Println(e)
	// Мы не можем присвоить 2 значения одной переменной.

	// Задание 10.
	var j int
	var u string
	var t int
	t, err = strconv.Atoi(u)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(t)
	}

	fmt.Println("Калькулятор сложения")
	fmt.Scan(&t)
	fmt.Scan(&j)
	k := t + j
	fmt.Println(k)
}
