package main

import (
    "fmt"
    "container/list"
    "sync"
    "time"
    "runtime"
)

type Queue struct {
    data *list.List
    mutex sync.Mutex
    cond *sync.Cond
}

func NewQueue() *Queue {
    q := &Queue{ data : list.New() }
    q.cond = sync.NewCond(&q.mutex)
    return q
}

func (q *Queue) Push(item interface{}) {
    q.mutex.Lock()
    q.data.PushBack(item)
    q.cond.Broadcast()
    q.mutex.Unlock()
}

func (q *Queue) Pop() interface{} {
    if q.Empty() {
        q.mutex.Lock()
        q.cond.Wait()
        q.mutex.Unlock()
    }
    q.mutex.Lock()
    elem := q.data.Front()
    item := q.data.Remove(elem)
    q.mutex.Unlock()
    return item
}

func (q *Queue) Len() int {
    q.mutex.Lock()
    l := q.data.Len()
    q.mutex.Unlock()
    return l
}

func (q *Queue) Empty() bool {
    q.mutex.Lock()
    isEmpty := (q.data.Len() == 0)
    q.mutex.Unlock()
    return isEmpty
}

func main () {
    q := NewQueue()

    go func() {
        i := 0
        for {
            q.Push(i)
            i++
            q.Push(i)
            i++
            time.Sleep(5 * time.Second)
        }
    }()
    go func () {
        for {
            q.Pop()
            time.Sleep(2 * time.Second)
        }
    }()

    go func () {
        for {
            fmt.Println("Queue Length:", q.Len())
            time.Sleep(1 * time.Second)
        }
    }()

    go func () {
        for {
            fmt.Println("Queue is empty?", q.Empty())
            time.Sleep(1 * time.Second)
        }
    }()

    for {
        runtime.Gosched()
    }

}
