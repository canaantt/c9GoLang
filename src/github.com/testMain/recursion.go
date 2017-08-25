package main

import "fmt"

func fact(n int) int {
    if n == 1{
        return 1
    } else {
        return n * fact(n-1)
    }
}

func recursion() {
    z := fact(5)
    fmt.Println("Factorial of %d\n", z)
}