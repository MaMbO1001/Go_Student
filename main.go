package main

import (
	"fmt"
)

func main() {
	// Задание 1.
	var test map[string]int
	test = map[string]int{
		"Viktor": 22,
		"Ed": 28,
		"Anna": 18,
	}
	fmt.Println(test)
	// Задание 2
	test["Alice"] = 23
	test["Margo"] = 32
	fmt.Println(test)
	// Задание 3
	delete(test, "Ed")
	fmt.Println(test)
	// Задание 4
	a, b := test["Viktor"]
	if b == true{
		fmt.Println(a)
	}
	// Задание 5
	for name, age := range test{
		fmt.Printf("%s is %d years old//n", name, age)
	}
}

