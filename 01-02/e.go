// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args {
		fmt.Printf("%v:%v", i, arg)
	}
}

//!-
