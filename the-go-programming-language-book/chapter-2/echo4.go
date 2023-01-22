// echo4 prints its command-line arguments and uses cmd line flags.
// This file demonstrates us of pointers.
package main

import (
	"flag" // allows you to access command line flags
	"fmt"
	"strings"
)

// var n = flag.Bool("<flag_name>", default, "<comment>")
var n = flag.Bool("n", false, "omit training newline") // returns a pointer
var sep = flag.String("s", " ", "separator") // returns a pointer

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep)) // sep is pointer, so dereference
	if !*n { // default value of 'n' is 'false'. Add newline if n is 'true'
		fmt.Println()
	}
}

/* examples of how to run this file:

go build echo4.go

./echo4.exe a bc def
>> a bc def

./echo4.exe -s -- a bc def
>> a--bc--def

./echo4.exe -n a bc def
>> a bc def >>$

./echo4.exe -help
Usage of D:\Projects\learning-go\the-go-programming-language-book\chapter-2\echo4.exe:
  -n    omit training newline
  -s string
        separator (default " ")

*/