package main

import (
    "fmt"
)

func main() {
    var input float64
    fmt.Print("Enter a floating point number: ")
    fmt.Scan(&input)

    truncated := int(input)
    fmt.Printf("Truncated integer: %d\n", truncated)
}
