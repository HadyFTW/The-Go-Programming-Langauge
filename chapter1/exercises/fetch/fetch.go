// Fetch prints the content found at a URL
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// resp.Body contains the server response as a stream data.
		fmt.Println(resp.Status)
		_, errs := io.Copy(os.Stdout, resp.Body)
		// close the stream to avoid lacking resource.
		resp.Body.Close()
		if errs != nil {
			fmt.Fprintf(os.Stderr, "reading %s: %v\n", url, err)
			os.Exit(1)
		}

	}
}
