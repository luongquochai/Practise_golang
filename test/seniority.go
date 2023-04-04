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
	s = strings.Trim(s, "[]")
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

func arrayToString(a []int) string {
	return strings.Trim(fmt.Sprint(a), "[]")
}

func printSlice(s []int) {
	for _, value := range s {
		fmt.Printf("%v ", value)
	}
}
func main() {
	var _ int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	_ = toNumber(scanner.Text())
	// fmt.Println(N)
	var arrNum []int
	scanner.Scan()
	scanner.Text()
	// fmt.Println(scanner.Text())
	arrNum = numbers(scanner.Text())
	// fmt.Println(arrNum)
	// fmt.Println(arrNum)
	// var output []int
	for _, i := range arrNum {
		k := int(0)
		for _, j := range arrNum {
			if i > j {
				k++
			}
		}
		fmt.Printf("%v ", k)
		// output = append(output, k)
	}
	// printSlice(output)
}
