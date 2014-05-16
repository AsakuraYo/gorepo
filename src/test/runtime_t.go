package main

import (
    "runtime"
    "fmt"
)

func main () {
    fmt.Println("Compiler", runtime.Compiler)
    fmt.Println("GOARCH", runtime.GOARCH)
    fmt.Println("GOOS", runtime.GOOS)
    fmt.Println("NumCPU", runtime.NumCPU())
    fmt.Println("NumGoroutine", runtime.NumGoroutine())
    fmt.Println("Version", runtime.Version())
}
