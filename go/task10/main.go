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

	evenCount := 0
	oddCount := 0

	maxVal := numbers[0]
	minVal := numbers[0]
	sum := 0

	for _, num := range numbers[0:] {
		sum += num

		if num > maxVal {
			maxVal = num
		}
		if num < minVal {
			minVal = num
		}
		
		if num%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}

	average := float64(sum) / float64(n)

	fmt.Println("\n--- Результаты ---")
	fmt.Printf("Сумма всех чисел:      %d\n", sum)
	fmt.Printf("Среднее значение:      %.2f\n", average)
	fmt.Printf("Максимальное значение: %d\n", maxVal)
	fmt.Printf("Минимальное значение:  %d\n", minVal)
	fmt.Printf("Количество чётных:     %d\n", evenCount)
	fmt.Printf("Количество нечётных:   %d\n", oddCount)
}