// main.go
package main

import (
    "config"
    "fmt"
)

func main() {
    fmt.Println(config.LoadIni())
    fmt.Println("hello")
}
