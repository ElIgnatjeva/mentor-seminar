package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	limit := int(math.Sqrt(float64(n)))
	for i := 3; i <= limit; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Print("Введите число для проверки: ")
	_, err := fmt.Fscan(reader, &n)
	if err != nil {
		fmt.Println("Ошибка: введено не целое число")
		return
	}

	if isPrime(n) {
		fmt.Printf("%d — простое число\n", n)
	} else {
		fmt.Printf("%d — не простое число\n", n)
	}
}