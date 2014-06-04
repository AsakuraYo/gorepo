package main

import (
    "fmt"
    "time"
)

func getMessageChannel(msg string, delay time.Duration) <-chan string {
    c := make(chan string)
    go func() {
        for i := 1; i <=3; i++ {
            c <- fmt.Sprintf("%s %d", msg, i)
            time.Sleep(time.Millisecond * delay)
        }
    }()
    return c
}

func main() {
    done := make(chan bool)
    close(done)

    // reading from a closed channels do not block and always return default
    // value for a channel type
    fmt.Println(<-done)
    fmt.Println(<-done)

    finish := make(chan bool)

    go func() {
        fmt.Println("goroutine message")
        // because the finish here is not used as a return value.
        // so we can just close it.
        close(finish)   // or  finish <- true
    }()

    fmt.Println("main routine message")
    <-finish

    c1 := getMessageChannel("first", 300)
    c2 := getMessageChannel("second", 50)
    c3 := getMessageChannel("third", 10)

    for i := 1; i <= 9; i++ {
        select {
        case msg := <-c1:
            println(msg)
        case msg := <-c2:
            println(msg)
        case msg := <-c3:
            println(msg)
        }
    }
}
