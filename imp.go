package main

import (
	"fmt"
	"math"
	"time"
	"math/rand"
	log "Go_Student/logger"
)

func main() {

	// Раздел 7. Импорты.
	// Задание 1
fmt.Println(math.Sqrt(28))

    // Задание 2
	s := time.Now()
	t := time.Now().Second()
	fmt.Println("Текущее время: ", s, "Секунды: ", t)
	var f float64 = float64(t)
	fmt.Println("Корень текуших секунд:", math.Sqrt(f))

	// Задание 3
    fmt.Println(rand.Intn(100))

    // Раздел 8. Организация....
	// Задание 1.
	log.Log("Смска из логгера,хы")
}