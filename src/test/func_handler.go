package main

import (
    "fmt"
    "time"
    "runtime"
)

var isShutDown = false;

func routineMain(routineID int, exitChan chan bool) {
    start := true
    for {
        select {
            case <-exitChan:
            break
            case start:
        }
        fmt.Println("Routine", routineID)
        time.Sleep(2 * time.Second)
    }
}

type RoutineManager struct {
    counter int
    routineID []int
    exitChans []chan bool
}

func NewRoutineManager() *RoutineManager {
    rm := &RoutineManager{ counter : 0 }
    return rm
}

func (rm *RoutineManager) CreateRoutine(handler func(int, chan bool)) {
    rm.counter++
    rm.routineID = append(rm.routineID, rm.counter)
    rm.exitChans = append(rm.exitChans, make(chan bool, 1))
    go handler(rm.counter, rm.exitChans[rm.counter - 1])
    rm.exitChans[rm.counter - 1] <- false
}

func main () {
    rtManager := NewRoutineManager()
    rtManager.CreateRoutine(routineMain)
    rtManager.CreateRoutine(routineMain)

    fmt.Println("-->NumGoroutine", runtime.NumGoroutine())
    time.Sleep(20 * time.Second)
    isShutDown = true;
}
