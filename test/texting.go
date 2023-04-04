package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()
	
	keypad = [4][3]string{
		[1, 2, 3'],
          ['4', '5', '6'],
          ['7', '8', '9'],
          ['*', '0', '#']
	}
}
