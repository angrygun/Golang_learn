package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	stringPlus()

	stringJoin()
}

func stringPlus() {
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}
	fmt.Println("stringPlus", s)
	fmt.Printf("stringPlus time: %.10fs elapsed\n", time.Since(start).Seconds())
}

func stringJoin() {
	start := time.Now()
	s := strings.Join(os.Args, " ")
	fmt.Println("stringJoin", s)
	fmt.Printf("stringJoin time: %.10fs elapsed\n", time.Since(start).Seconds())
}
