package helpers_test

import (
	"fmt"
	"go-template/config/helpers"
)

// Add returns the sum of two integers.
func Add(a, b int) int {
    return a + b
}

// ExampleAdd demonstrates how to use the Add function.
//
// This example demonstrates the usage of the Add function by adding two integers.
//    sum := Add(1, 2)
//    fmt.Println(sum) // Output: 3
func ExampleAdd() {
    sum := Add(1, 2)
    fmt.Println(sum)
    // Output: 3
}


func ExampleGetEnv() {
    fmt.Println(helpers.GetEnv("NOT-SET", "default"))
    // Output: default
}