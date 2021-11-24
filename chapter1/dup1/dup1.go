// Dup1 prints the text of each line that appears more than once in the standard input, preceded by its count.
package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
	map holds set of key/value pairs and provides constant-time operations to store, retrieve or test for an item in the set.
	The key: can be of any type whose values can compared with ==
	The value: can be of any type at all.
	make: creates a new empty map
*/

func main() {
	// create a new empty map, key: string, value: int
	counts := make(map[string]int)
	// create a convenient interface for reading data such as a file of lines ends with CRLF.
	input := bufio.NewScanner(os.Stdin)
	// call the next token every iteration and remove the new line from the end.
	// returns true if there's a line, otherwise, false
	for input.Scan() {
		/*
			line := input.Text()
			counts[line] = counts[line] + 1

			convert the input into a text [line]
		*/
		if input.Text() == "done" {
			break
		}
		counts[input.Text()]++
	}
	//	NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s", n, line)
		}
	}
}
