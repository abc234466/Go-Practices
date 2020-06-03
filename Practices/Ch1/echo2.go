package main

import (
	"fmt"
	"os"
)

// func main() {
// 	s, sep := "", ""
// 	// pairs fo index, arg
// 	for _, arg := range os.Args[1:] {
// 		s += sep + arg
// 		sep = " "
// 	}
// 	fmt.Println(s)
// }

// Print index, value every line

func main() {
	for line, arg := range os.Args[0:] {
		fmt.Println(line, arg)
	}
}
