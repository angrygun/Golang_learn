package main

import (
	"strings"
	"testing"
)

func BenchmarkStingPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := []string{"Welcome", "To", "China"}
		s, sep := "", ""
		for j := 0; j < len(input); j++ {
			s += sep + input[j]
			sep = " "
		}

		if s != "Welcome To China" {
			b.Error("Unexcepted result:" + s)
		}
	}
}

func BenchmarkStringJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := []string{"Welcome", "To", "China"}
		result := strings.Join(input, " ")

		if result != "Welcome To China" {
			b.Error("Unexcepted result:" + result)
		}
	}
}
