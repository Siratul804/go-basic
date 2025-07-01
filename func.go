package main

import "fmt"

func add(a int, b int) int {
    return a + b
}

func swap(a, b string) (string, string) {
    return b, a
}

func main() {
	
	fmt.Println(add(1, 2))
	fmt.Println(swap("siratul", "islam"))
}
