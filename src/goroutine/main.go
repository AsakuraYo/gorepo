package main

import "fmt"
import "time"
import "math/rand"


func workRoutine(ch chan int, workerNo int) {
    r := rand.New(rand.NewSource(time.Now().Unix()))
    for n := 0; n <= workerNo; n++ {
        sec := r.Intn(10)
        if n == workerNo {
            fmt.Println("Worker", workerNo, "will work", sec, "second.")
            time.Sleep(time.Duration(sec) * time.Second)
            break
        }
    }
    fmt.Println("Worker", workerNo, "finish")
    ch <- workerNo
}


func main() {
    const taskCount = 10
    const timeLimit = 5

    // start tasks
    chs := make([]chan int, taskCount)
    for i := 0; i < taskCount; i++ {
        chs[i] = make(chan int)
        go workRoutine(chs[i], i)
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

