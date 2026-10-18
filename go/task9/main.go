package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var balance float64
	fmt.Print("Введите начальный баланс: ")
	_, err := fmt.Fscan(reader, &balance)
	if err != nil || balance < 0 {
		fmt.Println("Ошибка: баланс должен быть неотрицательным числом")
		return
	}

	for {
		fmt.Println("\n--- Меню ---")
		fmt.Println("1. Посмотреть баланс")
		fmt.Println("2. Пополнить баланс")
		fmt.Println("3. Снять средства")
		fmt.Println("0. Выход")
		fmt.Print("Выберите действие: ")

		var choice int
		_, err := fmt.Fscan(reader, &choice)
		if err != nil {
			reader.ReadString('\n')
			fmt.Println("Ошибка: введите число от 0 до 3")
			continue
		}

		switch choice {
		case 1:
			fmt.Printf("Текущий баланс: %.2f\n", balance)

		case 2:
			var amount float64
			fmt.Print("Введите сумму пополнения: ")
			_, err := fmt.Fscan(reader, &amount)
			if err != nil || amount <= 0 {
				fmt.Println("Ошибка: сумма пополнения должна быть положительной")
				continue
			}
			balance += amount
			fmt.Printf("Баланс успешно пополнен. Новый баланс: %.2f\n", balance)

		case 3:
			var amount float64
			fmt.Print("Введите сумму снятия: ")
			_, err := fmt.Fscan(reader, &amount)
			if err != nil || amount <= 0 {
				fmt.Println("Ошибка: сумма снятия должна быть положительной")
				continue
			}
			if amount > balance {
				fmt.Printf("Ошибка: недостаточно средств. Доступно: %.2f\n", balance)
			} else {
				balance -= amount
				fmt.Printf("Средства успешно сняты. Новый баланс: %.2f\n", balance)
			}

		case 0:
			fmt.Println("Выход из программы.")
			return

		default:
			fmt.Println("Ошибка: неверный пункт меню. Введите 0, 1, 2 или 3.")
		}
	}
}