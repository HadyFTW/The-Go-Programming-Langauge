package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	filenames := os.Args[1:]
	if len(filenames) == 0 {
		fmt.Println("There's no input file!")
	} else {
		for _, filename := range filenames {
			data, err := ioutil.ReadFile(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Dup3: %v\n", err)
				continue
			}
			fmt.Printf("Type of data: %T\n", data) // []uint8 - byte.
			for _, line := range strings.Split(string(data), "\n") {
				counts[line]++
			}

			for line, nCounts := range counts {
				if nCounts > 1 {
					fmt.Println(line, nCounts)
				}
			}

		}
	}

}
