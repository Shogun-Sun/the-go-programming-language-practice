package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]
	if len(files) == 0 {
		counts := make(map[string]int)
		countLines(os.Stdin, counts)
		printDuplicates(counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}

			fileCounts := make(map[string]int)

			countLines(f, fileCounts)
			f.Close()

			if hasDuplicates(fileCounts) {
				fmt.Printf("Файл %s содержит повторы:\n", arg)
				printDuplicates(fileCounts)
			}
		}
	}
}

func printDuplicates(counts map[string]int) {
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func hasDuplicates(counts map[string]int) bool {
	for _, n := range counts {
		if n > 1 {
			return true
		}
	}

	return false
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line] = counts[line] + 1
	}
}
