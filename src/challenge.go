package main

import (
    "fmt"
)

func findReplacementDay(P, N int, parts []string) int {
    partCount := make(map[string]int)
    replacedCount := 0

    for day, part := range parts {
        partCount[part]++

        if partCount[part] == 1 {
            replacedCount++
            if replacedCount == P {
                return day + 1
            }
        }
    }

    return -1 // Paradox avoided
}

func main() {
    var P, N int
    fmt.Scan(&P, &N)

    parts := make([]string, N)
    for i := 0; i < N; i++ {
        fmt.Scan(&parts[i])
    }

    result := findReplacementDay(P, N, parts)

    if result == -1 {
        fmt.Println("paradox avoided")
    } else {
        fmt.Println(result)
    }
}
