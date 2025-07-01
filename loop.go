package main

import "fmt"

func main() {
	// for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// while-loop 
	sum := 1
	for sum < 10 {
		sum += sum
		fmt.Println("Current sum:", sum) // Added print to see progress
	}

	// infinite loop
	// Adding break condition to prevent infinite execution
	count := 0
	for {
		fmt.Println("loop iteration:", count)
		count++
		if count >= 3 {
			break
		}
	}
}
