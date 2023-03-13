package main

import "fmt"

func main() {
	c := 0

	var p *int = nil

	p = &c // address of c
	*p++ // increment c value
	c++

	fmt.Println(p, *p, c, &c)
}