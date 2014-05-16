package main

import (
    "fmt"
)

func main () {
    fmt.Println("Hello")

    s := make([]int, 0, 3)
    for i := 0; i < 64; i++ {
        s = append(s, i)
        fmt.Println(len(s), cap(s))
    }

    var s2 []int
    for i := 0; i < 64; i++ {
        s2 = append(s2, i)
        fmt.Println("s2", len(s2), cap(s2))
    }
}
