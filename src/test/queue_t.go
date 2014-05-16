package main

import (
    "queue"
    "runtime"
    "time"
    "fmt"
)

var q *queue.TaskQueue = queue.NewTaskQueue()

func PushRoutine(finish chan bool) {
    for i := 1; i < 10; i++ {
        q.Push(i)
        fmt.Println("Push", i)
        q.Push(i)
        fmt.Println("Push", i)
        q.Push(i)
        fmt.Println("Push", i)
        time.Sleep(4 * time.Second)
    }
    finish <- true
}

func PopRoutine(finish chan bool, num int) {
    for i := 1; i < 15; i++ {
        if tmp, err := q.Pop(); err != nil {
            fmt.Printf("Pop[%d] timeout\n", num)
        } else {
            fmt.Printf("Pop[%d] %d\n", num, tmp)
        }
        time.Sleep(2 * time.Second)
    }
    finish <- true
}

func main() {
    q = queue.NewTaskQueue()

    chs := make([]chan bool, 3)
    chs[0] = make(chan bool)
    chs[1] = make(chan bool)
    chs[2] = make(chan bool)
    go PushRoutine(chs[0])
    go PopRoutine(chs[1], 1)
    go PopRoutine(chs[2], 2)

    for {
        runtime.Gosched()
        if <-chs[0] && <-chs[1] && <-chs[2] {
            break;
        }
    }
    //q.Pop()
}

