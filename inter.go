package main

import (
	"fmt"
	"sort"
)
    // Rectangle на 1 и 2 задания)
type Rectangle struct{
	width int
	height int
}

// Задание 1.

type Shape interface {
	Area()
}

func (r Rectangle) Area() int {
	return r.width * r.height
}

	// Задание 2.

type Resizable interface {
	Resize()
	DoubleSize()
}

func (r Rectangle) Resize() int {
	return (r.height * r.width) * 3
}

func (r *Rectangle) DoubleSize() {
	r.height *= 2
	r.width *= 2
}


// Задание 3.

type Stringer interface {
	String() string
}

func (p Person) String() string {
	return fmt.Sprintf("%s, %d года", p.Name, p.Age)
}

type Person struct {
	Name string
	Age int
}


// Задание 4.

type Seconds int

type ConvertToMinutes interface {
	Minutes()
}

func (s Seconds) Minutes() int {
	return int(s) * 60
}

// Задание 5.

type FloatSlice []float64

func (f FloatSlice) Len() int {
	return len(f)
}

func (f FloatSlice) Less(i, j int) bool {
    return f[i] < f[j]
}

func (f FloatSlice) Swap(i, j int) {
    f[i], f[j] = f[j], f[i] 
}


// Задание 6.

func ffc (ffa interface{}) {
	if a, ok := ffa.(int); ok {
		fmt.Println("int:",a)
	} else if a, ok := ffa.(string); ok {
		fmt.Println("String:",a)
	} else if a, ok := ffa.(bool); ok {
		fmt.Println("Bool:", a)
	} else if a, ok := ffa.(float64); ok {
		fmt.Println("Float:", a)
	} else {
		fmt.Println("Без понятия:", a)
	}
}



func main() {

	// Задание 1.
	
	r := Rectangle{width: 5, height: 12}
    fmt.Println(r.Area())
	

	// Задание 2.

    fmt.Println(r.Resize())
    r.DoubleSize()
	fmt.Println(r)


	// Задание 3.

    p := Person{Name: "Viktor", Age: 22}
	fmt.Println(p.String())


	// Задание 4.

	s := Seconds(5)
	fmt.Println(s.Minutes())


	// Задание 5.

	var f FloatSlice
	f = append(f, 1.1, 55.1, 3.17, 25.9, 4.1, 67,1)
	fmt.Println(f)
	sort.Sort(f)
	fmt.Println(f)


	// Задание 6.

	ffc(14)
	ffc("Hello world")
    ffc(true)
	ffc(112.2)
	ffc(&f)
}