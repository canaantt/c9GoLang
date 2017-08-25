package main

import "fmt"

func loop1() {
	for i := 3; i <= 7; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
}

func loop2() {
	for i := 3; i <= 7; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
}

func loop3() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func loop4() {
	for i := 9; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func loop5() {
	for i := 1; i <= 11; i++ {
		if i < 7 {
			for j := 1; j <= i; i++ {
				fmt.Print("* ")
			}
		} else {
			x := (i - 6) * 2
			for j := i; j > x; j-- {
				fmt.Print("* ")
			}
		}
		fmt.Println()
	}
}
func main() {
	// 	fmt.Println("learning break in forloop")
	// 	loop1()
	// 	fmt.Println("continue in forloop skip")
	// 	loop2()
	// 	fmt.Println("see pattern")
	// 	loop3()
	// 	loop4()
	loop5()
}
