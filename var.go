package main

import "fmt"

var x int    
var y = 42
var z = "siratul"
const pi = 3.1416

func main() {
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	fmt.Println("pi:", pi)
	z := "gophers!" // only inside func body
	fmt.Println("z:", z)
}
