package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Print("Введите число N: ")
	_, err := fmt.Fscan(reader, &n)
	if err != nil || n < 1 {
		fmt.Println("Ошибка: N должно быть натуральным числом (N >= 1)")
		return
	}

	sum := 0
	count := 0

	for i := 1; i <= n; i++ {
		if i%2 == 0 && i%3 == 0 {
			sum += i
			count++
		}
	}

	fmt.Printf("Сумма чисел от 1 до %d, кратных 3: %d\n", n, sum)
	fmt.Printf("(Таких чисел: %d)\n", count)
}