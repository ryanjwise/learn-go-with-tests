package main

import (
	"fmt"
	"learn-go-with-tests/m/pkg/adder"
	"learn-go-with-tests/m/pkg/hello"
)

func runAdder(x, y int) {
	fmt.Printf("Assuming it is passed numbers '%d', and '%d', the result of adder is: %d\n", x, y, adder.Add(x, y))
}

func runHello() {
	name := ""
	language := ""
	fmt.Printf("When passed '%v' and '%v', the result of hello is: %v\n", name, language, hello.Hello(name, language))
	name = "Frank"
	fmt.Printf("When passed '%v' and '%v', the result of hello is: %v\n", name, language, hello.Hello(name, language))
	language = "Spanish"
	fmt.Printf("When passed '%v' and '%v', the result of hello is: %v\n", name, language, hello.Hello(name, language))
	language = "French"
	fmt.Printf("When passed '%v' and '%v', the result of hello is: %v\n", name, language, hello.Hello(name, language))
}

func main() {
	runAdder(2, 3)
	runAdder(5, 15)
	runHello()
}
