package main

import (
    "fmt"
    "time"
)

var ch chan int = make(chan int, 2)
var finish chan bool = make(chan bool, 1)

func test() {
    ch <- 1
    fmt.Println("write", 1, "done")
    ch <- 2
    fmt.Println("write", 2, "done")
    ch <- 3
    fmt.Println("write", 3, "done")
    time.Sleep(time.Second)
    ch <- 4
    fmt.Println("write", 4, "done")
    ch <- 5
    fmt.Println("write", 5, "done")
    ch <- 6
    fmt.Println("write", 6, "done")

    finish <- true
}

func main() {
    go test()
    for i := 0; i < 6; i++ {
        time.Sleep(2 * time.Second)
        fmt.Println("len", len(ch))
        fmt.Println("get", <-ch)
    }
}
