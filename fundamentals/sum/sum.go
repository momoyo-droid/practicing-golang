package main

import "fmt"

func sum(a int, b int) int { return a + b}

func print(value int) {
	fmt.Println(value)
}

func main() {
	result := sum(10,10)
	print(result)
}