package main

import (
	"fmt"
)

func main() {
	// Задание 1.
    num := []int{1, -3, 5, -12, 28}
	for _, num := range num {
		if num > 0 {
			fmt.Println("Good")
		} else if num < 0 {
			fmt.Println("Not good")
		}
	}

	// Задание 2.
	ggs := []int {1, 3, 52, 78, 99, 10, -10, -120}
	for _, ggs := range ggs {
		if ggs >= 10 {
			fmt.Println("Больше либо равно 10")
		} else if ggs < 10 {
			fmt.Println("Меньше 10")
		}
	}
    // Задание 3.
	hh := []int {10, 11, 25, 20, 30, -30, 37, 92}
	for _, hh := range hh {
		if hh == 10 {
			fmt.Println("Число равно 10, всо отлично)")
		} else if hh == 20 {
            fmt.Println("Число равно 20")
		} else if hh == 30 {
			fmt.Println("Число равно 30")
		} else {
			fmt.Println("Данное число не подходит по условию")
		}
	}

	// Задание 4.
	days := []int {1, 2, 3, 4, 5, 6, 7}
	
	for _, day := range days {
		switch day {
		case 1 :
			fmt.Println("Понедельник")
		case 2 :
			fmt.Println("Вторник")	
		case 3 :
			fmt.Println("Среда")	
		case 4 :
			fmt.Println("Четверг")	
		case 5 :
			fmt.Println("Пятница")	
		case 6 :
			fmt.Println("Суббота")	
		case 7 :
			fmt.Println("Воскресенье")	
		}
	}

	// Задание 5.
	mms := []int {1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 13}

	for _, mm := range mms{
	switch mm % 2 == 0 {
	case true:
		fmt.Println("Четное")
	case false:
		fmt.Println("Нечетное")
	} 
}

    // Задание 6.
	bol := []bool {true, false, false, true, false, true, true}

	for _, bo := range bol {
		switch bo {
		case bo || true:
			fmt.Println("Good")

			default :
			fmt.Println("Not Good")
		}
	}

	// Задание 7.
	fal := []int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, fals := range fal {
		switch fals {
		case 1 :
			fmt.Println("Первый")
		case 2 :
			fmt.Println("Второй")
			fallthrough
		case 11 :
			fmt.Println("Третий")
		case 4 :
			fmt.Println("Четвертый")
		case 5 :
			fmt.Println("Пятый")
		case 6 :
			fmt.Println("Шестой")
		case 7 :
			fmt.Println("Седьмой")
		case 8 :
			fmt.Println("Восьмой")
		case 9 :
			fmt.Println("Девятый")
		case 10 :
			fmt.Println("Десятый")
		}
	}
}