package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	numDice := 1

	if len(os.Args) > 1 {
		i, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println(err)
		}
		numDice = i
	}

	seed := rand.NewSource(time.Now().UnixNano())
	randomNumber := rand.New(seed)

	die := []string{
		" ------   ",
		"|      |  ",
		"|      |  ",
		"|      |  ",
		" ------   ",
	}

	for i := 0; i < 5; i++ {
		for j, n := 0, numDice; j < n; j++ {
			if i == 3 {
				fmt.Printf("|    %d |  ", randomNumber.Intn(5)+1)
			} else {
				fmt.Print(die[i])
			}
		}
		fmt.Println()
	}
}
