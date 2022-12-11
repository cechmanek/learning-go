// find duplicate lines in a file and print their number of occurences
// call this with '$go run dup1.go < text_file_name'
// or compile and run binary '$go build dup1.go && ./dup1.exe < text_file_name'
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int) // creates a map with string keys & int values
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() { // 'while input.Scan==True' in other languages
		counts[input.Text()]++ // won't raise error for new keys, they get initialized
	}
	// this is just a basic counting hash
	// note: we're ignoring potential errors from input.Err()

	// 
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}
