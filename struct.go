package main

import (
	"fmt"
)

func main() {

	// Задание 1.
	myBook := Book {
		Title : "Мой путь программиста",
		Author: "Виктор",
		Pages: 620,
	}
	fmt.Println("Название книги: ", myBook.Title)
	fmt.Println("Автор книги: ", myBook.Author)
	fmt.Println("Кол-во страниц: ", myBook.Pages)

	// Задание 2.
	p := Person {
        Name : "Виктор",
		Age : 22,
		isMarried : true,
	}
	fmt.Println(p.Name)
	fmt.Println(p.Age)
	fmt.Println(p.isMarried)

	// Задание 3.
    pp := Pers {
        Name : "Виктор",
		Age : 22,
		isMarried : true,
	}
	var ages *int = &pp.Age
	*ages = 25 
	fmt.Println(pp.Name)
	fmt.Println(pp.Age)
	fmt.Println(pp.isMarried)

	// Задание 4.
	s := Employee {
		Name: "Viktor",
	}
	ss := Job {
		Title: "Proger",
		Salary: "18000 рублей",
	}
	fmt.Println("Имя: ", s.Name)
	fmt.Println("Должность: ", ss.Title)
	fmt.Println("Зарплата: ", ss.Salary)

	// Задание 5.
	d1 := Student{"Viktor", 1}
    d2 := Student{"Viktor", 2}
    fmt.Println(d1 == d2)
}

	// Задание 1.
	type Book struct {
		Title string
		Author string
		Pages int
	}

// Задание 2.
type Person struct {
	Name string
	Age int
	isMarried bool
}

// Задание 3.
type Pers struct {
	Name string
	Age int
	isMarried bool
}

//Задание 4.
type Employee struct {
	Name string
	Job
}
type Job struct {
	Title string
	Salary string
}

// Задание 5.
type Student struct {
	name string
	numberss int
}