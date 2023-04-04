package main

import (
	"fmt"

	"practis/youtube/genericlist"
)

func main() {
	glist := genericlist.New[string]()

	glist.Insert("Bob")   //0
	glist.Insert("John")  //1
	glist.Insert("Alice") //2
	glist.Insert("Hal")   //3

	glist.RemoveByValue("Alice") //
	glist.Remove(2)
	fmt.Printf("%+v\n", glist)
}
