package main

import (
	"bufio"
	"sort"

	// "encoding/json"
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
func arrayToString(a []int) string {
	return strings.Trim(fmt.Sprint(a), "[]")
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

func getOutPut(s []string) {
	var list string
	var list1 []int
	list = strings.Join(s, " ")
	// fmt.Printf("%T",list)
	list1 = numbers(list)
	// fmt.Printf("%T",list1)
	elements := make(map[int]bool) //create empty map
	output := make([]int, 0)       //create output slice

	for i := range list1 {
		if _, ok := elements[list1[i]]; !ok {
			elements[list1[i]] = true
			// fmt.Printf("%v ", list1[i])
			output = append(output, list1[i])
		}
	}
	sort.Sort(sort.IntSlice(output))
	fmt.Println(arrayToString(output))
	// fmt.Printf("%v ", output)
}

func main() {
	var _ int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	// Get T
	_ = toNumber(scanner.Text())
	// fmt.Println(T)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		if len(strings.TrimSpace(line)) == 0 {
			break
		}
		lines = append(lines, line)
	}
	// fmt.Printf("%T", lines[1])
	// var arr string
	// var pos int
	for i := 1; i < len(lines); i = i + 3 {
		// pos = i+2
		arr := lines[i : i+2]
		getOutPut(arr)
		// fmt.Printf("\n")
	}
}
