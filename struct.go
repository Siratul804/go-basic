package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func (p Person) Greet() {
    fmt.Printf("Hi, I'm %s and I'm %d\n", p.Name, p.Age)
}

func main() {
    p := Person{Name: "Siratul", Age: 20}
    p.Greet() // Hi, I'm Siratul and I'm 20
}
