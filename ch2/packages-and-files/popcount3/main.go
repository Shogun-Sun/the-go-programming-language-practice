package main

import "fmt"

func main() {
	fmt.Println(PopCount(20))
}

func PopCount(x uint64) int {
	var count int
	for i := 0; i < 64; i++ {
		count += int(x & 1)
	}

	return count
}
