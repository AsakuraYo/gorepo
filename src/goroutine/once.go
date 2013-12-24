package main

import "sync"

var a string
var once sync.Once

func setup() {
    print("setup...\n")
    a = "Hello world"
}

func doPrint(ch chan int) {
    once.Do(setup)
    print(a + "\n")
    ch <- 1
}

func main() {
    chs := make([]chan int, 2)
    chs[0] = make(chan int)
    chs[1] = make(chan int)
    go doPrint(chs[0])
    go doPrint(chs[1])

    for _, ch := range(chs) {
        <-ch
    }
}
