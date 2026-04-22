package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type BankClient interface {
	// Deposit deposits given amount to clients account
	Deposit(amount int)

	// Withdrawal withdraws given amount from clients account.
	// return error if clients balance less the withdrawal amount
	Withdrawal(amount int) error

	// Balance returns clients balance
	Balance() int
}

type Client struct {
	accountMutex sync.RWMutex

	account_amount int
}

func (c *Client) Deposit(amount int) {
	c.accountMutex.Lock()
	c.account_amount += amount
	c.accountMutex.Unlock()
}

func (c *Client) Withdrawal(amount int) error {
	c.accountMutex.Lock()

	if c.account_amount < amount {
		c.accountMutex.Unlock()
		return fmt.Errorf("на балансе недостаточно средств")
	}
	c.account_amount -= amount

	c.accountMutex.Unlock()

	return nil

}

func (c *Client) Balance() int {
	c.accountMutex.RLock()

	result := c.account_amount

	c.accountMutex.RUnlock()

	return result
}

func main() {
	fmt.Println("Testoviy Zapusk")
	var command string
	var addAmount int
	var withdrawalAmount int

	var client = &Client{}
	var err error

	for range 10 {

		go func() {

			for {
				client.Deposit(rand.Intn(10) + 1)
				time.Sleep(time.Duration(rand.Float64()*500.0+500.0) * time.Millisecond)
			}
		}()

	}

	for range 5 {

		go func() {

			for {
				err = client.Withdrawal(rand.Intn(5) + 1)
				if err != nil {
					fmt.Println(err.Error())
				}
				time.Sleep(time.Duration(rand.Float64()*500.0+500.0) * time.Millisecond)
			}
		}()

	}

	for {
		fmt.Println("")
		fmt.Println("Введите команду")
		fmt.Println("You can use commands: balance, deposit, withdrawal, exit")

		if _, err := fmt.Scanln(&command); err != nil {
			fmt.Println("Ошибка ввода пользователя", err)
		}

		switch command {
		case "balance":
			fmt.Println("Ваш баланс: ", client.Balance())
		case "deposit":
			if _, err := fmt.Scanln(&addAmount); err != nil {
				fmt.Println("Ошибка ввода значения вносимых средств")
			}
			if addAmount >= 0 {
				client.Deposit(addAmount)
				fmt.Println("Средства успешно внесены")
			} else {
				fmt.Println("Вносимых средств должно быть больше 0")
			}
		case "withdrawal":
			if _, err := fmt.Scanln(&withdrawalAmount); err != nil {
				fmt.Println("Ошибка ввода значения списываемых средств")
			}
			if withdrawalAmount >= 0 {
				if err = client.Withdrawal(withdrawalAmount); err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Println("Средства успешно списаны")
				}

			} else {
				fmt.Println("Списываемых средств должно быть больше 0")
			}
		case "exit":
			fmt.Println("До свидания!")
			return
		default:
			fmt.Println("Unsupported command. You can use commands: balance, deposit, withdrawal, exit")
		}

	}
}
