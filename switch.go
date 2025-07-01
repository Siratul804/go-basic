package main

import (
	"fmt"
	"time"
)

var dayNow = time.Now().Weekday()

func main() {
	switch day := time.Now().Weekday(); day {
	case time.Saturday, time.Sunday:
		fmt.Println("weekend")
	default:
		fmt.Println("weekday")
	}

	fmt.Println(dayNow)

}
