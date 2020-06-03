package main

import (
	"fmt"
	"os"
)

// func main() {
// 	var s, sep string
// 	for i := 0; i < len(os.Args); i++ {
// 		s += sep + os.Args[i]
// 		sep = " "
// 	}
// 	fmt.Println(s)
// }

// Print command name at the same time
func main() {
	var s, sep string = "", ""
	for i := 0; i < len(os.Args[i]); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
