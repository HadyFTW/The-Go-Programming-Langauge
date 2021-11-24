package main

import (
	"fmt"
	"os"
	"strings"
)

/*
Exercise 1.1: Modify the echo program to also print os.Args[0], the name of the command that invoked it.
*/
func exercise11() {
	fmt.Println("The name of the command that invoked it:", os.Args[0])
	fmt.Println(strings.Join(os.Args[1:], " "))
}

/*
Exercise 1.2: Modify the echo program to print the index and value of each of its arguments, one per line.
*/
func exercise12() {
	for index, arg := range os.Args[1:] {
		fmt.Println("index = ", index, "argument = ", arg)
	}
}

func main() {
	exercise11()
	fmt.Println("###################################")
	exercise12()
}
