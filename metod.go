package main

import (
	"fmt"
)
type Seconds int

func main() {
	// Задание 1.
 r := Rectangle {
	Width: 23.4,
	Height: 12.8,
 }
 fmt.Println(r.Height, r.Width)
 fmt.Println(r.Area())

 // Задание 2.
 fmt.Println(r.isSquare())

 // Задание 3.
 fmt.Println(r.DoubleSize())

 // Задание 4.
 e := Rectangle {
	Width: 33.3,
	Height: 55.5,
 }
 fmt.Println(e.Area())
 fmt.Println(e.isSquare())

 // Задание 5.

 s := Seconds(345)
 fmt.Println(s.Minutes())

 // Задание 6.
 p := Person{
 	Name:    "Viktor",
 	Address: Address{
	    Street: "Пушкинская",
	    City: "Санкт-Петербург",
	    Country: "Russia",
     },
	}
p.FullAddress()
}

// Задание 1.
type Rectangle struct {
	Width float64
	Height float64
}
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Задание 2.
func (r Rectangle) isSquare() bool {
	return r.Width == r.Height
}

// Задание 3.
func (r Rectangle) DoubleSize() float64 {
	return r.Height * r.Width * 2
}

// Задание 5.
func (s Seconds) Minutes() float64 {
	return float64(s) / 60
}

// Задание 6.
type Person struct {
	Name string
	Address
}
type Address struct {
	Street string
	City string
	Country string
}
func (p Person) FullAddress() {
    fmt.Printf("Full Address: %s, %s, %s", p.Country, p.City, p.Street)
}