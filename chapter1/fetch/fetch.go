// Fetch prints the content found at a URL
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// resp.Body contains the server response as a stream data.
		b, err := ioutil.ReadAll(resp.Body)
		// close the stream to avoid leacking resource.
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", b)

	}
}
