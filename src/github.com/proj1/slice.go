package main

import "fmt"

func slice1() {
	x := make([]int, 0)

	x = append(x, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	fmt.Println(x)
}

func slice2() {
	x := make([]int, 3, 10)
	// 	x[1] = 2
	x = append(x, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("x ->", x)
	fmt.Println("x[2:5] ->", x[2:5])
	fmt.Println("x[:5] ->", x[:5])
	fmt.Println("x[:] ->", x[:])
}

func stringSlice() {
	s := make([]string, 3, 10)
	fmt.Println("string slice made is: ", s)
	s[0] = "a"
	s[2] = "b"
	s = append(s, "c", "d")
	fmt.Println(s)
}

func main() {
	// slice1()
	// slice2()
	stringSlice()
}
