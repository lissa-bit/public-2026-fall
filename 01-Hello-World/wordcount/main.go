//go:build !solution

package main

import (
	"bufio"
	"fmt"
	"os"
)

func countFile(filename string, counts map[string]int) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	input := bufio.NewScanner(file)
	for input.Scan() {
		counts[input.Text()]++
	}

	return input.Err()
}

func main() {
	counts := make(map[string]int)

	for _, filename := range os.Args[1:] {
		if err := countFile(filename, counts); err != nil {
			fmt.Fprintf(os.Stderr, "wordcount: %v\n", err)
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
