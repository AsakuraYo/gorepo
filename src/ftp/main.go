package main

import (
    "fmt"
    ftp4go "code.google.com/p/ftp4go"
)

func main () {
    ftp := ftp4go.NewFTP(2)
    ftp.Connect("127.0.0.1", 21, "")
    ftp.Login("ljq", "070521", "PASV")
    ftp.Cwd("/home/ljq/test")
    dir, _ := ftp.Pwd()
    fmt.Println(dir)
    size, err := ftp.Size("/home/ljq/test/a")
    fmt.Println(size)
}
