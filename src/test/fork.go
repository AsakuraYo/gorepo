package main

import (
    "syscall"
//    "time"
)

func main() {
    pid, _, sysErr := syscall.RawSyscall(syscall.SYS_FORK, 0, 0, 0)
    if pid > 0 {
        println(pid, sysErr)
        return
    } else if pid == 0 {
        println("finish")
        return
    } else {
        println(sysErr)
    }
}
