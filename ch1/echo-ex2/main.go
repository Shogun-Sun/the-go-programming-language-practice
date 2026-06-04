package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("List of args:")
	for index, item := range os.Args[1:] {
		result := fmt.Sprintf("[%d]: %s", index, item)
		fmt.Println(result)
	}
}
