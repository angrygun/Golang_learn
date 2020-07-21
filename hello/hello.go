package main

import "fmt"

const englishHelloPrefix = "Hello "

// Hello function for learning go TDD
// @title Hello
// @description ro learn
// @param name string
// @return string
func Hello(name string, language string) string {
    if name == "" {
        name = "World"
    }
  
    if language == "Spanish" {
        return "Hola " + name + ""
    }
  
  	return englishHelloPrefix + name
}

func main() {
    fmt.Println(Hello("World", ""))
}
