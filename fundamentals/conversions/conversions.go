package main

import "fmt"
import "strconv"

func main() {
	x := 2.5
	y := 2
	fmt.Println(x/float64(y))

	n := 6.9
	fmt.Println(int(n))	

	fmt.Println("A " + string(rune(98))) // ascii table
	fmt.Println("B " + strconv.Itoa(200)) // int to string

	num, _ := strconv.Atoi("Ana") // string to int
	fmt.Println(num)

	f, _ := strconv.ParseBool("true")

	if f{
		fmt.Println("Its true!")
	}
}