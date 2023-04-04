package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {
	var guest int
	var ArrivalTime, RentalTime []int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	guest = toNumber(scanner.Text())

	// fmt.Println("guest: ", guest)
	scanner.Scan()
	ArrivalTime = numbers(scanner.Text())
	scanner.Scan()
	RentalTime = numbers(scanner.Text())

	var leaves []int
	for i := 0; i < guest; i++ {
		leaves = append(leaves, ArrivalTime[i]+RentalTime[i])
	}
	// fmt.Println(leaves)
	sort.Sort(sort.IntSlice(ArrivalTime))
	sort.Sort(sort.IntSlice(leaves))
	// fmt.Println("ArrivalTime: ", ArrivalTime)
	// fmt.Println("RentalTime: ", RentalTime)

	i, j := 0, 0
	var numberOfRooms int
	var maxNumberOfRooms int
	// var maxRooms int
	for i < guest && j < guest {
		if ArrivalTime[i] < leaves[j] {
			numberOfRooms += 1
			maxNumberOfRooms = Max(maxNumberOfRooms, numberOfRooms)
			i++
		} else if ArrivalTime[i] > leaves[j] {
			numberOfRooms -= 1
			j++
		} else {
			i++
			j++
		}
	}
	fmt.Println(maxNumberOfRooms)
	return

}
