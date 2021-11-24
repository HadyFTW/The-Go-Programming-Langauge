package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

/*
	goroutine: concurrent function execution
	channel: communication mechanism that allows one goroutine to pass values
	of a SPECIFIED TYPE to another goroutine.

	the function main runs in a goroutine by default and the go statement creates additional goroutines.
*/

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	/*
		having main do all the printing ensures that output
		from each goroutine is processed as a unit with no danger
		of interleaving if two goroutines finish at the same time.
	*/
	for _, url := range os.Args[1:] {
		fileToSave := strings.Split(url, "//")[1] + ".txt"
		data := fmt.Sprint(<-ch)
		err := ioutil.WriteFile(fileToSave, []byte(data), 0755)
		if err != nil {
			fmt.Printf("Unable to write file: %v", err)
		}
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	//start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	// nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	data, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("While reading %s: %v", url, err)
		return
	}
	//secs := time.Since(start).Seconds()
	ch <- string(data)
}
