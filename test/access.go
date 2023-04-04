package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func numbers(s string) int {
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
	var N int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	N = numbers(scanner.Text())
	// loop based on Number of file

	files := make(map[string]map[string]bool)

	for i := 0; i < N; i++ {
		var fileName, str string
		var operations []string
		scanner.Scan()
		str = scanner.Text()
		// fmt.Println(fileName)

		for i, oper := range strings.Split(str, " ") {
			switch i {
			case 0:
				fileName = oper
			default:
				operations = append(operations, oper)
			}
		}

		files[fileName] = make(map[string]bool)
		for _, operation := range operations {
			files[fileName][operation] = true
		}

	}
	// fmt.Println(files)
	var M int
	scanner.Scan()
	M = numbers(scanner.Text())

	for i := 0; i < M; i++ {
		var fileName, operation, str string
		var strSplit []string
		scanner.Scan()
		str = scanner.Text()
		// fmt.Println(str)
		strSplit = strings.Split(str, " ")
		operation = strSplit[0]
		if operation == "write" {
			operation = "W"
		} else if operation == "read" {
			operation = "R"
		} else if operation == "execute" {
			operation = "X"
		}
		fileName = strSplit[1]

		// fmt.Println(fileName)
		// fmt.Println(operation)
		// fmt.Println("-======")
		// Kiểm tra xem tên tệp có tồn tại trong bảng hash không
		if _, ok := files[fileName]; !ok {
			fmt.Println("Access denied")
			continue
		}
		// fmt.Println(files[fileName][operation])

		// Kiểm tra xem hoạt động được yêu cầu có được phép hay không
		if _, ok := files[fileName][operation]; ok {
			fmt.Println("OK")
		} else {
			fmt.Println("Access denied")
		}
	}
}
