package main

import "fmt"

func sum(nums ...int) {
	fmt.Println("number received: ", nums)
	total := 0

	for _, n := range nums {
		total += n
	}
	fmt.Println("the total is: ", total)
}

func sumIndexed(nums ...int) {
	fmt.Println("number received: ", nums)
	total := 0

	for i, n := range nums {
		total += n
		fmt.Println("index is: ", i)
	}
	fmt.Println("the total is: ", total)
}

func main() {
// 	n := []int{5, 6, 7}
// 	sum(n...)
// 	sumIndexed(n...)
    // recursion()
    sortfun()
    structUsage()
}
