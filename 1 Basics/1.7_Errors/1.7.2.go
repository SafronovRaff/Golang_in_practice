package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// account представляет счет
type account1 struct {
	balance   int
	overdraft int
}

func main() {
	var acc account1
	var trans []int
	defer func() {
		fmt.Print("-> ")
		err := recover()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(acc, trans)
	}()
	acc, trans = parseInput1()
}

// parseInput считывает счет и список транзакций из os.Stdin.
func parseInput1() (account1, []int) {
	accSrc, transSrc := readInput1()
	acc := parseAccount1(accSrc)
	trans := parseTransactions(transSrc)
	return acc, trans
}

// readInput возвращает строку, которая описывает счет
// и срез строк, который описывает список транзакций.
// эту функцию можно не менять
func readInput1() (string, []string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	accSrc := scanner.Text()
	var transSrc []string
	for scanner.Scan() {
		transSrc = append(transSrc, scanner.Text())
	}
	return accSrc, transSrc
}

// parseAccount парсит счет из строки
// в формате balance/overdraft.
func parseAccount1(src string) account1 {
	parts := strings.Split(src, "/")
	balance, _ := strconv.Atoi(parts[0])
	overdraft, _ := strconv.Atoi(parts[1])
	if overdraft < 0 {
		panic("expect overdraft >= 0")
	}
	if balance < -overdraft {
		panic("balance cannot exceed overdraft")
	}
	return account1{balance, overdraft}
}

// parseTransactions парсит список транзакций из строки
// в формате [t1 t2 t3 ... tn].
func parseTransactions(src []string) []int {
	trans := make([]int, len(src))
	for idx, s := range src {
		t, _ := strconv.Atoi(s)
		trans[idx] = t
	}
	return trans
}
