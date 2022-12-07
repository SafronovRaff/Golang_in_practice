package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	tex, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	tex = strings.Trim(tex, "\n \r")
	num := strings.Count(tex, " ")

	if num == 0 {
		fmt.Println(num)
	} else {
		fmt.Println(num + 1)
	}
}
