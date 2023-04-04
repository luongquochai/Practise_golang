package main

import (
	"bufio"
	"fmt"
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

func matrix(n int, s string) {
	matrix := make([][]string, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]string, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j] = s
		}
	}
	// fmt.Printf("%v", matrix)
	// // fmt.Printf("%T", matrix)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("matrix[%d][%d]: ", i, j)
			fmt.Print(matrix[i][j])
		}
		fmt.Println()

	}

	// fmt.Printf("%v", matrix)

	// return matrix
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var N int
	N = toNumber(scanner.Text())
	fmt.Println(N)
	var image = [3][3]string{
		{"o", "x", "o"},
		{"o", "x", "x"},
		{"o", "o", "o"},
	}

	matrix := make([][]string, 3*N)
	for i := 0; i < 1*N; i++ {
		matrix[i] = make([]string, 3*N)
	}
	var total string
	for i := 0; i < 3*N; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				total = image[i][k]
			}
			matrix[i][j] = total[]
		}
	}

	for i := 0; i < 3*N; i++ {
		for j := 0; j < 3*N; j++ {
			fmt.Print(matrix[i][j], "\n")
		}
	}

	// fmt.Println(len(matrix))
	// fmt.Println(image)

	// matrix3 := make([][]string, 3*N)
	// for i := 0; i < 3*N; i++ {
	// 	matrix3[i] = make([]string, 3*N)
	// }

	// matrixTemp := make([][]string, 3)
	// for i := 0; i < 3; i++ {
	// 	matrixTemp[i] = make([]string, 3)
	// }

	// // fmt.Printf("%v", len(matrix2))
	// for i := 0; i < 3; i++ {
	// 	for j := 0; j < 3; j++ {
	// 		matrix(N, image[i][j])
	// 	}
	// }

}
