package main

import (
    "bufio"
    "os"
    "io"
    "fmt"
)

func main () {
    target, err := os.OpenFile("/tmp/c", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0664)
    if err != nil {
        fmt.Println(err)
        return
    }
    //defer target.Close()
    n, err := target.WriteString("ABC\n")
    fmt.Println(n, err)

    writer := bufio.NewWriter(target)
    n, err = io.WriteString(writer, "Hello?")
    fmt.Println(n, err)
}
