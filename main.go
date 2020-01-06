package main

import (
	"fmt"
	"strconv"
)

var reply string

func main() {
	var a, b string
	fmt.Scan(&a, &b)
	solve([]string{a, b})
}

func solve(buf []string) {
	a, _ := strconv.Atoi(buf[0])
	b, _ := strconv.Atoi(buf[1])
	reply = strconv.Itoa(a ^ b)
	fmt.Println(reply)
}
