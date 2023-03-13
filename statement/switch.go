package main

import "fmt"

func whatDaySingleCase(day int) {
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")
	}
}

func whatDayMultiCase(day int) {
	switch day {
	case 1, 3, 5:
		fmt.Println("Odd weekday")
	case 2, 4:
		fmt.Println("Even weekday")
	case 6, 7:
		fmt.Println("Weekend")
	}

}

func main() {
	whatDaySingleCase(1)
	whatDaySingleCase(5)
	whatDaySingleCase(6)
	whatDaySingleCase(2)
	whatDaySingleCase(9)

	whatDayMultiCase(4)
	whatDayMultiCase(3)
	whatDayMultiCase(6)

}
