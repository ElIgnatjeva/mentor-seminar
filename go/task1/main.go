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

	fmt.Print("Введите первое число: ")
	aStr, _ := reader.ReadString('\n')
	aStr = strings.TrimSpace(aStr)
	a, err := strconv.ParseFloat(aStr, 64)
	if err != nil {
		fmt.Println("Ошибка: первое число введено некорректно")
		return
	}

	fmt.Print("Введите второе число: ")
	bStr, _ := reader.ReadString('\n')
	bStr = strings.TrimSpace(bStr)
	b, err := strconv.ParseFloat(bStr, 64)
	if err != nil {
		fmt.Println("Ошибка: второе число введено некорректно")
		return
	}

	fmt.Print("Введите операцию (+, -, *, /): ")
	op, _ := reader.ReadString('\n')
	op = strings.TrimSpace(op)

	var result float64
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			fmt.Println("Ошибка: деление на ноль невозможно")
			return
		}
		result = a / b
	default:
		fmt.Println("Ошибка: неизвестная операция")
		return
	}

	if result == float64(int64(result)) {
		fmt.Printf("Результат: %d\n", int64(result))
	} else {
		fmt.Printf("Результат: %g\n", result)
	}
}