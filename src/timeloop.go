package main

import (
    "fmt"
)

func main() {
    var N int
    fmt.Scan(&N)

    if N < 1 || N > 100 {
        fmt.Println("Please enter a number between 1 and 100.")
        return
    }

    for i := 1; i <= N; i++ {
        fmt.Printf("%d Abracadabra\n", i)
    }
}
