package main

import (
	"fmt"
	"github.com/SafronovRaff/ping"
	"net/http"
)

func main() {
	client := &http.Client{}
	pinger := ping.Pinger{client}
	url := "https://ya.ru"
	alive := pinger.Ping(url)
	fmt.Println(url, "is alive =", alive)
}