package main

import "fmt"

func whatTime(time int){
	if time < 11 {
		fmt.Println("Good Morning!")
	}else if time >= 11 && time < 18{
		fmt.Println("Good Evening!")
	}else{
		fmt.Println("Good Night!")
	}
}


func main() {
	whatTime(10)
	whatTime(11)
	whatTime(12)
	whatTime(19)

}