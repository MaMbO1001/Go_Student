package main

import (
	"fmt"
)

func main() {
	var test [5]int
	test[0] = 2
	test[1] = 4
	test[2] = 6
	test[3] = 8
	test[4] = 10
	for i := 0; i < len(test); i++ {
		fmt.Println(test[i])
	}
}
