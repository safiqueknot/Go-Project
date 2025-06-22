package main

import (
	"fmt"
)

func main() {
    age := 5
    if age < 18 { // greater than
        fmt.Println("You are not eligible to be married.")
    } else if age > 18 { //less the
        fmt.Println("You are eligible to be married.")
    } else if age == 18{
        fmt.Println("You are eligible to be married and have a family.")
    }
}