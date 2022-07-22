package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Command to execute: go run duplicate_line3.go ./../data/name.txt

func main() {
	counts := make(map[string]int)

	for _, fileName := range os.Args[1:] {
		data, err := ioutil.ReadFile(fileName)

		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}

		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
