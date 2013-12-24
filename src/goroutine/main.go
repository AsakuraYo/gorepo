package main

import "fmt"
import "time"
import "math/rand"


func Count(ch chan int, i int) {
    r := rand.New(rand.NewSource(time.Now().Unix()))
    for n := 0; n <= i; n++ {
        sec := r.Intn(10)
        if n == i {
            fmt.Println("Task", i, "will work", sec, "second.")
            time.Sleep(time.Duration(sec) * time.Second)
            break
        }
    }
    fmt.Println("Task", i, "finish")
    ch <- i
}


func main() {
    const taskCount = 10
    const timeLimit = 5

    // start tasks
    chs := make([]chan int, taskCount)
    for i := 0; i < taskCount; i++ {
        chs[i] = make(chan int)
        go Count(chs[i], i)
    }

    // start time limit
    timeout := make(chan bool, 1)
    go func() {
        time.Sleep(timeLimit * time.Second)
        fmt.Println("-----Times up-----")
        for {
            timeout <- true
        }
    }()

    // wait for tasks
    for i, ch := range(chs) {
        select {
            case <-ch:
            case <-timeout:
                fmt.Println(i, " timeout")
        }
    }
    fmt.Println("Program exit")
}

