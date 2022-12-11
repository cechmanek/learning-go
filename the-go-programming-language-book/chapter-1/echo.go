// a simple example of reading command line arguments mimicing 'echo' tool
package main

import (
	"fmt"
	"os"
)

func main() {
	var s string   // type specification comes after variable name
	var sep string // no semi-colon needed
	//var s, sep string // can declare multiple vars of the same type on one line
	// s and sep are declared, but not defined so are initialized as empty strings

	// os.Args[0] is name of current executable, so passed in args start at [1]
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " " // no need to reset this each loop, need it empty on first pass
	}

	fmt.Println("The name of this program is ", os.Args[0])
	fmt.Println("the args passed to this program are: ", s)
}
