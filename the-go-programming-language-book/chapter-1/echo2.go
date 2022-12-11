// another example of reading command line arguments mimicing 'echo' tool
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string

	// os.Args[0] is name of current executable, so passed in args start at [1]
	/* standard for loops follow the pattern
	for 'initialization'; 'condition'; 'post' {
		...
	}

	we can also iterate over slices
	for index, val := range os.Args[1:] {
		...

	}
	*/

	for _, val := range os.Args[1:] {
		// index is not being used, so tell go to ignore it by naming it _
		s += sep + val
		sep = " " // no need to reset this each loop, need it empty on first pass
	}

	fmt.Println("The name of this program is ", os.Args[0])
	fmt.Println("the args passed to this program are: ", s)
}
