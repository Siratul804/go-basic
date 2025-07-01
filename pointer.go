package main

import "fmt"

func zero(val *int) {
    *val = 0
}

func main() {
	
	x := 5
    zero(&x)
    fmt.Println(x) // 0

}
