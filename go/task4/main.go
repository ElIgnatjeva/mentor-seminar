package main

import (
	"fmt"
)

func printTable(n int) {
	fmt.Printf("Таблица умножения для %d:\n", n)
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", n, i, n*i)
	}
}

func main() {
	var n int
	fmt.Print("Введите число: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Ошибка: введите целое число")
		return
	}

	printTable(n)
}