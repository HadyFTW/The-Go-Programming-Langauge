package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	filenames := os.Args[1:]
	result := countLines(filenames)
	for k, v := range result {
		fmt.Println(k)
		displayCounts(v)
		fmt.Println("####################################")
	}
}

func displayCounts(counts map[string]int) {
	for line, nCounts := range counts {
		if nCounts > 1 {
			fmt.Println(line, nCounts)
		}
	}
}

func countLines(filenames []string) map[string]map[string]int {
	relatedFileToLines := make(map[string]map[string]int)
	for _, filename := range filenames {
		counts := make(map[string]int)
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Dup: %v\n", err)
			continue
		}
		splitFileToLines(data, counts)
		relatedFileToLines[filename] = counts
	}
	return relatedFileToLines
}

func splitFileToLines(data []byte, counts map[string]int) {
	for _, line := range strings.Split(string(data), "\n") {
		counts[line]++
	}
}
