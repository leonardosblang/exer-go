package main

import (
	"fmt"
)



func main() {
    unsort := []int{25, 6, 2, 1, 5, 23, 3, 4, 7, 9}
    sorted := bubbleSorting(unsort)
    fmt.Println(sorted)
}

func bubbleSorting(input []int) []int {
    swap := true
    for swap {
        swap = false
        for i := 1; i < len(input); i++ {

            if input[i-1] > input[i] {
                input[i], input[i-1] = input[i-1], input[i]
                swap = true
            }
        }
    }
    return input
}