package main

import "fmt"

func printResult(v float64){
	if v >= 7{
		fmt.Println("passed")
	} else{
		fmt.Println("failed")
	}
}


func main() {
	printResult(6.5)
	printResult(4.0)
	printResult(9.0)
}