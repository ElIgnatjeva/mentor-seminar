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

	var n int
	fmt.Print("Введите количество чисел N: ")
	_, err := fmt.Fscan(reader, &n)
	if err != nil || n <= 0 {
		fmt.Println("Ошибка: N должно быть натуральным числом (N >= 1)")
		return
	}

	reader.ReadString('\n')

	fmt.Print("Введите числа через пробел: ")
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)

	fields := strings.Fields(line)
	numbers := make([]int, 0, len(fields))
	for _, f := range fields {
		num, err := strconv.Atoi(f)
		if err != nil {
			fmt.Println("Пропущено (не число):", f)
			continue
		}
		numbers = append(numbers, num)
	}

	if len(numbers) < n {
		fmt.Printf("Ошибка: введено %d чисел, а требовалось %d\n", len(numbers), n)
		return
	}

	numbers = numbers[:n]  

	positiveCount := 0
	negativeCount := 0
	zeroCount := 0

	for _, num := range numbers[0:] {
		if num > 0 {
			positiveCount++
		} else if num < 0 {
			negativeCount++
		} else {
			zeroCount++
		}
	}

	fmt.Printf("Положительных чисел: %d\n", positiveCount)
	fmt.Printf("Отрицательных чисел: %d\n", negativeCount)
	fmt.Printf("Нулей: %d\n", zeroCount)
}