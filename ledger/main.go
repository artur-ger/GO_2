package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type Transaction struct {
	ID          int
	Amount      float64
	Category    string
	Description string
	Date        time.Time
}

var transactions = []Transaction{}

func AddTransaction(tx Transaction) error {
	if tx.Amount == 0 {
		return errors.New("amount must not be zero")
	}
	if strings.TrimSpace(tx.Category) == "" {
		return errors.New("category must not be empty")
	}
	if tx.Date.IsZero() {
		tx.Date = time.Now()
	}

	tx.ID = len(transactions) + 1
	transactions = append(transactions, tx)
	return nil
}

func ListTransactions() []Transaction {
	result := make([]Transaction, len(transactions))
	copy(result, transactions)
	return result
}

func main() {
	fmt.Println("Ledger service started")

	testTransactions := []Transaction{
		{Amount: 1500, Category: "Продукты", Description: "Покупки в супермаркете", Date: time.Date(2026, 9, 25, 0, 0, 0, 0, time.Local)},
		{Amount: 350.5, Category: "Транспорт", Description: "Такси до работы", Date: time.Date(2026, 9, 26, 0, 0, 0, 0, time.Local)},
		{Amount: 2990, Category: "Развлечения", Description: "Билеты в кино", Date: time.Date(2026, 9, 27, 0, 0, 0, 0, time.Local)},
		{Amount: 0, Category: "Прочее", Description: "Некорректная транзакция"},
	}

	for _, tx := range testTransactions {
		if err := AddTransaction(tx); err != nil {
			fmt.Printf("Ошибка добавления транзакции %q: %v\n", tx.Description, err)
		}
	}

	fmt.Println("Список транзакций:")
	for _, tx := range ListTransactions() {
		fmt.Printf("  #%d | %s | %-12s | %8.2f | %s\n",
			tx.ID, tx.Date.Format("2006-01-02"), tx.Category, tx.Amount, tx.Description)
	}
}
