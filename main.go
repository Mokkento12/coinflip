package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
    // Инициализация генератора случайных чисел
    rand.Seed(time.Now().UnixNano())

    // Объявление переменных для подсчета
    var heads, tails int

    // Цикл для многократного подбрасывания монетки
    for i := 0; i < 10; i++ {
        switch coinflip() {
        case "heads":
            heads++
        case "tails":
            tails++
        default:
            fmt.Println("Приземлились!")
        }
    }

    // Вывод результата
    fmt.Printf("Heads: %d, Tails: %d\n", heads, tails)
}

// Функция для имитации подбрасывания монетки
func coinflip() string {
    // Генерация случайного числа: 0 или 1
    if rand.Intn(2) == 0 {
        return "heads"
    }
    return "tails"
}