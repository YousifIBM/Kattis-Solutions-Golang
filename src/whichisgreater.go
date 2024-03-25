package main

import "fmt"

func main() {
    var a, b uint64
    fmt.Scan(&a, &b)
    if a > b {
        fmt.Println(1)
    } else if a < b {
        fmt.Println(0)
    } else {
        fmt.Println(0)
    }
}
