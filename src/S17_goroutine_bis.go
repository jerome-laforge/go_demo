package main

import (
    "time"
    "fmt"
    "runtime"
)

func f(left, right chan int) {
    left <- 1 + <-right
}

func main() {
    fmt.Println(runtime.GOMAXPROCS(runtime.NumCPU()))
    fmt.Println(runtime.GOMAXPROCS(runtime.NumCPU()))
    start := time.Now()
    const n = 10000
    leftmost := make(chan int)

    right := leftmost
    left := leftmost
    for i := 0; i < n; i++ {
        right = make(chan int)
        go f(left, right)
        left = right
    }

    fmt.Println(runtime.NumGoroutine())

    go func(c chan int) { c <- 0 }(right)

    fmt.Println(<-leftmost, time.Since(start))
}
