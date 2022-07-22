package main

import (
	"fmt"
	"os"
	"strings"
)

// This is the first version.
//func main() {
//	var s, sep string
//
//	fmt.Println(os.Args)
//	for i := 1; i < len(os.Args); i++ {
//		s += sep + os.Args[i]
//		sep = " "
//	}
//
//	fmt.Println(s)
//}

func main() {
	fmt.Println(strings.Join(os.Args[:], " "))

	for i := 0; i < len(os.Args); i++ {
		fmt.Println(i)
		fmt.Println(os.Args[i])
	}
}
