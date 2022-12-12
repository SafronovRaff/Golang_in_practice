package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

/*Счет в банке
Напишите логику работы с лицевым счетом в банке. У счета есть баланс (сумма на счете) и овердрафт (на какую сумму можно уйти в минус). К счету последовательно применяются транзакции списания и пополнения. Транзакции изменяют баланс счета, овердрафт не меняется.

Со счета нельзя списать больше, чем (баланс + овердрафт). Зачислить или списать нулевую сумму тоже нельзя.

Например:

начальное состояние счета: баланс = 100, овердрафт = 10
транзакции: 10, -50, 20
результат: баланс = 80, овердрафт = 10
Или:

начальное состояние счета: баланс = 100, овердрафт = 10
транзакции: -100, -10, -10
результат: недостаточно средств на счете
Гарантируется, что сумма на счете, овердрафт и размер транзакции не выйдут за пределы типа int, как по отдельности, так и все вместе. Гарантируется, что овердрафт больше либо равен 0.

Sample Input:

{100 10} [10 -50 20]
Sample Output:

{80 10}*/
// errInsufficientFunds сигнализирует,
// что на счете недостаточно денег,
// чтобы выполнить списание
var errInsufficientFunds error = errors.New("insufficient funds")

// errInvalidAmount сигнализирует,
// что указана некорректная сумма транзакции
var errInvalidAmount error = errors.New("invalid transaction amount")

// account представляет счет
type accountt struct {
	balance   int
	overdraft int
}

// deposit зачисляет деньги на счет.
func (acc *accountt) deposit(amount int) error {
	// ...
	if amount > 0 {
		acc.balance = acc.balance + amount
	} else {
		return errInvalidAmount
	}
	return nil
}

// withdraw списывает деньги со счета.
func (acc *accountt) withdraw(amount int) error {
	if acc.balance+acc.overdraft < amount {
		return errInsufficientFunds
	} else {
		acc.balance = acc.balance - amount
	}
	return nil

}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

type test struct {
	acc   accountt
	trans []int
}

var tests = map[string]test{
	"{100 10} [10 -50 20]":   {accountt{100, 10}, []int{10, -50, 20}},
	"{30 0} [-20 -10]":       {accountt{30, 0}, []int{-20, -10}},
	"{30 0}, [-20 -10 -10]":  {accountt{30, 0}, []int{-20, -10, -10}},
	"{30 0}, [-100]":         {accountt{30, 0}, []int{-100}},
	"{0 0}, [10 20 30]":      {accountt{0, 0}, []int{10, 20, 30}},
	"{0 0}, [10 -10 20 -20]": {accountt{0, 0}, []int{10, -10, 20, -20}},
	"{20 10}, [-20 -10]":     {accountt{20, 10}, []int{-20, -10}},
	"{20 10}, [-20 -10 -10]": {accountt{20, 10}, []int{-20, -10, -10}},
	"{0 100}, [-20 -10]":     {accountt{0, 100}, []int{-20, -10}},
	"{0 30}, [-20 -10]":      {accountt{0, 30}, []int{-20, -10}},
	"{0 30}, [-20 -10 -10]":  {accountt{0, 30}, []int{-20, -10, -10}},
	"{70 30}, [-100 100]":    {accountt{70, 30}, []int{-100, 100}},
	"{100 10}, [10 0 20]":    {accountt{100, 10}, []int{10, 0, 20}},
}

func main() {
	var err error
	name, err := readString()
	if err != nil {
		log.Fatal(err)
	}
	testCase, ok := tests[name]
	if !ok {
		log.Fatalf("Test case '%v' not found", name)
	}
	for _, t := range testCase.trans {
		if t >= 0 {
			err = testCase.acc.deposit(t)
		} else {
			err = testCase.acc.withdraw(-t)
		}
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	if err == nil {
		fmt.Println(testCase.acc)
	}
}

// readString считывает и возвращает строку из os.Stdin
func readString() (string, error) {
	rdr := bufio.NewReader(os.Stdin)
	str, err := rdr.ReadString('\n')
	if err != nil && err != io.EOF {
		return "", err
	}
	return strings.TrimSpace(str), nil
}
