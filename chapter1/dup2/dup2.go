// Dup2 prints the count and text of lines that appear more than once
// in the input. It reads from stdin or from a list of named files.
// This program operates in "streaming" mode using `bufio` in which input is read and broken into lines as needed.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		displayCounts(countLines(os.Stdin, counts))
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Dup2: %v\n", err)
				continue
			}
			displayCounts(countLines(f, counts))
			f.Close()
		}
	}

}

func countLines(f *os.File, counts map[string]int) map[string]int {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if input.Text() == "done" {
			break
		}
		counts[input.Text()]++
	}
	return counts
}

func displayCounts(counts map[string]int) {
	for line, nCounts := range counts {
		fmt.Printf("%s\t\t%d\n", line, nCounts)
	}
}
