package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

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

func main() {
	var n, v, t int
	var list []int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	list = numbers(scanner.Text())
	n, v = list[0], list[1]
	t = int(math.Max(1, float64(n-v)))
	t = int(math.Min(float64(n-1), float64(v))) + (t+1)*t/2 - 1
	fmt.Println(t)

	// 	a = max(1,n-v)
	// print min(n-1,v)+(a+1)*a/2-1

}
