package main

import "fmt"
import "math"

func main() {
	const PI float64 = 3.14
	var radius = 3.2

	area := PI * math.Pow(radius, 2)
	fmt.Print("Area:", area,"\n")

	const (
		a = 1
		b = 2
	)
	
	fmt.Print(a+b,"\n")

	var (
		c = 1
		d = 2
	)

	fmt.Print(c+d,"\n")

	g, h, i := 2, "hi", true

	fmt.Println(g,h,i,)
}