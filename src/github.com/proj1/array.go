package main

import "fmt"

func arr1() {
	var a [7]int
	a[0] = 1
	a[6] = 7

	fmt.Println(a[4])
	fmt.Println(a[6])
	fmt.Println(a)
}

func arr2() {
	a := [3]string{"apple", "banana", "cranberry"}
	fmt.Println(a)
	fmt.Println(len(a))
}
func main() {
	// 	arr1()
	arr2()
}
