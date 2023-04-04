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
	var folderList []string
	var folderUnique []string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	N = numbers(scanner.Text())
	for i := 0; i < N; i++ {
		scanner.Scan()
		folderList = append(folderList, scanner.Text())
	}
	folderUnique = getFolderNames(folderList)
	for _, folder := range folderUnique {
		fmt.Println(folder)
	}
}

func getFolderNames(names []string) []string {
	hash := map[string]int{}
	unique := make([]string, 0, len(names))
	existName := map[string]bool{}
	for _, name := range names {
		// name not exist, record into existName and hash, append to unique
		if existName[name] == false {
			hash[name] = 0
			existName[name] = true
			unique = append(unique, name)
			continue
		}
		// name exist, but not exist in hash, it's format must be XXX(X)
		if _, exist := hash[name]; !exist {
			hash[name] = 0
			for {
				hash[name]++
				newName := name + "(" + strconv.Itoa(hash[name]) + ")"
				if existName[newName] == true {
					continue
				}
				existName[newName] = true
				unique = append(unique, newName)
				break
			}
			continue
		}
		// name exist and exist in hash
		for {
			hash[name]++
			newName := name + "(" + strconv.Itoa(hash[name]) + ")"
			if existName[newName] == true {
				continue
			}
			existName[newName] = true
			unique = append(unique, newName)
			break
		}
	}
	return unique
}
