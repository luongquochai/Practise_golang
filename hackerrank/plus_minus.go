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
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func plusMinus(arr []int32) {
	// Write your code here
	lenArr := int32(len(arr))
	var countPos, countNeg, countZero int32
	for i := int32(0); i < lenArr; i++ {``
		// fmt.Println(arr[i])
		if arr[i] > 0 {
			countPos++
		} else if arr[i] < 0 {
			countNeg++
		} else {
			countZero++
		}
	}
	// resultPos := float64(countPos)/float64(lenArr)
	fmt.Printf("%.6f\n", float64(countPos)/float64(lenArr))
	fmt.Printf("%.6f\n", float64(countNeg)/float64(lenArr))
	fmt.Printf("%.6f\n", float64(countZero)/float64(lenArr))
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}
	// fmt.Println(arr)
	plusMinus(arr)
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
