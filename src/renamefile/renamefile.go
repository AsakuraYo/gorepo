package main

import "fmt"
import "os"

func main() {
    if len(os.Args) != 3 {
        fmt.Println("Usage:", os.Args[0], "<oldPath> <newPath>")
        os.Exit(1)
    }
    if err := os.Rename(os.Args[1], os.Args[2]); err != nil {
        fmt.Println("Rename", os.Args[1], "->", os.Args[2], "ERROR:", err.Error())
    }
}
