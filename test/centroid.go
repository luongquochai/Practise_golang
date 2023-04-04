package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func numbers(s string) []int {
	var n []int
	for _, f := range strings.Fields(s) {
		i, err := strconv.Atoi(f)
		if err == nil {
			n = append(n, i)
		}
	}
	return n
}

func toNumber(s string) int {
	var n int
	for _, f := range strings.Fields(s) {
		i, err := strconv.Atoi(f)
		if err == nil {
			n = i
		}
	}
	return n
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var N int
	N = toNumber(scanner.Text())
	var list []int
	scanner.Scan()
	list = numbers(scanner.Text())

	var i, j int
	for i < N && j < N {
		var num int
		num = list[i] * list[j]
		i++
		j++
		fmt.Println(num)
	}
}
