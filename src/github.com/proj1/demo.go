package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func Boolean() bool {
	var b bool
	b = true
	return b
}

func showvar() {
	var a string = "mystring"
	fmt.Println(a)
	var e int
	fmt.Println(e)
	var b, c int = 1, 2
	fmt.Println(b, c)
	d := "Golang"
	f := a + " : " + d
	fmt.Println(f)
}
func main() {
	fmt.Println("Welcome to the first Go Project")
	c := add(5, 6)
	fmt.Println(c)
	d := Boolean()
	fmt.Println(d)
	showvar()
}
