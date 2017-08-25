package main

import "fmt"
import "sort"

func sortfun() {
	s := []string{"z", "x", "b", "z", "y"}
	sort.Strings(s)
	fmt.Println("sorted s is: ", s)

	nums := []int{3, 4, 2, 5, 1}
	sort.Ints(nums)
	fmt.Println("sorted number is: ", nums)
}
