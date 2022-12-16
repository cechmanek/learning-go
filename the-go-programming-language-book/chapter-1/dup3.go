// yet another way to detect duplicate lines in a set of files
package main

import (
	"fmt"
	"io/ioutil" // imports only one sub package from io
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename) // returns a byte slice so cast to string later
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
		}

		for _, line := range strings.Split(string(data), "\n") { // cast byte slice to string
			counts[line]++
		}
	}

	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
	}
}
