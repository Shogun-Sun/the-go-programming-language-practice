package main

import (
	"strings"
)

func EchoSlow(args []string) string {
	var s, sep string
	for i := 0; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}

	return s
}

func EchoFast(args []string) string {
	return strings.Join(args, " ")
}
