package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func verifyPassword(s string) bool {
	var hasNumber, hasUpperCase, hasLowercase, hasSpecial bool
	for _, c := range s {
		switch {
		case unicode.IsNumber(c):
			hasNumber = true
		case unicode.IsUpper(c):
			hasUpperCase = true
		case unicode.IsLower(c):
			hasLowercase = true
		case c == '#' || c == '|':
			return false
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			hasSpecial = true
		}
	}
	return hasNumber && hasUpperCase && hasLowercase && hasSpecial
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	password := scanner.Text()
	if verifyPassword(password) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
