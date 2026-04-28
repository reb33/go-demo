package main

import (
	"fmt"
)

// В цикле спрашиваем ввод транзакций: -10, 10, 40.5, до первого 0
// Добавлять каждую в массив транзакций
// Вывести массив
// Вывести сумма баланса в консоль

func main() {
	transactions := []float64{}
	for {
		transaction := scanTransaction()
		if transaction == 0 {
			break
		}
		transactions = append(transactions, transaction)
	}
	fmt.Println(transactions)
	fmt.Println(calculateBalance(transactions))
}

func scanTransaction() (float64) {
	fmt.Print("Введите транзакцию: ")
	var transaction float64
	fmt.Scan(&transaction)
	return transaction
}

func calculateBalance(transactions []float64) (float64) {
	var balance float64 = 0
	for _, transaction := range transactions {
		balance += transaction
	}
	return balance
}