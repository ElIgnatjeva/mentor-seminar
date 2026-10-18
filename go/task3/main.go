package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите число N: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	n, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Ошибка: введено не целое число")
		return
	}

	if n < 1 {
		fmt.Println("Ошибка: N должно быть натуральным числом (N >= 1)")
		return
	}

	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}

	average := float64(sum) / float64(n)

	fmt.Printf("Сумма чисел от 1 до %d: %d\n", n, sum)
	fmt.Printf("Среднее значение: %.2f\n", average)
}