
package main

import "fmt"

func main() {
    fmt.Println("GO Condition")

    a := 4

	// Go  te not use parentheses
    if a >= 60 {
        fmt.Println("Excellent")
    } else if a >= 40 && a < 60 {
        fmt.Println("Fixed Golang")
    } else {
        fmt.Println("Not Found")
    }

    fmt.Println("GO Understanding")
}
