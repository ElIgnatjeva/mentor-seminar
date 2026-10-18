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

	if n > 0 {
		fmt.Println("Число положительное")
	} else if n < 0 {
		fmt.Println("Число отрицательное")
	} else {
		fmt.Println("Число равно нулю (ни положительное, ни отрицательное)")
	}

	if n%2 == 0 {
		fmt.Println("Число чётное")
	} else {
		fmt.Println("Число нечётное")
	}

	absN := n
	if absN < 0 {
		absN = -absN
	}

	div3 := absN%3 == 0
	div5 := absN%5 == 0

	switch {
	case div3 && div5:
		fmt.Println("Число делится и на 3, и на 5")
	case div3:
		fmt.Println("Число делится на 3, но не делится на 5")
	case div5:
		fmt.Println("Число делится на 5, но не делится на 3")
	default:
		fmt.Println("Число не делится ни на 3, ни на 5")
	}
}