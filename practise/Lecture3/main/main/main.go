package main

import (
	"fmt"
    "fibonachi/fibonachi"
    "factorial/factorial"
)

func main() {
    n := 10
    fibList := fibonachi.Generate(n)
    fmt.Println("Fibonacci list", fibList)


	num := 5
	result := factorial.Calculate(num)
	fmt.Printf("Factorial of %d is %d\n", num, result)
}