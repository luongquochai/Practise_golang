package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'staircase' function below.
 *
 * The function accepts INTEGER n as parameter.
 */
func staircase(n int32) {
	str := "#"
	strSpace := " "
	// var sb strings.Builder
	// s := make([]string, n)
	// Write your code here
	for i := int32(0); i < n; i++ {
		space := n - i //5
		// fmt.Println(space)
		var sb strings.Builder
		for k := int32(0); k < space-1; k++ {
			sb.WriteString(strSpace)
		}
		for j := int32(0); j < n-space+1; j++ {
			sb.WriteString(str)
		}

		fmt.Printf("%v\n", sb.String())

		// fmt.Println()
	}

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	staircase(n)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
