package problem3

import (
    "fmt"
    "time"
)

func RunSolution() {
    start := time.Now()

    solution := Solve()

    duration := time.Since(start)

    fmt.Printf("Solution: %d", solution)
    fmt.Println("\nTime to solve:")
    fmt.Println(duration)
}