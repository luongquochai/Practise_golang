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
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */
func Sum(arr []int32) int64 {
	response := int64(0)
	for _, number := range arr {
		response += int64(number)
	}
	return response
}

func MinMax(array []int64) (int64, int64) {
	var max int64 = array[0]
	var min int64 = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func miniMaxSum(arr []int32) {
	// arr = []int(arr)
	// Write your code here
	var arrSum []int64
	sumArr := Sum(arr)
	// fmt.Printf("Sum: %d",sumArr)
	for i := int32(0); i < int32(len(arr)); i++ {
		// arrSum = arrSum + int64(arr[i])

		arrSum = append(arrSum, sumArr-int64(arr[i]))
	}
	// fmt.Printf("%v",arrSum)
	minVal, maxVal := MinMax(arrSum)

	fmt.Printf("%d %d", minVal, maxVal)

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
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
