package main

import (
	"testing"
)

func getArgs() []string {
	args := make([]string, 100)
	for i := 0; i < 100; i++ {
		args[i] = "test"
	}
	return args
}

func BenchmarkEchoSlow(b *testing.B) {
	args := getArgs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		EchoSlow(args)
	}
}

func BenchmarkEchoFast(b *testing.B) {
	args := getArgs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		EchoFast(args)
	}
}
