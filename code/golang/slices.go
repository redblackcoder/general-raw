package main

import (
    "fmt"
    "slices"
)

func fib(n int) int {
    if n < 2 {
        return 1
    }

    return fib(n - 2) + fib(n - 1)
}

func main() {
    list := []int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    fmt.Println(list[2:6])

    slices.Reverse(list)
    fmt.Println(list[1:5])

    fmt.Println(fib(5))
}

